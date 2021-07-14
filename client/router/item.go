package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddItem(w http.ResponseWriter, r *http.Request, listingId int) map[string]string {
	listingIdS := strconv.Itoa(listingId)
	var passed bool
	var errorsList map[string]string
	itemName := r.FormValue("item_name")
	itemPrice := r.FormValue("item_price")
	itemDescription := r.FormValue("item_description")
	formValues := map[string]validator.StringInput{
		"item_name": {
			Value:          itemName,
			RequiredLength: 1,
		},
		"item_price": {
			Value:          itemPrice,
			RequiredLength: 1,
		},
		"item_description": {
			Value:          itemDescription,
			RequiredLength: 10,
		},
	}
	errorsList, passed = validator.GeneralFormValidator(formValues)
	itemPriceInFloat, err := strconv.ParseFloat(itemPrice, 64)
	if err != nil {
		errorsList["invalid_syntax"] = "Please insert only numbers"
	}
	if passed && err == nil {
		itemDetails := httpcontroller.Item{
			Id:          0,
			Name:        itemName,
			Price:       itemPriceInFloat,
			Description: itemDescription,
			ListingId:   listingId,
		}
		c := make(chan error)
		go httpcontroller.ProcessAddItem(itemDetails, c)
		err = <-c
		if err != nil {
			errorsList["response_error"] = err.Error()
		}
		http.Redirect(w, r, "/bowner/"+httpcontroller.CurrentUser.Email+"/listing/"+listingIdS+"/view", http.StatusSeeOther)
		return nil
	}
	return errorsList
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["item_id"]
	listingId := params["listing_id"]
	c := make(chan error)
	go httpcontroller.ProcessDeleteItem(itemId, listingId, c)
	err := <-c
	if err != nil {
		log.Println(err)
	} else {
		http.Redirect(w, r, "/bowner/"+httpcontroller.CurrentUser.Email+"/listing/"+listingId+"/view", http.StatusSeeOther)
	}
}
