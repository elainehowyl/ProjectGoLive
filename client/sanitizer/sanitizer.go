package sanitizer

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"

	"github.com/microcosm-cc/bluemonday"
)

func RegistrationSanitization(formValues map[string]validator.StringInput, errorsList map[string]string) (map[string]string, bool) {
	p := bluemonday.StrictPolicy()
	count := 0
	for key, value := range formValues {
		afterSanitize := p.Sanitize(value.Value)
		if afterSanitize != value.Value {
			errorsList[key] = "Please ensure that field doesn't contain other special characters aside from: !@#$%^*"
			count++
		}
	}
	if count > 0 {
		return errorsList, false
	}
	return errorsList, true
}
