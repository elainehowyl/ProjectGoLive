package validator

import (
	"errors"
	"fmt"
)

type StringInput struct {
	Value          string
	RequiredLength int
}

func LengthValidator(input string, requiredLength int) error {
	if len(input) == 0 {
		return errors.New("Required field is empty")
	}
	if len(input) < requiredLength {
		return fmt.Errorf("Please enter at least %v characters", requiredLength)
	}
	return nil
}

func FormValidatorForString(formValues map[string]StringInput) (map[string]string, bool) {
	errorsList := map[string]string{}
	errCount := 0
	for key, value := range formValues {
		err := LengthValidator(value.Value, value.RequiredLength)
		if err != nil {
			errorsList[key] = err.Error()
			errCount++
		} else {
			errorsList[key] = ""
		}
	}
	if errCount > 0 {
		return errorsList, false
	}
	return errorsList, true
}
