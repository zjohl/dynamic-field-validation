package validation

import (
	"encoding/json"
	"io"
)

type HTTPRequester func(string) (io.ReadCloser, error)

func GetValidator(url string, requester HTTPRequester) (*Validator, error) {
	reader, err := requester(url)
	if err != nil {
		return nil, err
	}
	var resp *Validator
	err = json.NewDecoder(reader).Decode(resp)
	return resp, err
}

func (v *Validator) InvalidCustomers() ([]byte, error) {
	allValid := true
	for _, customer := range v.Customers {
		allValid = allValid && customer.IsValid(v.Validations)
	}
	//TODO: collect values and marshal
	return nil, nil
}

type Validator struct {
	Validations []map[string]Validation
	Customers   []Customer
	Pagination  Pagination
}

type Validation struct {
	Name     string
	Required bool
	Type     string // TODO: make enum
	Length   Length
}

type Length struct {
	Min int
	Max int
}

type Customer map[string]interface{}

func (c *Customer) IsValid([]map[string]Validation) bool {
	return false
}
