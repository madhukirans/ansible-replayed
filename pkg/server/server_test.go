package server

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/madhukirans/replayed/pkg/types"
)

func TestHealthCheckHandler(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	config = types.GetReplayedConfig()
	//StartServer(config)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHandler(t *testing.T) {
	tt := []struct {
		name  string
		value string
	}{
		{name: "test1", value: "abcdef"},
		{name: "test2", value: ""},
		{name: "test3", value: "x"},
	}

	config = types.GetReplayedConfig()
	InitServer(config)
	var body string
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("POST", "/", strings.NewReader(tc.value))
			if err != nil {
				t.Fatal(err)
			}
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(Handler)
			handler.ServeHTTP(rr, req)
			assert.Equal(t, http.StatusAccepted, rr.Code)

			req1, err1 := http.NewRequest("GET", "/", nil)
			if err1 != nil {
				t.Fatal(err1)
			}
			rr1 := httptest.NewRecorder()
			handler1 := http.HandlerFunc(Handler)
			handler1.ServeHTTP(rr1, req1)
			assert.Equal(t, http.StatusOK, rr1.Code)

			body = body + tc.value
			//assert.Equal(t, w.Body.String(), body)
		})
	}
}

func BenchmarkHttp(b *testing.B) {
	config = types.GetReplayedConfig()
	InitServer(config)
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest("POST", "/", strings.NewReader("xxx"))
		if err != nil {
			b.Fatal(err)
		}
		w := httptest.NewRecorder()
		Handler(w, req)

		req, err = http.NewRequest("GET", "/", nil)
		if err != nil {
			b.Fatal(err)
		}
		w = httptest.NewRecorder()
		Handler(w, req)
	}
	//
	//config = types.GetReplayedConfig()
	//StartServer(config)
	//rr := httptest.NewRecorder()
	//handler := http.HandlerFunc(Handler)
	//rr1 := httptest.NewRecorder()
	//handler1 := http.HandlerFunc(Handler)
	//
	//for i := 0; i < b.N; i++ {
	//	req, err := http.NewRequest("POST", "/", strings.NewReader(""))
	//	if err != nil {
	//		b.Fatal(err)
	//	}
	//
	//	handler.ServeHTTP(rr, req)
	//	assert.Equal(b, http.StatusAccepted, rr.Code)
	//
	//	req1, err1 := http.NewRequest("GET", "/", nil)
	//	if err1 != nil {
	//		b.Fatal(err1)
	//	}
	//
	//	handler1.ServeHTTP(rr1, req1)
	//	assert.Equal(b, http.StatusOK, rr1.Code)
	//}
}

