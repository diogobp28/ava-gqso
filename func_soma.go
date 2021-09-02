package main

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

	w http.ResponseWriter,
	r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	t1, err := strconv.ParseFloat(partes[2], 64)
	
	if err != nil {
		w.WriteHeader(400)
		return
	}

	t2, err := strconv.ParseFloat(partes[3], 64)
	
	if err != nil {
		w.WriteHeader(400)
		return
	}

	resultado := soma(t1, t2)
	fmt.Fprintf(w, "%f", resultado)
}