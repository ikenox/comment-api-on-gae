package middleware_test

import (
	"commenting/middleware"
	"encoding/json"
	"google.golang.org/appengine/aetest"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCommentAndView(t *testing.T) {
	s, done := newTestServerClient(t)
	defer done()

	rec := s.Request("POST", "/comment", struct {
		PageId string
		Name   string
		Text   string
	}{
		"pageId1",
		"commenter1",
		"text",
	})
	if rec.Code != http.StatusOK {
		t.Errorf("got %v\nwant %v", rec.Code, http.StatusOK)
	}

	rec = s.Request("GET", "/comment?PageId=pageId1", nil)
	if rec.Code != http.StatusOK {
		t.Errorf("got %v\nwant %v", rec.Code, http.StatusOK)
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
	handler := middleware.NewHandler()
	return &testServerClient{
			inst:    inst,
			handler: handler,
			t:       t,
		}, func() {
			inst.Close()
		}
}

func (s *testServerClient) Request(method string, url string, params interface{}) *httptest.ResponseRecorder {
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
