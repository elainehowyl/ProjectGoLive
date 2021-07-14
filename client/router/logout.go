package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/httpcontroller"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	c := make(chan error)
	go httpcontroller.ProcessBOwnerLogout(httpcontroller.CurrentUser.Email, c)
	err := <-c
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		httpcontroller.MyCookie = nil
		httpcontroller.CurrentUser = nil
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
