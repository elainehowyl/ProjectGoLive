package router

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type PassToTemplate struct {
	APIData    struct{}
	ErrorsList map[string]string
}

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("templates/*"))
}

func RegisterRoutes() {
	log.Println("Client starting...")
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/listing/{listing_id}", ListingForPublic)
	r.HandleFunc("/listing/{listing_id}/review/add", AddReview)
	r.HandleFunc("/login", Login)
	r.HandleFunc("/logout", Logout)
	r.HandleFunc("/register", RegisterBOwner)
	//r.HandleFunc("/listing/id", ViewListing)
	//r.HandleFunc("/listing/id/review", AddReview)
	r.HandleFunc("/bowner/{email}", MyProfile)
	//r.HandleFunc("/bowner/{email}/logout", Logout)
	r.HandleFunc("/bowner/{email}/listing/add", AddListing)
	r.HandleFunc("/bowner/{email}/listing/{listing_id}/view", ListingDetails)
	r.HandleFunc("/bowner/{email}/listing/{listing_id}/delete", DeleteListing)
	r.HandleFunc("/bowner/listing/{listing_id}/item/{item_id}/edit", EditItem)
	r.HandleFunc("/bowner/listing/{listing_id}/item/{item_id}/delete", DeleteItem)
	r.HandleFunc("/favicon.ico", http.NotFound)
	//http.ListenAndServe(":8080", r)
	http.ListenAndServeTLS(":8080", "./cert.pem", "./key.pem", r)
}
