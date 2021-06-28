package httpcontroller

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/database"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CustomerCredentials struct {
	Email    string
	Password string
}

type CustomerDetails struct {
	Email    string
	Username string
	Password string
}

func ProcessCustomerRegistration(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			var customer CustomerDetails
			json.Unmarshal(reqBody, &customer)
			c := make(chan error)
			go database.AddCustomer(db, customer.Email, customer.Password, customer.Username, c)
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

func ProcessCustomerLogin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			var customer CustomerCredentials
			json.Unmarshal(reqBody, &customer)
			c := make(chan error)
			go database.VerifyCustomerIdentity(db, customer.Email, customer.Password, c)
			err = <-c
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Username or/and password does not match"))
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
