package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func subtracao(a, b float64) float64 {
	return a - b
}

func subtracaoHandler(
	x http.ResponseWriter,
	y *http.Request) {
	path := strings.Split(r.URL.Path, "/")func TestSub_OK(t *testing.T) {
		assert.Equal(t, 1., sub(2, 1))
	}
	
	func TestSubHandler_OK(t *testing.T) {
		req := httptest.NewRequest(
			"GET",
			"http://test.com/sub/2/1",
			nil)
		w := httptest.NewRecorder()
		subHandler(w, req)
	
		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)
		result, _ := strconv.ParseFloat(string(body), 64)
		assert.Equal(t, 1., result)
	}
	
	func TestSubHandler_Error2Param(t *testing.T) {
		req := httptest.NewRequest(
			"GET",
			"http://test.com/sub/1/4d",
			nil)
		w := httptest.NewRecorder()
		subHandler(w, req)
		resp := w.Result()
		assert.Equal(t, 400, resp.StatusCode)
	}
	
	func TestSubHandler_Error1Param(t *testing.T) {
		req := httptest.NewRequest(
			"GET",
			"http://test.com/sub/4d/1",
			nil)
		w := httptest.NewRecorder()
		subHandler(w, req)
		resp := w.Result()
		assert.Equal(t, 400, resp.StatusCode)
	}
	op1, err := strconv.ParseFloat(path[2], 64)
	
	if err != nil {
		x.WriteHeader(400)
		return
	}

	op2, err := strconv.ParseFloat(path[3], 64)
	if err != nil {
		x.WriteHeader(400)
		return
	}
	result := subtracao(op1, op2)
	fmt.Fprintf(x, "%f", result)
}