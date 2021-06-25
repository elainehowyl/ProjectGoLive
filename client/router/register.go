package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	var passed bool
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
		errorsList, passed = validator.FormValidatorForString(formValues)
		if passed {
			fmt.Println("serving data to api server")
			encryptedpw, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("customer_password")), bcrypt.MinCost)
			newCustomer := map[string]string{
				"email":    r.FormValue("customer_email"),
				"username": r.FormValue("customer_username"),
				"password": string(encryptedpw[:]),
			}
			err := httpcontroller.AddCustomer(newCustomer)
			if err != nil {
				fmt.Printf("display the error: %v in template or something\n", err)
			}
		}
	}
	Tpl.ExecuteTemplate(w, "registercustomer.gohtml", errorsList)
}

func RegisterBOwner(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "registerbowner.gohtml", nil)
}
