package session

import (
	"ProjectGoLiveElaine/ProjectGoLive/client/envfile"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CookieInfo struct {
	Email    string
	MyCookie *http.Cookie
}

var (
	//TrackSessions = make(map[string]string)
	TrackSessions = make(map[string]CookieInfo)
)

func GenerateToken(email string) string {
	mySigningKey := envfile.RetrieveEnv("MY_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"email":      email,
	})
	signedToken, _ := token.SignedString([]byte(mySigningKey))
	return signedToken
}

func GenerateCookie(signedToken string) *http.Cookie {
	//expiration := time.Now().Add(15 * time.Minute)
	expiration := time.Now().Add(10 * time.Second)
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

///delete(user.MapSessions, myCookie.Value)
