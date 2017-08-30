package validation

import (
	"encoding/json"
	"io"
)

type Validator struct {
	requester HTTPRequester
	url       string
}

type HTTPRequester func(string) (io.ReadCloser, error)

func NewValidator(url string, requester HTTPRequester) Validator {
	return Validator{
		requester: requester,
		url:       url,
	}
}

func (v *Validator) InvalidCustomers() ([]byte, error) {
	_, err := v.getFirstPage()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (v *Validator) getFirstPage() (*apiResponse, error) {
	reader, err := v.requester(v.url)
	if err != nil {
		return nil, err
	}
	var resp *apiResponse
	err = json.NewDecoder(reader).Decode(resp)
	return resp, err
}

type apiResponse struct {
	Validations []Validation
	Customers   []Customer
	Pagination  Pagination
}

type Validation struct {
	Required bool
	Type     string // TODO: make enum
	Length   Length
}

type Length struct {
	Min int
	Max int
}

type Customer map[string]interface{}

type Pagination struct {
	CurrentPage int
	PerPage     int
	Total       int
}
