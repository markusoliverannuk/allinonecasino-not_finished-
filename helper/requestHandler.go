package helper

import (
	"fmt"
	"net/http"
)

var SessionManager []string

// Handles all the incoming requests
func RequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	default:
		http.Redirect(w, r, "/homepage", http.StatusSeeOther)

	}
}

func HomePageExecutionHandler(w http.ResponseWriter, r *http.Request) {
	currUser := currentUser(r)
	loadContent(w, r, currUser)
}

func LoginPageExecutionHandler(w http.ResponseWriter, r *http.Request) {
	if loginPageTMPLT.Execute(w, nil) != nil {
		http.Error(w, "Error while rendering loginpage!", http.StatusInternalServerError)
	}
}

func RegisterPageExecutionHandler(w http.ResponseWriter, r *http.Request) {
	if registerPageTMPLT.Execute(w, nil) != nil {
		http.Error(w, "Error while rendering registerpage!", http.StatusInternalServerError)
	}
}

func LogoutAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	currUser := currentUser(r)
	fmt.Println("User ", currUser.Name, " has logged out.")
	http.SetCookie(w, CreateExpiredCookie())
	removeSession(currUser.Id)
	http.Redirect(w, r, "/loginpage", http.StatusSeeOther)

}

func RegisterAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	registerUserToDatabase(w, r)
	fmt.Println("New user has signed up!!")
}

func LoginAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	Login(w, r)
}

func Buy100Coins(w http.ResponseWriter, r *http.Request) {
	currUser := currentUser((r))
	buyCoins(w, r, currUser, 100)
	fmt.Println("100 coin purchase succesful for user:", currUser.Name)

}
func Buy50Coins(w http.ResponseWriter, r *http.Request) {
	currUser := currentUser((r))
	buyCoins(w, r, currUser, 50)
	fmt.Println("50 coin purchase succesful for user:", currUser.Name)

}
func Buy20Coins(w http.ResponseWriter, r *http.Request) {
	currUser := currentUser((r))
	buyCoins(w, r, currUser, 20)
	fmt.Println("20 coin purchase succesful for user:", currUser.Name)

}

func BetSubmitHandler(w http.ResponseWriter, r *http.Request) {

	currUser := currentUser(r)
	getUserBetValue(w, r, currUser)

}

func CoinflipHandler(w http.ResponseWriter, r *http.Request) {
	currUser := currentUser(r)
	loadCoinFlip(w, r, currUser)
}
func RouletteHandler(w http.ResponseWriter, r *http.Request) {
	currUser := currentUser(r)
	loadRoulette(w, r, currUser)
}
func WheelSpinHandler(w http.ResponseWriter, r *http.Request) {
	spinHandler(w, r)
}
func CoinflipLobbyHandler(w http.ResponseWriter, r *http.Request) {
	currUser := currentUser(r)
	loadCoinFlipLobby(w, r, currUser)
}
