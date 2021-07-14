package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/httpcontroller"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func SetUp() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:49519)/proj_db")
	if err != nil {
		log.Panicln(err.Error())
		//panic(err.Error())
	} else {
		log.Println("Database opened")
		//fmt.Println("Database opened")
	}
	defer db.Close()
	r := mux.NewRouter()
	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.ProcessBOwnerLogin(w, r, db)
	})
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.ProcessBOwnerRegistration(w, r, db)
	})
	r.HandleFunc("/logout", httpcontroller.ProcessBOwnerLogout)
	r.HandleFunc("/listings", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.GetAllListings(w, r, db)
	})
	r.HandleFunc("/listing/{id}", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.GetOneListingPublic(w, r, db)
	})
	r.HandleFunc("/listing/{id}/review/add", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.ProcessAddReview(w, r, db)
	})
	r.HandleFunc("/profile/{email}", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.RetrieveBOwnerInfo(w, r, db)
	})
	r.HandleFunc("/{email}/listing/add", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.ProcessAddListing(w, r, db)
	})
	r.HandleFunc("/{email}/listing/{listing_id}/view", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.GetOneListing(w, r, db)
	})
	r.HandleFunc("/{email}/listing/{listing_id}/delete", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.ProcessDeleteListing(w, r, db)
	})
	r.HandleFunc("/{email}/listing/{listing_id}/item/add", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.ProcessAddItem(w, r, db)
	})
	r.HandleFunc("/{email}/listing/{listing_id}/item/{item_id}/delete", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.ProcessDeleteItem(w, r, db)
	})
	r.HandleFunc("/favicon.ico", http.NotFound)
	http.ListenAndServe(":5000", r)
}
