package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := NewRouter()

	if r.routes == nil {
		t.Error("expected routes to be initialized")
	}

	r.AddRoute("GET", "/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GET /"))
	})

	r.AddRoute("POST", "/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("POST /"))
	})

	tests := []struct {
		method       string
		path         string
		expectedCode int
	}{
		{"GET", "/", http.StatusOK},
		{"POST", "/", http.StatusOK},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.path, nil)
		rr := httptest.NewRecorder()

		r.ServeHTTP(rr, req)

		if rr.Code != test.expectedCode {
			t.Errorf("expected %d, but got %d", test.expectedCode, rr.Code)
		}
	}
}
