package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func UserPwd(field validator.FieldLevel) bool {
	// rexp at least 8, and at least with a-z, A-Z, and 0-9
	return regexp.MustCompile(`^[a-zA-Z\d]{8,}$`).MatchString(field.Field().String()) &&
		regexp.MustCompile(`[a-z]`).MatchString(field.Field().String()) &&
		regexp.MustCompile(`[A-Z]`).MatchString(field.Field().String()) &&
		regexp.MustCompile(`\d`).MatchString(field.Field().String())

}
