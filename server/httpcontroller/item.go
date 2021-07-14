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
	"strconv"

	"github.com/gorilla/mux"
)

func ProcessAddItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == http.MethodPost {
			myCookie, _ := r.Cookie("myCookie")
			if _, ok := session.TrackSessions[myCookie.Value]; !ok {
				// check if there is an existing session
				err := "Verification of session failed. You are not authorized to submit this request."
				log.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(err))
				return
			}
			params := mux.Vars(r)
			if params["email"] != session.TrackSessions[myCookie.Value].Email {
				// check if current logged in user is mapped to the correct session
				err := "Verification of identity failed. You are not authorized to submit this request."
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
			}
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.SetCookie(w, myCookie)
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			var item model.Item
			json.Unmarshal(reqBody, &item)
			c := make(chan error)
			go database.AddItem(db, item, c)
			err = <-c
			if err != nil {
				http.SetCookie(w, myCookie)
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			http.SetCookie(w, myCookie)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("New item created successfully"))
		}
	}
}

func ProcessDeleteItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodDelete {
		myCookie, _ := r.Cookie("myCookie")
		if _, ok := session.TrackSessions[myCookie.Value]; !ok {
			// check if there is an existing session
			err := "Verification of session failed. You are not authorized to submit this request."
			log.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err))
			return
		}
		params := mux.Vars(r)
		if params["email"] != session.TrackSessions[myCookie.Value].Email {
			// check if current logged in user is mapped to the correct session
			err := "Verification of identity failed. You are not authorized to submit this request."
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
		}
		itemIdS := params["item_id"]
		itemId, _ := strconv.Atoi(itemIdS)
		c := make(chan error)
		go database.DeleteSingleItem(db, itemId, c)
		err := <-c
		if err != nil {
			http.SetCookie(w, myCookie)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("There is an internal server error. Try again later."))
			return
		}
		http.SetCookie(w, myCookie)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Item is successfully deleted."))
	}
}
