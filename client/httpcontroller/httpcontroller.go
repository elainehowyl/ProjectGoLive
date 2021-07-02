package httpcontroller

import "net/http"

var (
	BaseURL  = "http://localhost:5000"
	MyCookie *http.Cookie
	Client   http.Client
)
