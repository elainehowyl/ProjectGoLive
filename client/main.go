package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	Tpl = template.Must(template.ParseGlob("templates/*"))
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)
	r.HandleFunc("/login", Login)
	r.HandleFunc("/register/customer", RegisterCustomer)
	r.HandleFunc("/register/bowner", RegisterBowner)
	r.HandleFunc("/listing/id/view", ViewListing)
	r.HandleFunc("/listing/id/view/review", AddReview)
	r.HandleFunc("/bowner/profile", BownerProfile)
	r.HandleFunc("/bowner/listing/add", AddListing)
	r.HandleFunc("/bowner/listing/id/view", ListingDetails)
	r.HandleFunc("/bowner/listing/id/edit", EditItem)
	r.HandleFunc("/favicon.ico", http.NotFound)
	//http.ListenAndServeTLS(":8080", "./cert.pem", "./key.pem", r)
	http.ListenAndServe(":8080", r)
}

func Index(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "registercustomer.gohtml", nil)
}

func RegisterBowner(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "registerbowner.gohtml", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func BownerProfile(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "bownerprofile.gohtml", nil)
}

func AddListing(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "addnewlisting.gohtml", nil)
}

func ListingDetails(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "bownerlistingdetails.gohtml", nil)
}

func EditItem(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "edititem.gohtml", nil)
}

func ViewListing(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "viewlisting.gohtml", nil)
}

func AddReview(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "addreview.gohtml", nil)
}
