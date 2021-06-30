package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"ProjectGoLiveElaine/ProjectGoLive/client/sanitizer"
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"fmt"
	"net/http"
)

func RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	var passed bool
	var passed2 bool
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
			"password": {
				Value:          r.FormValue("customer_password"),
				RequiredLength: 12,
			},
		}
		errorsList, passed = validator.FormValidatorForRegistration(formValues)
		errorsList, passed2 = sanitizer.RegistrationSanitization(formValues, errorsList)
		if passed && passed2 {
			newCustomer := map[string]string{
				"email":    r.FormValue("customer_email"),
				"username": r.FormValue("customer_username"),
				"password": r.FormValue("customer_password"),
			}
			c := make(chan error)
			go httpcontroller.AddCustomer(newCustomer, c)
			err := <-c
			if err != nil {
				fmt.Printf("display the error: %v in template or something\n", err)
				errorsList["response_error"] = err.Error()
			} else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
		}
	}
	Tpl.ExecuteTemplate(w, "registercustomer.gohtml", errorsList)
}

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
			newBowner := map[string]string{
				"email":    r.FormValue("bowner_email"),
				"contact":  r.FormValue("bowner_contact"),
				"password": r.FormValue("bowner_password"),
			}
			c := make(chan error)
			go httpcontroller.AddBOwner(newBowner, c)
			err := <-c
			if err != nil {
				fmt.Printf("display the error: %v in template or something\n", err)
				errorsList["response_error"] = err.Error()
			} else {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
			}
		}
	}
	Tpl.ExecuteTemplate(w, "registerbowner.gohtml", errorsList)
}
