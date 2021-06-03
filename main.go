package main

import (
	"fmt"
	"net/http"
)

func calculatorHandler(w http.ResponseWriter, req *http.Request) {

	numA := req.URL.Query().Get("num_a")
	numB := req.URL.Query().Get("num_b")
	operator := req.URL.Query().Get("operator")

	fmt.Fprintf(w, fmt.Sprint(numA, numB, operator))
}

func main() {

	http.HandleFunc("/calculator", calculatorHandler)

	fmt.Println("serving HTTP...")
	http.ListenAndServe(":8989", nil)
}
