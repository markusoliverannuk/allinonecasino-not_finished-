package helper

import (
	"fmt"
	"net/http"
)

func currentUser(r *http.Request) User {
	var currUserId string
	currCookie, err := r.Cookie("session")
	if err == nil {
		errForCookie := Db.QueryRow("SELECT id FROM users WHERE id = ?", currCookie.Value).Scan(&currUserId)
		if errForCookie != nil {
			fmt.Println(errForCookie)
		}
	}
	currUser, _ := getUserById(currUserId)
	return currUser
}
