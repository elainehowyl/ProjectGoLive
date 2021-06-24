package router

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	Tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
