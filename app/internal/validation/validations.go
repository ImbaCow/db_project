package validation

import (
	ozzoValidation "github.com/go-ozzo/ozzo-validation"
)

// RequiredIf ...
func RequiredIf(cond bool) ozzoValidation.RuleFunc {
	return func(value interface{}) error {
		if cond {
			return ozzoValidation.Validate(value, ozzoValidation.Required)
		}
		return nil
	}
}
