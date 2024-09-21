package handlers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func NewValidator() *CustomValidator {
	validator := validator.New()
	customValidator := &CustomValidator{Validator: validator}

	return customValidator
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

// ValidateAndReturnErrors performs validation and returns detailed error messages
func (cv *CustomValidator) ValidateAndReturnErrors(i interface{}) error {
	if err := cv.Validate(i); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			errorMessages := make(map[string]string)
			for _, validationErr := range validationErrors {
				field := validationErr.Field()
				tag := validationErr.Tag()

				// Customize error messages based on the validation tag
				switch tag {
				case "required":
					errorMessages[field] = "This field is required"
				case "min":
					errorMessages[field] = "This field is too short"
				case "max":
					errorMessages[field] = "This field is too long"
				case "email":
					errorMessages[field] = "Invalid email address"
				default:
					errorMessages[field] = validationErr.Error()
				}
			}

			// Return an echo.HTTPError to halt execution
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMessages})
		}

		// Return a generic validation failure error
		return echo.NewHTTPError(http.StatusBadRequest, map[string]string{"error": "Validation failed"})
	}
	return nil
}
