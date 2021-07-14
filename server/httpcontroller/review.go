package httpcontroller

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/database"
	"ProjectGoLiveElaine/ProjectGoLive/server/model"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ProcessAddReview(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == http.MethodPost {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			var review model.AddReview
			json.Unmarshal(reqBody, &review)
			c := make(chan error)
			go database.AddReview(db, review, c)
			err = <-c
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("There is an internal server error. Try again later."))
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("New review created successfully"))
		}
	}
}
