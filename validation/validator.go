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
	invalid := map[string][]string{}
	for _, customer := range v.Customers {
		fields := customer.InvalidFields(v.Validations)
		if len(fields) != 0 {
			id := customer["id"].(int)
			invalid[string(id)] = fields
		}
	}

	return json.Marshal(invalid)
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

func (v *Validator) UnmarshalJSON(data []byte) error {
	var intermDataStruc struct {
		Validations []map[string]Validation `json:"validations"`
		Customers   []Customer              `json:"customers"`
		Pagination  Pagination              `json:"pagination"`
	}
	err := json.Unmarshal(data, &intermDataStruc)
	if err != nil {
		return err
	}
	v.Customers = intermDataStruc.Customers
	v.Pagination = intermDataStruc.Pagination
	v.Validations = map[string]Validation{}
	for _, validation := range intermDataStruc.Validations {
		for key, val := range validation {
			v.Validations[key] = val
		}
	}
	return nil
}
