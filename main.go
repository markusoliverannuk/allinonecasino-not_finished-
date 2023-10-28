package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"main.go/helper"
)

type BetData struct {
	BetAmount string `json:"betAmount"`
}

func main() {

	helper.Db = helper.InitializeDB("databaseSql/forumDatabase.db") //Initialize DB
	helper.CreateDbTable(helper.Db)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) //Links css file to html
	http.HandleFunc("/", helper.RequestHandler)
	http.HandleFunc("/homepage", helper.HomePageExecutionHandler)
	http.HandleFunc("/coinflip", helper.CoinflipHandler)
	http.HandleFunc("/submitBetAndWinning", helper.BetSubmitHandler)
	http.HandleFunc("/spin", helper.WheelSpinHandler)
	http.HandleFunc("/roulette", helper.RouletteHandler)
	http.HandleFunc("/coins100-form", helper.Buy100Coins)
	http.HandleFunc("/coins50-form", helper.Buy50Coins)
	http.HandleFunc("/coins20-form", helper.Buy20Coins)
	http.HandleFunc("/loginpage", helper.LoginPageExecutionHandler)
	http.HandleFunc("/login", helper.LoginAuthenticationHandler)
	http.HandleFunc("/logout", helper.LogoutAuthenticationHandler)
	http.HandleFunc("/register", helper.RegisterPageExecutionHandler)
	http.HandleFunc("/insertRegisterData", helper.RegisterAuthenticationHandler)
	http.HandleFunc("/coinfliplobby", helper.CoinflipLobbyHandler)
	log.Fatal(http.ListenAndServe(":4240", nil))
}
