package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var passed bool
	errorsList := make(map[string]string)
	if r.Method == http.MethodPost {
		formValues := map[string]validator.StringInput{
			"login_email": {
				Value:          r.FormValue("login_email"),
				RequiredLength: 1,
			},
			"login_password": {
				Value:          r.FormValue("login_password"),
				RequiredLength: 1,
			},
		}
		errorsList, passed = validator.FormValidatorForString(formValues)
		if passed {
			encryptedpw, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("customer_password")), bcrypt.MinCost)
			loginCredentials := map[string]string{
				"email":    r.FormValue("login_email"),
				"password": string(encryptedpw[:]),
			}
			if r.FormValue("role") == "customer" {
			}
			if r.FormValue("role") == "bowner" {
			}
		}
	}
	Tpl.ExecuteTemplate(w, "login.gohtml", errorsList)
}
