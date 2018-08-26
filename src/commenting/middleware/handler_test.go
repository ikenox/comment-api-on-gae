package middleware_test

import (
	"commenting/middleware"
	"encoding/json"
	"github.com/koron/go-dproxy"
	"github.com/stretchr/testify/assert"
	"google.golang.org/appengine/aetest"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCommentAndReply(t *testing.T) {
	s, done := newTestServerClient(t)
	defer done()

	{
		// post comment
		rec := s.JsonRequest("POST", "/comment", struct {
			PageId string `json:"pageId"`
			Name   string `json:"name"`
			Text   string `json:"text"`
		}{
			"pageId1",
			"commenter1",
			"text1",
		})
		assertStatusCode(t, rec, http.StatusOK)
	}

	// get comment list
	{
		rec := s.JsonRequest("GET", "/comment?pageId=pageId1", nil)
		assertStatusCode(t, rec, http.StatusOK)
		var v interface{}
		bytes, _ := ioutil.ReadAll(rec.Body)
		json.Unmarshal(bytes, &v)
		p := dproxy.New(v)

		comment := p.M("data").A(0).M("comment")
		if val, err := comment.M("text").Value(); err != nil {
			t.Errorf(err.Error())
		} else {
			assert.Equal(t, val, "text1")
		}
		if val, err := comment.M("pageId").Value(); err != nil {
			t.Errorf(err.Error())
		} else {
			assert.Equal(t, val, "pageId1")
		}
		if _, err := comment.M("commentedAt").String(); err != nil {
			t.Errorf(err.Error())
		}
		if _, err := comment.M("commentId").Int64(); err != nil {
			t.Errorf(err.Error())
		}

		commenter := p.M("data").A(0).M("commenter")
		if val, err := commenter.M("name").String(); err != nil {
			t.Errorf(err.Error())
		} else {
			assert.Equal(t, val, "commenter1")
		}
		if _, err := commenter.M("commenterId").Int64(); err != nil {
			t.Errorf(err.Error())
		}
	}

}

func assertStatusCode(t *testing.T, rec *httptest.ResponseRecorder, code int) {
	t.Helper()
	if rec.Code != code {
		body, _ := ioutil.ReadAll(rec.Result().Body)
		t.Errorf("status code expected %v got %v, response body: %s", code, rec.Code, string(body))
	}
}

type testServerClient struct {
	t       *testing.T
	inst    aetest.Instance
	handler http.Handler
}

func newTestServerClient(t *testing.T) (*testServerClient, func()) {
	inst, err := aetest.NewInstance(&aetest.Options{StronglyConsistentDatastore: true})
	if err != nil {
		t.Fatal(err)
	}
	handler := middleware.NewServer()
	return &testServerClient{
			inst:    inst,
			handler: handler,
			t:       t,
		}, func() {
			inst.Close()
		}

}

func (s *testServerClient) JsonRequest(method string, url string, params interface{}) *httptest.ResponseRecorder {
	marshal, err := json.Marshal(params)
	if err != nil {
		s.t.Fatal(err)
	}

	req, err := s.inst.NewRequest(method, url, strings.NewReader(string(marshal)))
	if err != nil {
		s.t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	s.handler.ServeHTTP(rec, req)
	return rec
}
