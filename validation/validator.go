package validation

import "io"

type Validator struct{}

type HTTPRequester func(string) (io.ReadCloser, error)

func NewValidator(url string, requester HTTPRequester) Validator {
	return Validator{}
}

func (v *Validator) InvalidCustomers() ([]byte, error) {
	return nil, nil
}
