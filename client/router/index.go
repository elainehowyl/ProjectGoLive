package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	c := make(chan []httpcontroller.Listing)
	go httpcontroller.GetAllListings(c)
	listings := <-c
	Tpl.ExecuteTemplate(w, "index.gohtml", listings)
}
