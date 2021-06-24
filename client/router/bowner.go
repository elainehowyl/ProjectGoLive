package router

import "net/http"

func MyProfile(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "bownerprofile.gohtml", nil)
}

func AddListing(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "addnewlisting.gohtml", nil)
}

func ViewMyListing(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "bownerlistingdetails.gohtml", nil)
}

func EditItem(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "edititem.gohtml", nil)
}
