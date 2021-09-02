package main

import (
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	assert.Equal(t, 3.0, soma(1, 2))
}

func TestSomaHandler(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://test/soma/1/2",
		nil)

	w := httptest.NewRecorder()
	somaHandler(w, req)
	resp := w.Result()

	assert.Equal(t, 200, resp.StatusCode)

	dados, _ := io.ReadAll(resp.Body)
	resultado, _ := strconv.ParseFloat(string(dados), 64)
	assert.Equal(t, 3., resultado)

}

func TestSomaHandler_ErrorParam2(t *testing.T) {
	
	req := httptest.NewRequest(
		"GET",
		"http://test/soma/1/OiOI",
		nil)

	w := httptest.NewRecorder()
	somaHandler(w, req)
	
	resp := w.Result()

	assert.Equal(t, 400, resp.StatusCode)

}

func TestSomaHandler_ErrorParam1(t *testing.T) {
	
	req := httptest.NewRequest(
		"GET",
		"http://test/soma/oooi/1",
		nil)

	w := httptest.NewRecorder()
	somaHandler(w, req)

	resp := w.Result()

	assert.Equal(t, 400, resp.StatusCode)
	
}