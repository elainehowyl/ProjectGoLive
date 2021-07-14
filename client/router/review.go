package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"ProjectGoLiveElaine/ProjectGoLive/client/sanitizer"
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddReview(w http.ResponseWriter, r *http.Request) {
	var passed bool
	params := mux.Vars(r)
	listingIdS := params["listing_id"]
	listingId, _ := strconv.Atoi(listingIdS)
	errorsList := make(map[string]string)
	if r.Method == http.MethodPost {
		reviewerName := r.FormValue("reviewer_name")
		addReview := r.FormValue("add_review")
		formValues := map[string]validator.StringInput{
			"reviewer_name": {
				Value:          reviewerName,
				RequiredLength: 1,
			},
			"add_review": {
				Value:          addReview,
				RequiredLength: 8,
			},
		}
		errorsList, passed = validator.GeneralFormValidator(formValues)
		err := sanitizer.SimpleSanitization(reviewerName)
		if err != nil {
			errorsList["name_sanitization"] = err.Error()
		}
		err2 := sanitizer.SimpleSanitization(addReview)
		if err2 != nil {
			errorsList["review_sanitization"] = err.Error()
		}
		if passed && err == nil && err2 == nil {
			newReview := httpcontroller.AddReview{
				Id:         0,
				Name:       reviewerName,
				Review:     addReview,
				Listing_id: listingId,
			}
			c := make(chan error)
			go httpcontroller.ProcessAddReview(newReview, c)
			err3 := <-c
			if err3 != nil {
				errorsList["response_error"] = err.Error()
			} else {
				http.Redirect(w, r, "/listing/"+listingIdS, http.StatusSeeOther)
			}
		}
	}
	Tpl.ExecuteTemplate(w, "addreview.gohtml", errorsList)
}
