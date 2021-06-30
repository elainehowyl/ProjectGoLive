package validator

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

type StringInput struct {
	Value          string
	RequiredLength int
}

func SpecialCharValidator(password string) bool {
	var special bool
	specialChar := "!@#$%^*"
	for _, char := range specialChar {
		result := strings.Contains(password, string(char))
		if result {
			special = true
		}
	}
	return special
}

func UpperCaseValidator(password string) bool {
	var upper bool
	for _, char := range password {
		result := unicode.IsUpper(char)
		if result {
			upper = true
		}

	}
	return upper
}

func PasswordValidator(password string) error {
	specialChar := SpecialCharValidator(password)
	upperCase := UpperCaseValidator(password)
	if !specialChar || !upperCase {
		return errors.New("Please ensure that your password contain at least ONE special character(!@#$%^*) and ONE uppercase character")
	}
	return nil
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

func FormValidatorForLogin(formValues map[string]StringInput) (map[string]string, bool) {
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

func FormValidatorForRegistration(formValues map[string]StringInput) (map[string]string, bool) {
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
	passwordPassed := PasswordValidator(formValues["password"].Value)
	if passwordPassed != nil {
		errorsList["password_char"] = passwordPassed.Error()
		errCount++
	}
	if errCount > 0 {
		return errorsList, false
	}
	return errorsList, true
}
