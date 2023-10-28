package helper

import (
	"fmt"
	"net/http"
)

func buyCoins(w http.ResponseWriter, r *http.Request, currUser User, amount int) error {
	var err error
	http.Redirect(w, r, "/homepage", http.StatusSeeOther)
	// Retrieve the user's current balance from the database
	var currentBalance int

	err = Db.QueryRow("SELECT balance FROM users WHERE id = ?", currUser.Id).Scan(&currentBalance)
	if err != nil {
		fmt.Println("hey")
		return err

	}

	newBalance := currentBalance + amount

	_, err = Db.Exec("UPDATE users SET balance = ? WHERE id = ?", newBalance, currUser.Id)
	if err != nil {
		return err
	}

	http.Redirect(w, r, "/homepage", http.StatusSeeOther)
	return nil
}
