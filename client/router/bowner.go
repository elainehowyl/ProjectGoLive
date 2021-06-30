package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"net/http"
)

func MyProfile(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "bownerprofile.gohtml", nil)
}

func AddListing(w http.ResponseWriter, r *http.Request) {
	errsList := make(map[string]string)
	var passed bool
	if r.Method == http.MethodPost {
		shopTitle := r.FormValue("shop_title")
		shopDescription := r.FormValue("shop_description")
		category := r.FormValue("shop_category")
		igURL := r.FormValue("ig_url")
		fbURL := r.FormValue("fb_url")
		websiteURL := r.FormValue("website_url")
		formValues := map[string]validator.StringInput{
			"shop_title": {
				Value:          shopTitle,
				RequiredLength: 1,
			},
			"shop_description": {
				Value:          shopDescription,
				RequiredLength: 15,
			},
		}
		errsList, passed = validator.GeneralFormValidator(formValues)
		if passed {
			listing := map[string]interface{}{
				"shop_title":       shopTitle,
				"shop_description": shopDescription,
				"ig_url":           igURL,
				"fb_url":           fbURL,
				"website_url":      websiteURL,
				"bowner_id":        0,
				"category_id":      category,
			}
			c := make(chan error)
			go httpcontroller.ProcessAddListing(listing, c)
			err := <-c
			if err != nil {
				errsList["response_error"] = err.Error()
			} else {
				http.Redirect(w, r, "/bowner/email", http.StatusSeeOther)
			}
		}
	}
	Tpl.ExecuteTemplate(w, "addnewlisting.gohtml", errsList)
}

func ViewMyListing(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "bownerlistingdetails.gohtml", nil)
}

func EditItem(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "edititem.gohtml", nil)
}
