package httpcontroller

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/database"
	"ProjectGoLiveElaine/ProjectGoLive/server/session"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ProcessBOwnerRegistration(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			var bOwner database.BOwner
			json.Unmarshal(reqBody, &bOwner)
			c := make(chan error)
			go database.AddBOwner(db, bOwner.Id, bOwner.Email, bOwner.Password, bOwner.Contact, c)
			err = <-c
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Email was already registered"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Account created successfully"))
		}
	}
}

func ProcessBOwnerLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			var bOwner database.BOwnerCredentials
			json.Unmarshal(reqBody, &bOwner)
			c := make(chan error)
			go database.VerifyBOwnerIdentity(db, bOwner.Email, bOwner.Password, c)
			err = <-c
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Username or/and password does not match"))
				return
			}
			myToken := session.GenerateToken(bOwner.Email, "bowner")
			myCookie := session.GenerateCookie(myToken)
			http.SetCookie(w, myCookie)
			//w.Header().Add("myCookie", myToken)
			w.WriteHeader(http.StatusOK)
			//myToken := GenerateToken()
			//http.SetCookie(w, myToken)
		}
	}
}
