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

type Middleware struct {
	Handler http.Handler
}

func (middleware *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Before handler.  Check Middleware")
	middleware.Handler.ServeHTTP(w, r)
	fmt.Println("After handler")
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello Handler")
		fmt.Println("Handler Success")
	})

	middleware := &Middleware{
		Handler: router,
	}

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	middleware.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "Hello Handler", string(body))
}

func TestServerMiddleware(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello Handler")
		fmt.Println("Handler Success")
	})

	middleware := &Middleware{
		Handler: router,
	}
	server := http.Server{
		Addr:    ":8080",
		Handler: middleware,
	}

	server.ListenAndServe()
}
