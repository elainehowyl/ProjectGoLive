package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"fmt"
	"net/http"
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
		errorsList, passed = validator.GeneralFormValidator(formValues)
		if passed {
			//encryptedpw, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("customer_password")), bcrypt.MinCost)
			loginCredentials := map[string]string{
				"email":    r.FormValue("login_email"),
				"password": r.FormValue("login_password"),
			}
			fmt.Println("ROLE FROM FORM: ", r.FormValue("role"))
			if r.FormValue("role") == "customer" {
				c := make(chan error)
				go httpcontroller.ProcessCustomerLogin(loginCredentials, c)
				err := <-c
				if err != nil {
					fmt.Println("Response Error on login:", err)
					errorsList["response_error"] = err.Error()
				} else {
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			}
			if r.FormValue("role") == "bowner" {
				c := make(chan error)
				go httpcontroller.ProcessBOwnerLogin(loginCredentials, c)
				err := <-c
				if err != nil {
					fmt.Println("Response Error on login", err)
					errorsList["response_error"] = err.Error()
				} else {
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			}
		}
	}
	Tpl.ExecuteTemplate(w, "login.gohtml", errorsList)
}
