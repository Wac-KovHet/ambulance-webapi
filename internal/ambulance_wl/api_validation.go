package ambulance_wl

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func formatValidationError(err error) map[string]string {
	errs := make(map[string]string)
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.ToLower(err.Field())
		errs[field] = err.Tag()
	}
	return errs
}