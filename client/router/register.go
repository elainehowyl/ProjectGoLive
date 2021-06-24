package router

import "net/http"

func RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "registercustomer.gohtml", nil)
}

func RegisterBOwner(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "registerbowner.gohtml", nil)
}
