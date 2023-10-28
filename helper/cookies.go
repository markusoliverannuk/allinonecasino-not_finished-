package helper

import (
	"net/http"
	"time"
)

func CreateCookie(userId string) *http.Cookie {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    userId,
		Path:     "/",
		MaxAge:   3600, // Expires in 60 seconds
		HttpOnly: true, // Cannot be accessed by JavaScript
	}
	return cookie
}

func CreateExpiredCookie() *http.Cookie {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true, // Cannot be accessed by JavaScript
	}
	return cookie
}
