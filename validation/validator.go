package validation

import (
	"encoding/json"
	"io"
)

type HTTPRequester func(string) (io.ReadCloser, error)

func GetValidator(url string, requester HTTPRequester) (*Validator, error) {
	validator, err := makeRequest(url, requester)
	if err != nil {
		return nil, err
	}

	for !validator.Pagination.LastPage() {
		suffix := "?page=" + string(validator.Pagination.CurrentPage+1)
		resp, err := makeRequest(url+suffix, requester)
		if err != nil {
			return nil, err
		}

		validator.Pagination = resp.Pagination
		validator.Customers = append(validator.Customers, resp.Customers...)
	}
	return validator, err
}

func makeRequest(url string, requester HTTPRequester) (*Validator, error) {
	reader, err := requester(url)
	if err != nil {
		return nil, err
	}

	resp := &Validator{}
	err = json.NewDecoder(reader).Decode(resp)
	return resp, err
}

func (v *Validator) InvalidCustomers() ([]byte, error) {
	////allValid := true
	//for _, customer := range v.Customers {
	//	//allValid = allValid && customer.IsValid(v.Validations)
	//}
	////TODO: collect values and marshal
	return nil, nil
}

type Validator struct {
	Validations map[string]Validation `json:"validations"`
	Customers   []Customer            `json:"customers"`
	Pagination  Pagination            `json:"pagination"`
}

type Validation struct {
	Required bool    `json:"required"`
	Type     string  `json:"type"`
	Length   *Length `json:"length"`
}

type Length struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

func (v *Validation) Validate(customer *Customer, name string) bool {
	val, ok := (*customer)[name]
	if !ok {
		return !v.Required
	}

	if v.Length != nil {
		length := len(val.(string))
		if v.Length.Min != 0 && length < v.Length.Min {
			return false
		}
		if v.Length.Max != 0 && length > v.Length.Max {
			return false
		}
	}

	if v.Type != "" {
		switch val.(type) {
		case string:
			return v.Type == "string"
		case bool:
			return v.Type == "bool"
		case int:
			return v.Type == "number"
		}
	}

	return true
}
