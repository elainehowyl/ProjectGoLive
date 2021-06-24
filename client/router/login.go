package router

import "net/http"

func Login(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "login.gohtml", nil)
}
