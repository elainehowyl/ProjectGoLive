package httpcontroller

import "net/http"

var (
	baseURL  = "http://localhost:5000"
	myCookie *http.Cookie
	client   http.Client
)
