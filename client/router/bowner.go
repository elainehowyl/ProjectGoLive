package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func DeleteListing(w http.ResponseWriter, r *http.Request) {
	var response_error string
	params := mux.Vars(r)
	listingId := params["listing_id"]
	c := make(chan error)
	go httpcontroller.ProcessDeleteListing(listingId, c)
	err := <-c
	if err != nil {
		response_error = err.Error()
	} else {
		http.Redirect(w, r, "/bowner/email", http.StatusSeeOther)
	}
	Tpl.ExecuteTemplate(w, "deleteresult.gohtml", response_error)
}

func ViewMyListing(w http.ResponseWriter, r *http.Request) {
	errorsList := make(map[string]string)
	if r.Method == http.MethodPost {
		itemName := r.FormValue("item_name")
		itemPrice := r.FormValue("item_price")
		itemDescription := r.FormValue("item_description")
		err := validator.LengthValidator(itemName, 1)
		if err != nil {
			errorsList["item_name"] = err.Error()
		}
		err2 := validator.LengthValidator(itemPrice, 1)
		if err2 != nil {
			errorsList["item_price"] = err.Error()
		}
		itemPriceInFloat, err3 := strconv.ParseFloat(itemPrice, 64)
		if err3 != nil {
			errorsList["invalid_syntax"] = "Please insert only numbers"
		}
		if err == nil && err2 == nil && err3 == nil {
			itemDetails := map[string]interface{}{
				"item_name":        itemName,
				"item_price":       itemPriceInFloat * 100,
				"item_description": itemDescription,
				"listing_id":       0,
			}
			c := make(chan error)
			go httpcontroller.ProcessAddItem(itemDetails, c)
			err = <-c
			if err != nil {
				errorsList["response_error"] = err.Error()
			} else {
				http.Redirect(w, r, "/bowner/email/listing/id/view", http.StatusSeeOther)
			}
		}
	}
	Tpl.ExecuteTemplate(w, "bownerlistingdetails.gohtml", errorsList)
}

func EditItem(w http.ResponseWriter, r *http.Request) {
	errorsList := make(map[string]string)
	if r.Method == http.MethodPost {
		itemName := r.FormValue("edit_item_name")
		itemPrice := r.FormValue("edit_item_price")
		itemDescription := r.FormValue("edit_item_description")
		err := validator.LengthValidator(itemName, 1)
		if err != nil {
			errorsList["edit_item_name"] = err.Error()
		}
		err2 := validator.LengthValidator(itemPrice, 1)
		if err2 != nil {
			errorsList["edit_item_price"] = err.Error()
		}
		itemPriceInFloat, err3 := strconv.ParseFloat(itemPrice, 64)
		if err3 != nil {
			errorsList["invalid_syntax"] = "Please insert only numbers"
		}
		if err == nil && err2 == nil && err3 == nil {
			itemDetails := map[string]interface{}{
				"item_name":        itemName,
				"item_price":       itemPriceInFloat * 100,
				"item_description": itemDescription,
				"listing_id":       0,
			}
			c := make(chan error)
			go httpcontroller.ProcessAddItem(itemDetails, c)
			err = <-c
			if err != nil {
				errorsList["response_error"] = err.Error()
			} else {
				http.Redirect(w, r, "/bowner/email/listing/id/view", http.StatusSeeOther)
			}
		}
	}
	Tpl.ExecuteTemplate(w, "edititem.gohtml", errorsList)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	var response_error string
	params := mux.Vars(r)
	itemId := params["item_id"]
	c := make(chan error)
	go httpcontroller.ProcessDeleteItem(itemId, c)
	err := <-c
	if err != nil {
		response_error = err.Error()
	} else {
		http.Redirect(w, r, "/bowner/email", http.StatusSeeOther)
	}
	Tpl.ExecuteTemplate(w, "deleteresult.gohtml", response_error)
}
