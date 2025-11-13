package helpers

import (
	"strings"

	"github.com/go-playground/validator/v10"
)


var (
	validate *validator.Validate
)

func GetValidator() *validator.Validate {
	if validate == nil {
		validate = validator.New()
	}
	return validate
}

func ParseValidationErrors(err error) map[string]string {
	errs := make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {

		for _, e := range validationErrors {
			field := e.Field()

			switch e.Tag() {
				case "required":
					errs[strings.ToLower(field)] = field + " harus diisi bro"
				case "min" :
					errs[strings.ToLower(field)] = field + " harus setidaknya " + e.Param() + " karakter"
				default:
					errs[strings.ToLower(field)] = field + " aturan validasi tidak valid"
			}
		}
	} else {
		errs["error"] = err.Error()
	}

	return errs
}