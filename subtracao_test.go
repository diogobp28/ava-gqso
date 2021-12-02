package main

import (
	"testing"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

	func TestSubtracao_CERTO(t *testing.T) {
		assert.Equal(t, 1., subtracao(2, 1))
	}
	
	func TestSubtracaoHandler_CERTO(t *testing.T) {
		req := httptest.NewRequest(
			"GET",
			"http://test.com/subtracao/2/1",
			nil)
		x := httptest.NewRecorder()
		subHandler(x, req)
	
		resp := x.Result()
		body, _ := io.ReadAll(resp.Body)
		result, _ := strconv.ParseFloat(string(body), 64)
		
		assert.Equal(t, 1., result)

	}
	
	func TestSubtracaoHandler_Error2Param(t *testing.T) {
		
		req := httptest.NewRequest(
			"GET",
			"http://test.com/subtracao/1/4d",
			nil)
		
			x := httptest.NewRecorder()
		subHandler(w, req)
		resp := x.Result()
		assert.Equal(t, 400, resp.StatusCode)
	}
	
	func TestSubtracaoHandler_Error1Param(t *testing.T) {
		
		req := httptest.NewRequest(
			"GET",
			"http://test.com/sub/4d/1",
			nil)
		
			x := httptest.NewRecorder()
		subHandler(x, req)
		resp := x.Result()
		
		assert.Equal(t, 400, resp.StatusCode)
	}
