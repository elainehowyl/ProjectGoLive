package httpcontroller

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/database"
	"ProjectGoLiveElaine/ProjectGoLive/server/model"
	"ProjectGoLiveElaine/ProjectGoLive/server/session"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
			var bOwner model.BOwner
			json.Unmarshal(reqBody, &bOwner)
			c := make(chan error)
			go database.AddBOwner(db, bOwner.Email, bOwner.Password, bOwner.Contact, c)
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
			var bOwner model.BOwnerCredentials
			json.Unmarshal(reqBody, &bOwner)
			c := make(chan error)
			go database.VerifyBOwnerIdentity(db, bOwner.Email, bOwner.Password, c)
			err = <-c
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Username or/and password does not match"))
				return
			}
			myToken := session.GenerateToken(bOwner.Email)
			myCookie := session.GenerateCookie(myToken)
			session.TrackSessions[myToken] = session.CookieInfo{
				Email:    bOwner.Email,
				MyCookie: myCookie,
			}
			http.SetCookie(w, myCookie)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func RetrieveBOwnerInfo(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	//params := mux.Vars(r)
	if r.Method == http.MethodGet {
		myCookie, _ := r.Cookie("myCookie")
		if _, ok := session.TrackSessions[myCookie.Value]; !ok {
			// check if there is an existing session
			err := "Verification of session failed. You are not authorized to visit this page."
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err))
			return
		}
		params := mux.Vars(r)
		if params["email"] != session.TrackSessions[myCookie.Value].Email {
			// check if current logged in user is mapped to the correct session
			err := "Verification of identity failed. You are not authorized to visit this page."
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err))
			return
		}
		// check if current cookie is expired
		// if expired, issue a new token and cookie
		expired := session.CookieExpired(session.TrackSessions[myCookie.Value].MyCookie)
		if expired {
			log.Println("Current session expired - deleting previous session and issuing new token.")
			delete(session.TrackSessions, myCookie.Value)
			myToken := session.GenerateToken(params["email"])
			myCookie = session.GenerateCookie(myToken)
			session.TrackSessions[myToken] = session.CookieInfo{
				Email:    params["email"],
				MyCookie: myCookie,
			}
		} else {
			myCookie = session.TrackSessions[myCookie.Value].MyCookie
		}
		// retrieve business owner profile
		c := make(chan *model.BOwnerData)
		go database.GetBOwnerData(db, params["email"], c)
		bownerDetails := <-c
		if bownerDetails == nil {
			http.SetCookie(w, myCookie)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("There were some errors in processing your request"))
			return
		}
		// retrieve business owner's listings
		c2 := make(chan []model.Listing)
		go database.GetListings(db, params["email"], c2)
		bownerListings := <-c2
		// if bownerListings == nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte("There were some errors in processing your request"))
		// 	return
		// }
		bownerDetails.Listings = bownerListings
		json.NewEncoder(w).Encode(&bownerDetails)
		http.SetCookie(w, myCookie)
	}
}

func ProcessBOwnerLogout(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == http.MethodPost {
			myCookie, _ := r.Cookie("myCookie")
			if _, ok := session.TrackSessions[myCookie.Value]; !ok {
				// check if there is an existing session
				err := "Verification of session failed. You are not authorized to visit this page."
				log.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(err))
				return
			}
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			var email string
			json.Unmarshal(reqBody, &email)
			if email != session.TrackSessions[myCookie.Value].Email {
				// check if current logged in user is mapped to the correct session
				err := "Verification of identity failed. You are not authorized to visit this page."
				log.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(err))
				return
			}
			// delete session
			delete(session.TrackSessions, myCookie.Value)
			log.Printf("Logout for %v is successful\n", email)
			w.WriteHeader(http.StatusNoContent)
		}
	}
}
