package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"ProjectGoLiveElaine/ProjectGoLive/client/sanitizer"
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"net/http"
)

func ViewListing(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "viewlisting.gohtml", nil)
}

func AddReview(w http.ResponseWriter, r *http.Request) {
	formValues := make(map[string]string)
	errMsgs := make(map[string]string)
	if r.Method == http.MethodPost {
		review := r.FormValue("add_review")
		err := validator.LengthValidator(review, 15)
		if err != nil {
			errMsgs["validation"] = err.Error()
		}
		err2 := sanitizer.SimpleSanitization(review)
		if err != nil {
			errMsgs["sanitization"] = err.Error()
		}
		if err == nil && err2 == nil {
			formValues["add_review"] = review
			formValues["customer_id"] = "?"
			formValues["listing_id"] = "?"
			c := make(chan error)
			go httpcontroller.ProcessAddReview(formValues, c)
			err = <-c
			if err != nil {
				errMsgs["response_error"] = err.Error()
			} else {
				http.Redirect(w, r, "/listing/id", http.StatusSeeOther)
			}
		}
	}
	Tpl.ExecuteTemplate(w, "addreview.gohtml", errMsgs)
}
