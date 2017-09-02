package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/zjohl/dynamic-field-validation/validation"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Errorf("Expected format: ./path-to-binary endpoint-url")
		os.Exit(1)
	}
	url := arguments[1]

	validator, err := validation.GetValidator(url, makeHTTPRequest)
	if err != nil {
		panic(err)
	}

	bytes, err := validator.InvalidCustomers()
	if err != nil {
		panic(err)
	}

	fmt.Fprint(os.Stdout, string(bytes))

	os.Exit(0)
}

func makeHTTPRequest(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, err
}
