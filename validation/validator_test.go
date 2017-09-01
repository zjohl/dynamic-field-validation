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
		Validations: map[string]Validation{
			"name": {
				Required: true,
				Type:     "string",
				Length: &Length{
					Min: 5,
				},
			},
			"email": {
				Required: true,
			},
			"age": {
				Type:     "number",
				Required: false,
			},
			"newsletter": {
				Required: true,
			},
		},
		Customers: []Customer{
			{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true},
			{"id": 2.0, "name": "Lily", "email": "lily@interview.com", "age": 24, "country": "China", "newsletter": false},
			{"id": 3.0, "name": "Bernardo", "email": "bernardo@interview.com", "age": 30, "country": "Brazil", "newsletter": "false"},
			{"id": 4.0, "name": "Gabriel", "email": "gabriel@interview.com", "age": 28, "country": "Canada", "newsletter": true},
			{"id": 5.0, "name": "Alex", "email": "alex@interview.com", "age": 29, "country": "United States", "newsletter": true},
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

	Describe("Validate", func() {
		Context("when a field is nil", func() {
			It("returns true if unrequired", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true}

				validation := Validation{Required: false, Type: "number"}

				Expect(validation.Validate(customer, "some-nil-field")).To(BeTrue())
			})

			It("returns false if required", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true}

				validation := Validation{Required: true, Type: "number"}

				Expect(validation.Validate(customer, "some-nil-field")).To(BeFalse())
			})
		})

		Context("when a field is a boolean", func() {
			It("returns false if type is not boolean", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true}

				validation := Validation{Required: true, Type: "number"}

				Expect(validation.Validate(customer, "newsletter")).To(BeFalse())
			})

			It("returns true if value is false", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": false}

				validation := Validation{Required: true, Type: "boolean"}

				Expect(validation.Validate(customer, "newsletter")).To(BeTrue())
			})
		})

		Context("when a field is a number", func() {
			It("returns false if type is not number", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true}

				validation := Validation{Required: true, Type: "string"}

				Expect(validation.Validate(customer, "id")).To(BeFalse())
			})

			It("returns true if type is number", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true}

				validation := Validation{Required: true, Type: "number"}

				Expect(validation.Validate(customer, "id")).To(BeTrue())
			})
		})

		Context("when a field is a string", func() {
			It("returns false if type is not string", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true}

				validation := Validation{Required: true, Type: "boolean"}

				Expect(validation.Validate(customer, "name")).To(BeFalse())
			})

			It("returns true if type is string", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true}

				validation := Validation{Required: true, Type: "string"}

				Expect(validation.Validate(customer, "name")).To(BeTrue())
			})

			It("returns true if length is greater than minimum", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true}

				validation := Validation{Required: true, Length: &Length{Min: 4}}

				Expect(validation.Validate(customer, "name")).To(BeTrue())
			})

			It("returns false if length is less than minimum", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true}

				validation := Validation{Required: true, Length: &Length{Min: 6}}

				Expect(validation.Validate(customer, "name")).To(BeFalse())
			})

			It("returns true if length is less than maximum", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true}

				validation := Validation{Required: true, Length: &Length{Max: 6}}

				Expect(validation.Validate(customer, "name")).To(BeTrue())
			})

			It("returns false if length is greater than maximum", func() {
				customer := &Customer{"id": 1.0, "name": "David", "email": "david@interview.com", "country": "France", "newsletter": true}

				validation := Validation{Required: true, Length: &Length{Max: 4}}

				Expect(validation.Validate(customer, "name")).To(BeFalse())
			})
		})
	})

	Describe("InvalidCustomers", func() {
		It("identifies the invalid customers", func() {
			expectedResponse, err := ioutil.ReadFile("../fixtures/expected_response.json")
			Expect(err).NotTo(HaveOccurred())

			bytes, err := exampleValidator.InvalidCustomers()
			Expect(err).NotTo(HaveOccurred())
			Expect(bytes).To(MatchJSON(expectedResponse))
		})
	})
})
