package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ParseToListing struct {
	ListingsAndItems *httpcontroller.Listing_Items
	ErrorsList       map[string]string
}

func ListingDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	listingIdS := params["listing_id"]
	listingId, _ := strconv.Atoi(listingIdS)
	var errorsList map[string]string
	//errorsList := make(map[string]string)
	var parse ParseToListing
	c := make(chan *httpcontroller.Listing_Items)
	go httpcontroller.GetOneListing(listingId, c)
	listingWItems := <-c
	parse.ListingsAndItems = listingWItems
	// if listingWItems == nil {
	// 	log.Println("There were some errors in retrieving your data")
	// 	http.Redirect(w, r, "/bowner/"+httpcontroller.CurrentUser.Email, http.StatusSeeOther)
	// 	return
	// }
	if r.Method == http.MethodPost {
		//errorsList := make(map[string]string)
		errorsList = AddItem(w, r, listingId)
		if errorsList["response_errors"] == "unauthorized" {
			log.Println("You are not authorized to submit this request")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		if errorsList != nil {
			log.Println("")
			parse.ErrorsList = errorsList
			//http.Redirect(w, r, "/bowner/"+httpcontroller.CurrentUser.Email+"/listing/"+listingIdS+"/view", http.StatusSeeOther)
		}
		//parse.ErrorsList = errorsList
	}
	//parse.ErrorsList = errorsList
	Tpl.ExecuteTemplate(w, "bownerlistingdetails.gohtml", parse)
}

func AddListing(w http.ResponseWriter, r *http.Request) {
	if httpcontroller.MyCookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	params := mux.Vars(r)
	errsList := make(map[string]string)
	var passed bool
	if r.Method == http.MethodPost {
		shopTitle := r.FormValue("shop_title")
		shopDescription := r.FormValue("shop_description")
		shopCategory := r.FormValue("shop_category")
		//categoryInt, _ := strconv.Atoi(category)
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
			"shop_category": {
				Value:          shopCategory,
				RequiredLength: 1,
			},
		}
		errsList, passed = validator.GeneralFormValidator(formValues)
		if passed {
			listing := httpcontroller.Listing{
				Id:              0,
				ShopTitle:       shopTitle,
				ShopDescription: shopDescription,
				IgURL:           igURL,
				FbURL:           fbURL,
				WebsiteURL:      websiteURL,
				Category: httpcontroller.Category{
					Id:    0,
					Title: shopCategory,
				},
			}
			c := make(chan error)
			go httpcontroller.ProcessAddListing(listing, c)
			err := <-c
			if err != nil {
				errsList["response_error"] = err.Error()
			} else {
				http.Redirect(w, r, "/bowner/"+params["email"], http.StatusSeeOther)
			}
		}
	}
	Tpl.ExecuteTemplate(w, "addnewlisting.gohtml", errsList)
}

func DeleteListing(w http.ResponseWriter, r *http.Request) {
	if httpcontroller.MyCookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	params := mux.Vars(r)
	listingId := params["listing_id"]
	c := make(chan error)
	go httpcontroller.ProcessDeleteListing(listingId, c)
	err := <-c
	if err != nil {
		log.Println(err)
		//response_error = err.Error()
	}
	http.Redirect(w, r, "/bowner/"+httpcontroller.CurrentUser.Email, http.StatusSeeOther)
	//Tpl.ExecuteTemplate(w, "deleteresult.gohtml", response_error)
}

func ListingForPublic(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	listingIdS := params["listing_id"]
	listingId, _ := strconv.Atoi(listingIdS)
	c := make(chan httpcontroller.Listing_Items_Reviews)
	go httpcontroller.GetListingWReviews(listingId, c)
	listingWReviews := <-c
	Tpl.ExecuteTemplate(w, "viewlisting.gohtml", listingWReviews)
}
