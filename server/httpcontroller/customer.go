package httpcontroller

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/database"
	"ProjectGoLiveElaine/ProjectGoLive/server/session"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// type CustomerCredentials struct {
// 	Email    string
// 	Password string
// }

// type CustomerDetails struct {
// 	Email    string
// 	Username string
// 	Password string
// }

func ProcessCustomerRegistration(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			var customer database.Customer
			json.Unmarshal(reqBody, &customer)
			c := make(chan error)
			go database.AddCustomer(db, customer.Id, customer.Email, customer.Password, customer.Username, c)
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
			var customer database.CustomerCredentials
			json.Unmarshal(reqBody, &customer)
			c := make(chan error)
			go database.VerifyCustomerIdentity(db, customer.Email, customer.Password, c)
			err = <-c
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Username or/and password does not match"))
				return
			}
			myToken := session.GenerateToken(customer.Email, "customer")
			myCookie := session.GenerateCookie(myToken)
			http.SetCookie(w, myCookie)
			w.WriteHeader(http.StatusOK)
		}
	}
}
