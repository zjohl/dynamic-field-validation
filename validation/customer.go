package validation

type Customer map[string]interface{}

func (c *Customer) InvalidFields(validations map[string]Validation) []string {
	invalid := []string{}
	for name, validation := range validations {
		if !validation.Validate(c, name) {
			invalid = append(invalid, name)
		}
	}
	return invalid
}
