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

func TestParamMoreThanTwo(t *testing.T) {
	router := httprouter.New()
	router.GET("/room/:roomId/student/:studentId", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		roomId := p.ByName("roomId")
		studentId := p.ByName("studentId")

		result := "Room " + roomId + " Student " + studentId
		fmt.Fprint(w, result)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/room/2/student/209", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "Room 2 Student 209", string(body))
}

func TestCatchAll(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		image := p.ByName("image")
		result := "Result: " + image
		fmt.Fprint(w, result)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/images/assets/img/SBY.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "Result: /assets/img/SBY.png", string(body))
}
