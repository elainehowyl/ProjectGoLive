package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"net/http"
	//"golang.org/x/crypto/bcrypt"
)

func RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	errorsList := make(map[string]string)
	if r.Method == http.MethodPost {
		formValues := map[string]validator.StringInput{
			"customer_email": {
				Value:          r.FormValue("customer_email"),
				RequiredLength: 1,
			},
			"customer_username": {
				Value:          r.FormValue("customer_username"),
				RequiredLength: 6,
			},
			"customer_password": {
				Value:          r.FormValue("customer_password"),
				RequiredLength: 12,
			},
		}
		errorsList = validator.FormValidatorForString(formValues)
	}
	Tpl.ExecuteTemplate(w, "registercustomer.gohtml", errorsList)
}

func RegisterBOwner(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "registerbowner.gohtml", nil)
}
