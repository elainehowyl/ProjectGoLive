package httpcontroller

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/database"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type BOwnerCredentials struct {
	Email    string
	Password string
}

type BOwnerDetails struct {
	Email    string
	Password string
	Contact  string
}

func ProcessBOwnerRegistration(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			var bOwner BOwnerDetails
			json.Unmarshal(reqBody, &bOwner)
			c := make(chan error)
			go database.AddBOwner(db, bOwner.Email, bOwner.Password, bOwner.Contact, c)
			err = <-c
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Registration failed. Email was already registered."))
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
				w.Write([]byte("unable to process request"))
				return
			}
			var bOwner BOwnerCredentials
			json.Unmarshal(reqBody, &bOwner)
			c := make(chan error)
			go database.VerifyBOwnerIdentity(db, bOwner.Email, bOwner.Password, c)
			err = <-c
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("username or/and password does not match"))
				return
			}
			//http.SetCookie(w, myCookie)
			//w.Header().Add("myCookie", signedToken)
			w.WriteHeader(http.StatusOK)
			//myToken := GenerateToken()
			//http.SetCookie(w, myToken)
		}
	}
}
