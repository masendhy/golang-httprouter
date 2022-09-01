package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello Router")
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	//body, _ := io.ReadAll(recorder.Result().Body)
	//fmt.Println(string(body))

	response := recorder.Result()
	bytes, _ := io.ReadAll(response.Body)
	fmt.Println(string(bytes))
	//assert.Equal(t, "Hello Router", string(bytes))

}
