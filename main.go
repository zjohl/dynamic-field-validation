package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Errorf("Expected format: ./path-to-binary endpoint-url")
	}

}

func makeHTTPRequest(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp.Body, err
}
