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
	path := strings.Split(r.URL.Path, "/")
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