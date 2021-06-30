package router

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var Tpl *template.Template

func init() {
	Tpl = template.Must(template.ParseGlob("templates/*"))
}

func RegisterRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/login", Login)
	r.HandleFunc("/register/customer", RegisterCustomer)
	r.HandleFunc("/register/bowner", RegisterBOwner)
	r.HandleFunc("/listing/id", ViewListing)
	r.HandleFunc("/listing/id/review", AddReview)
	r.HandleFunc("/bowner/email", MyProfile)
	r.HandleFunc("/bowner/email/listing/add", AddListing)
	r.HandleFunc("/bowner/email/listing/id/view", ViewMyListing)
	r.HandleFunc("/bowner/email/listing/{listing_id}/delete", DeleteListing)
	r.HandleFunc("/bowner/email/listing/{listing_id}/item/{item_id}/edit", EditItem)
	r.HandleFunc("/bowner/email/listing/{listing_id}/item/{item_id}/delete", DeleteItem)
	r.HandleFunc("/favicon.ico", http.NotFound)
	http.ListenAndServe(":8080", r)
	//http.ListenAndServeTLS(":8080", "./cert.pem", "./key.pem", r)
}
