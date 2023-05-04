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

func TestHttpRouterParam(t *testing.T) {
	router := httprouter.New()
	router.GET("/product/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		result := "Product " + id
		fmt.Fprint(w, result)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/product/1", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "Product 1", string(body))
}
