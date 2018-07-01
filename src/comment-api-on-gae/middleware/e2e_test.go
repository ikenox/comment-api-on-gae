package middleware_test

import (
	"comment-api-on-gae/middleware"
	"encoding/json"
	"google.golang.org/appengine/aetest"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestComment(t *testing.T) {
	s, done := newTestServerClient(t)
	defer done()

	rec := s.Request("POST", "/comment", struct {
		PageId string
		Name   string
		Text   string
	}{
		"Valid-page_id",
		"commenter1",
		"text",
	})
	if rec.Code != http.StatusCreated {
		t.Errorf("got %v\nwant %v", rec.Code, http.StatusCreated)
	}

	rec = s.Request("GET", "/comment", nil)
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
	opt := aetest.Options{StronglyConsistentDatastore: true}
	inst, err := aetest.NewInstance(&opt)
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
