package validation_test

import (
	"io"

	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/zjohl/dynamic-field-validation/validation"
)

var _ = Describe("Validator", func() {

	It("identifies the invalid customers", func() {
		mockHTTPRequester := func(string) (io.ReadCloser, error) {
			return os.Open("../fixtures/example_api_response.json")
		}

		validator := NewValidator("some-url", mockHTTPRequester)

		expectedResponse, err := ioutil.ReadFile("../fixtures/expected_response.json")
		Expect(err).NotTo(HaveOccurred())
		Expect(validator.InvalidCustomers()).To(Equal(expectedResponse))
	})
})
