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
			newCustomer := httpcontroller.Customer{
				Id:       0,
				Email:    r.FormValue("customer_email"),
				Username: r.FormValue("customer_username"),
				Password: r.FormValue("customer_password"),
			}
			c := make(chan error)
			go httpcontroller.ProcessCustomerRegistration(newCustomer, c)
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
			newBowner := httpcontroller.BOwner{
				ID:       0,
				Email:    r.FormValue("bowner_email"),
				Contact:  r.FormValue("bowner_contact"),
				Password: r.FormValue("bowner_password"),
			}
			c := make(chan error)
			go httpcontroller.ProcessBOwnerRegistration(newBowner, c)
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
