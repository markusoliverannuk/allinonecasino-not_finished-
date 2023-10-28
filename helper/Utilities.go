package helper

import (
	"strconv"
	"time"
)

// Generates a beautiful date format
func GetDateTime() string {
	currentDateTime := time.Now()
	day := strconv.Itoa(currentDateTime.Day())
	month := currentDateTime.Month().String()
	year := strconv.Itoa(currentDateTime.Year())

	return day + ". " + month + " " + year
}

// Removes session from sessionmanager
func removeSession(userId string) {
	for index, session := range SessionManager {
		if userId == session {
			SessionManager[index] = SessionManager[len(SessionManager)-1]
			SessionManager = SessionManager[:len(SessionManager)-1]
		}
	}
}

// Checks if session is still valid or no
func ifSessionExists(userId string) bool {
	for _, session := range SessionManager {
		if session == userId {
			return true
		}
	}
	return false
}
