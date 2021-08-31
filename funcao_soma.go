package main // operação soma da calc em go

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func soma(x, y float64) float64 {
	return x + y
}

func somaHandler(
	
	teste1 http.ResponseWriter,
	teste2 *http.Request) {
	c := strings.Split(teste2.URL.Path, "/")
	num1, err := strconv.ParseFloat(c [2], 64)
	
	if err != nil {
		teste1.WriteHeader(400)
		
		return
	}

	num2, err := strconv.ParseFloat(c[3], 64)
	
	if err != nil {
		teste1.WriteHeader(400)
		
		return
	}

	result := soma(num1, num2)
	fmt.Fprintf(teste1, "%f", result)
	
} 