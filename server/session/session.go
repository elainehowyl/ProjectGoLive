package session

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

var (
	TrackSessions = make(map[string]string)
)

func GenerateToken(email, role string) string {
	MY_SECRET_KEY, _ := uuid.NewV4()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  role,
	})
	signedToken, _ := token.SignedString([]byte(MY_SECRET_KEY.String()))
	return signedToken
}

func GenerateCookie(signedToken string) *http.Cookie {
	expiration := time.Now().Add(15 * time.Minute)
	myCookie := &http.Cookie{
		Name:     "myCookie",
		Value:    signedToken,
		Expires:  expiration,
		HttpOnly: true,
		Path:     "/",
		Domain:   "127.0.0.1",
		Secure:   true,
	}
	return myCookie
}

func CookieExpired(cookie *http.Cookie) bool {
	now := time.Now().Unix()
	cookieExp := cookie.Expires.Unix()
	if now >= cookieExp {
		return true
	}
	return false
}
