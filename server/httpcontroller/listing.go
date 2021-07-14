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

func GetOneListing(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodGet {
		params := mux.Vars(r)
		listingIdS := params["listing_id"]
		listingId, _ := strconv.Atoi(listingIdS)
		// get one listing
		c := make(chan *model.Listing_Items)
		go database.GetOneListing(db, listingId, c)
		listing := <-c
		if listing == nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("There were some errors in processing your request"))
			return
		}
		// get related items for current listing id
		c2 := make(chan []model.Item)
		go database.GetItems(db, listingId, c2)
		items := <-c2
		// if items == nil {
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	w.Write([]byte("There were some errors in processing your request"))
		// 	return
		// }
		listing.Items = items
		json.NewEncoder(w).Encode(&listing)
	}
}

func GetOneListingPublic(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodGet {
		params := mux.Vars(r)
		listindIdS := params["id"]
		listingId, _ := strconv.Atoi(listindIdS)
		// get general listing and details
		c := make(chan *model.Listing_Items_Reviews)
		go database.GetOneListingWReviews(db, listingId, c)
		listings := <-c
		// get items
		c2 := make(chan []model.Item)
		go database.GetItems(db, listingId, c2)
		items := <-c2
		listings.Items = items
		// get reviews
		c3 := make(chan []model.Review)
		go database.GetReviews(db, listingId, c3)
		reviews := <-c3
		listings.Reviews = reviews
		json.NewEncoder(w).Encode(&listings)
	}
}

func GetAllListings(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodGet {
		c := make(chan []model.Listing)
		go database.GetAllListings(db, c)
		listings := <-c
		json.NewEncoder(w).Encode(&listings)
	}
}

func ProcessAddListing(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Header.Get("Content-type") == "application/json" {
		if r.Method == "POST" {
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
				w.WriteHeader(http.StatusUnprocessableEntity)
				w.Write([]byte("Unable to process request"))
				return
			}
			var listing model.Listing
			json.Unmarshal(reqBody, &listing)
			c := make(chan int)
			go database.AddCategory(db, listing.Category, c)
			//go database.AddListing(db, listing, c)
			id := <-c
			if id == 0 {
				// retrieve category id
				go database.GetCategory(db, listing.Category.Title, c)
				id = <-c
			}
			listing.Category.Id = id
			c2 := make(chan error)
			go database.AddListing(db, listing, params["email"], c2)
			err = <-c2
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("There is an internal server error. Try again later."))
				return
			}
			http.SetCookie(w, myCookie)
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("New listing created successfully"))
		}
	}
}

func ProcessDeleteListing(w http.ResponseWriter, r *http.Request, db *sql.DB) {
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
		listingIdS := params["listing_id"]
		listingId, _ := strconv.Atoi(listingIdS)
		// delete all items linked to this listing (diassociate foreign keys)
		c := make(chan error)
		go database.DeleteItem(db, listingId, c)
		err := <-c
		if err != nil {
			http.SetCookie(w, myCookie)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("There is an internal server error. Try again later."))
			return
		}
		// delete listing
		go database.DeleteListing(db, listingId, c)
		err = <-c
		if err != nil {
			http.SetCookie(w, myCookie)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("There is an internal server error. Try again later."))
			return
		}
		http.SetCookie(w, myCookie)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Listing is successfully deleted."))
	}
}
