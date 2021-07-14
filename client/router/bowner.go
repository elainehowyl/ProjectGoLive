package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/envfile"
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"ProjectGoLiveElaine/ProjectGoLive/client/validator"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func MyProfile(w http.ResponseWriter, r *http.Request) {
	if httpcontroller.MyCookie == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	var email string
	signingKey := envfile.RetrieveEnv("MY_SECRET_KEY")
	token, err := jwt.ParseWithClaims(httpcontroller.MyCookie.Value, &httpcontroller.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if claims, ok := token.Claims.(*httpcontroller.MyCustomClaims); ok && token.Valid {
		email = claims.Email
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	c := make(chan *httpcontroller.Profile)
	go httpcontroller.GetBOwnerData(email, c)
	profile := <-c
	if profile == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
		//fmt.Println("not authorized to view page.")
	}
	Tpl.ExecuteTemplate(w, "bownerprofile.gohtml", profile)
}

func EditItem(w http.ResponseWriter, r *http.Request) {
	errorsList := make(map[string]string)
	params := mux.Vars(r)
	itemId := params["item_id"]
	listingId := params["listing_id"]
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
				"listing_id":       listingId,
			}
			c := make(chan error)
			go httpcontroller.ProcessUpdateItem(itemId, itemDetails, c)
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
