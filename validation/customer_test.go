package validation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/zjohl/dynamic-field-validation/validation"
)

var _ = Describe("Customer", func() {

	Describe("InvalidFields", func() {

		exampleValidations := map[string]Validation{
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
		}

		It("returns one field if it is invalidd", func() {
			customer := Customer{"id": 3, "name": "Bernardo", "email": "bernardo@interview.com", "age": 30, "country": "Brazil", "newsletter": "false"}

			Expect(customer.InvalidFields(exampleValidations)).To(BeEmpty())
		})

		It("returns nil if all fields are vali", func() {
			customer := Customer{"id": 2, "name": "Lily", "email": "lily@interview.com", "age": 24, "country": "China", "newsletter": false}

			Expect(customer.InvalidFields(exampleValidations)).To(Equal([]string{"name"}))
		})

		It("returns many fields if they are invalid", func() {
			customer := Customer{"id": 1, "name": "David", "country": "France", "age": nil}

			Expect(customer.InvalidFields(exampleValidations)).To(Equal([]string{"email", "newsletter"}))
		})
	})
})
