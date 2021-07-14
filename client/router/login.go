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
			loginCredentials := map[string]string{
				"email":    r.FormValue("login_email"),
				"password": r.FormValue("login_password"),
			}
			c := make(chan error)
			go httpcontroller.ProcessBOwnerLogin(loginCredentials, c)
			err := <-c
			if err != nil {
				fmt.Println("Response Error on login", err)
				errorsList["response_error"] = err.Error()
			} else {
				http.Redirect(w, r, "/bowner/"+loginCredentials["email"], http.StatusSeeOther)
			}
		}
	}
	Tpl.ExecuteTemplate(w, "login.gohtml", errorsList)
}
