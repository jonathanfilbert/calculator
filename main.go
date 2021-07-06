package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

func calculatorHandler(w http.ResponseWriter, req *http.Request) {

	numA, _ := strconv.Atoi(req.URL.Query().Get("num_a"))
	numB, _ := strconv.Atoi(req.URL.Query().Get("num_b"))
	operator := req.URL.Query().Get("operator")
	result := 0

	switch operator {
	case "plus":
		result = numA + numB
	case "minus":
		result = numA - numB
	case "times":
		result = numA * numB
	case "divide":
		result = numA / numB
	case "power":
		result = int(math.Pow(float64(numA), float64(numB)))
	}

	fmt.Fprintf(w, fmt.Sprint(result))
}

func main() {
	// Declaration variables
	port := ":8989"

	http.HandleFunc("/calculator", calculatorHandler)

	fmt.Println("serving HTTP...")
	http.ListenAndServe(port, nil)

}
