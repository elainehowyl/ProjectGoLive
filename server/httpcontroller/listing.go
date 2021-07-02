package httpcontroller

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/database"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ProcessAddListing(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			var listing database.Listing
			json.Unmarshal(reqBody, &listing)
			c := make(chan error)
			go database.AddListing(db, listing, c)
			err = <-c
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("There is an internal server error. Try again later."))
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("New listing created successfully"))
		}
	}
}
