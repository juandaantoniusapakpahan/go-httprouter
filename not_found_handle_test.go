package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestNotFoundHandler(t *testing.T) {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Not Found Brooo")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/ggwp", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "Not Found Brooo", string(body))
}

func TestServerNotFoundHandler(t *testing.T) {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Not Found Brooo")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()

}
