package router

import "net/http"

func ViewListing(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "viewlisting.gohtml", nil)
}

func AddReview(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "addreview.gohtml", nil)
}
