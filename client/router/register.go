package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"ProjectGoLiveElaine/ProjectGoLive/client/sanitizer"
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"net/http"
)

func RegisterBOwner(w http.ResponseWriter, r *http.Request) {
	var passed bool
	var passed2 bool
	errorsList := make(map[string]string)
	if r.Method == http.MethodPost {
		formValues := map[string]validator.StringInput{
			"bowner_email": {
				Value:          r.FormValue("bowner_email"),
				RequiredLength: 1,
			},
			"bowner_contact": {
				Value:          r.FormValue("bowner_contact"),
				RequiredLength: 8,
			},
			"password": {
				Value:          r.FormValue("bowner_password"),
				RequiredLength: 12,
			},
		}
		errorsList, passed = validator.FormValidatorForRegistration(formValues)
		errorsList, passed2 = sanitizer.RegistrationSanitization(formValues, errorsList)
		if passed && passed2 {
			newBowner := httpcontroller.BOwner{
				Email:    r.FormValue("bowner_email"),
				Contact:  r.FormValue("bowner_contact"),
				Password: r.FormValue("bowner_password"),
			}
			c := make(chan error)
			go httpcontroller.ProcessBOwnerRegistration(newBowner, c)
			err := <-c
			if err != nil {
				errorsList["response_error"] = err.Error()
			} else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
		}
	}
	Tpl.ExecuteTemplate(w, "registerbowner.gohtml", errorsList)
}
