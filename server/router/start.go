package router

import (
	"ProjectGoLiveElaine/ProjectGoLive/server/httpcontroller"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func SetUp() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:53361)/proj_db")
	if err != nil {
		log.Panicln(err.Error())
		//panic(err.Error())
	} else {
		log.Println("Database opened")
		//fmt.Println("Database opened")
	}
	defer db.Close()
	r := mux.NewRouter()
	r.HandleFunc("/bowner/login", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.ProcessBOwnerLogin(w, r, db)
	})
	r.HandleFunc("/customer/login", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.ProcessCustomerLogin(w, r, db)
	})
	r.HandleFunc("/bowner/register", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.ProcessBOwnerRegistration(w, r, db)
	})
	r.HandleFunc("/customer/register", func(w http.ResponseWriter, r *http.Request) {
		httpcontroller.ProcessCustomerRegistration(w, r, db)
	})
	r.HandleFunc("/favicon.ico", http.NotFound)
	http.ListenAndServe(":5000", r)
}
