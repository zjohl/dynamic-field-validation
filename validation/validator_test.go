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

	exampleValidator := &Validator{
		Validations: []map[string]Validation{
			{
				"name": {
					Required: true,
					Type:     "string",
					Length: Length{
						Min: 5,
					},
				},
			},
			{
				"email": {
					Required: true,
				},
			},
			{
				"age": {
					Type:     "number",
					Required: false,
				},
			},
			{
				"newsletter": {
					Required: true,
				},
			},
		},
		Customers: []Customer{
			{"id": 1, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true},
			{"id": 2, "name": "Lily", "email": "lily@interview.com", "age": 24, "country": "China", "newsletter": false},
			{"id": 3, "name": "Bernardo", "email": "bernardo@interview.com", "age": 30, "country": "Brazil", "newsletter": "false"},
			{"id": 4, "name": "Gabriel", "email": "gabriel@interview.com", "age": 28, "country": "Canada", "newsletter": true},
			{"id": 5, "name": "Alex", "email": "alex@interview.com", "age": 29, "country": "United States", "newsletter": true},
		},
		Pagination: Pagination{
			CurrentPage: 1,
			PerPage:     5,
			Total:       5,
		},
	}

	Describe("GetValidator", func() {
		It("creates a validator from a json endpoint", func() {

			mockHTTPRequester := func(string) (io.ReadCloser, error) {
				return os.Open("../fixtures/example_api_response.json")
			}

			validator, err := GetValidator("some-url", mockHTTPRequester)
			Expect(err).NotTo(HaveOccurred())
			Expect(validator.Pagination).To(Equal(exampleValidator.Pagination))
		})

		It("makes a request for each page", func() {
			var count int

			mockHTTPRequester := func(string) (io.ReadCloser, error) {
				switch count {
				case 0:
					count++
					return os.Open("../fixtures/example_api_response_page_1.json")
				case 1:
					count++
					return os.Open("../fixtures/example_api_response_page_2.json")
				case 2:
					count++
					return os.Open("../fixtures/example_api_response_page_3.json")
				case 3:
					count++
					return os.Open("../fixtures/example_api_response_page_4.json")
				default:
					Fail("Unexpected number of http requests made")
					return nil, nil
				}
			}

			validator, err := GetValidator("some-url", mockHTTPRequester)
			Expect(err).NotTo(HaveOccurred())

			Expect(count).To(Equal(4))
			Expect(validator.Customers).To(HaveLen(16))
		})
	})

	Describe("InvalidCustomers", func() {
		It("identifies the invalid customers", func() {
			expectedResponse, err := ioutil.ReadFile("../fixtures/expected_response.json")
			Expect(err).NotTo(HaveOccurred())
			Expect(exampleValidator.InvalidCustomers()).To(Equal(expectedResponse))
		})
	})
})
