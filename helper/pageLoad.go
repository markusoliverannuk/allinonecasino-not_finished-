package helper

import (
	"fmt"
	"net/http"
)

func loadCoinFlip(w http.ResponseWriter, r *http.Request, currUser User) {
	var hp HomePage
	hp.CurrentUser = currUser

	e := coinflipPageTMPLT.Execute(w, hp)
	if e != nil {
		fmt.Println(e)
		http.Error(w, "Error while rendering coinflip!", http.StatusInternalServerError)
	}
}
func loadRoulette(w http.ResponseWriter, r *http.Request, currUser User) {
	var hp HomePage
	hp.CurrentUser = currUser

	e := roulettePageTMPLT.Execute(w, hp)
	if e != nil {
		fmt.Println(e)
		http.Error(w, "Error while rendering roulette!", http.StatusInternalServerError)
	}
}
func loadContent(w http.ResponseWriter, r *http.Request, currUser User) {
	var hp HomePage
	hp.CurrentUser = currUser

	e := homePageTMPLT.Execute(w, hp)
	if e != nil {
		fmt.Println(e)
		http.Error(w, "Error while rendering homepage/crash!", http.StatusInternalServerError)
	}
}
func loadCoinFlipLobby(w http.ResponseWriter, r *http.Request, currUser User) {
	var hp HomePage
	hp.CurrentUser = currUser

	e := coinfliplobbyPageTMPLT.Execute(w, hp)
	if e != nil {
		fmt.Println(e)
		http.Error(w, "Error while rendering coinflip lobby!", http.StatusInternalServerError)
	}
}
