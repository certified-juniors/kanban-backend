package validation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// ValidatorInstance глобальная переменная для валидатора
var ValidatorInstance = validator.New()

// ValidateStruct валидирует структуру и возвращает ошибки валидации
func ValidateStruct(s interface{}) map[string]string {
	errMap := make(map[string]string)

	if err := ValidatorInstance.Struct(s); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errMap["general"] = "invalid validation error"
			return errMap
		}

		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			tag := err.Tag()
			value := err.Param()

			errorMessage := formatValidationError(fieldName, tag, value)
			errMap[fieldName] = errorMessage
		}
	}

	return errMap
}

// formatValidationError форматирует ошибку валидации
func formatValidationError(fieldName, tag, value string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("field '%s' is required", fieldName)
	case "min":
		return fmt.Sprintf("field '%s' must be at least %s", fieldName, value)
	case "max":
		return fmt.Sprintf("field '%s' must be at most %s", fieldName, value)
	case "email":
		return fmt.Sprintf("field '%s' must be a valid email address", fieldName)
	case "datetime":
		return fmt.Sprintf("field '%s' must be a valid date-time format: %s", fieldName, value)
	case "gtefield":
		return fmt.Sprintf("field '%s' must be greater than or equal to '%s'", fieldName, value)
	case "ltefield":
		return fmt.Sprintf("field '%s' must be less than or equal to '%s'", fieldName, value)
	case "oneof":
		return fmt.Sprintf("field '%s' must be one of the following values: %s", fieldName, value)
	default:
		return fmt.Sprintf("field '%s' failed validation with tag '%s'", fieldName, tag)
	}
}
