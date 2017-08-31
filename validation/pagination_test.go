package validation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/zjohl/dynamic-field-validation/validation"
)

var _ = Describe("Pagination", func() {

	Describe("LastPage", func() {
		It("returns true for only one item", func() {
			pagination := &Pagination{
				CurrentPage: 1,
				PerPage:     12,
				Total:       1,
			}
			Expect(pagination.LastPage()).To(BeTrue())
		})

		It("returns true for only one page", func() {
			pagination := &Pagination{
				CurrentPage: 1,
				PerPage:     5,
				Total:       5,
			}
			Expect(pagination.LastPage()).To(BeTrue())
		})

		It("returns false if it is not the last page", func() {
			pagination := &Pagination{
				CurrentPage: 2,
				PerPage:     7,
				Total:       18,
			}
			Expect(pagination.LastPage()).To(BeFalse())
		})

		It("returns true if it is the last page", func() {
			pagination := &Pagination{
				CurrentPage: 3,
				PerPage:     12,
				Total:       30,
			}
			Expect(pagination.LastPage()).To(BeTrue())
		})
	})
})
