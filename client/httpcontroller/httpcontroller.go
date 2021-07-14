package httpcontroller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

type MyCustomClaims struct {
	Authorized bool   `json:"authorized"`
	Email      string `json:"email"`
	jwt.StandardClaims
}

var (
	BaseURL  = "http://localhost:5000"
	MyCookie *http.Cookie
	Client   http.Client
)
