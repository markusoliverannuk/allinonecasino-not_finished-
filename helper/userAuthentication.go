package helper

import (
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Login func check all the things with log in
func Login(w http.ResponseWriter, r *http.Request) {
	//If we get post method
	if r.Method == http.MethodPost {
		var passWordToCompare string
		var userId string
		Db.QueryRow("SELECT id, password FROM users WHERE name = ? OR email = ?", r.FormValue("name"), r.FormValue("name")).Scan(&userId, &passWordToCompare)
		//Check user session
		if ifSessionExists(userId) {
			loginPageTMPLT.Execute(w, "User is already logged in!")
			return
		}
		//Check whether user exists or no check email and username
		if bcrypt.CompareHashAndPassword([]byte(passWordToCompare), []byte(r.FormValue("password"))) == nil {
			http.SetCookie(w, CreateCookie(userId))
			SessionManager = append(SessionManager, userId)
			http.Redirect(w, r, "/homepage", http.StatusSeeOther)
			fmt.Println("User ", userId, " has logged in.")
		} else {
			loginPageTMPLT.Execute(w, "Invalid username/email or password!")
		}
	}
}

// Register func check and insert all neccessary things about registering user
func registerUserToDatabase(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//Check wheter the gmail and username is registered or no
		var email string
		Db.QueryRow("SELECT email FROM users WHERE email = ?", r.FormValue("email")).Scan(&email)
		if email != "" {
			registerPageTMPLT.Execute(w, "User already registered!")
			return
		}
		fmt.Println(r.FormValue("email"))
		Db.QueryRow("SELECT email FROM users WHERE name = ?", r.FormValue("name")).Scan(&email)
		if email != "" {
			registerPageTMPLT.Execute(w, "User already registered!")
			return
		}
		//Generate password for user
		hash, err1 := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
		if err1 != nil {
			fmt.Println("Something went wrong with encryption of password!")
			return
		}
		//Generate UUID for each user
		uniqId, err := uuid.NewV4()
		if err != nil {
			fmt.Println("Error occured while generating uuid for user")
			return
		}
		//construct user
		u := User{
			Name:         r.FormValue("name"),
			Email:        r.FormValue("email"),
			PassWordHash: string(hash),
			Id:           uniqId.String(),
			Role:         1,
			Balance:      0,
		}
		//Insert user data to database
		_, errForInsertingUserData := Db.Exec("INSERT INTO users (id, name, email, password, role, balance) VALUES (?, ?, ?, ?, ?, ?)", u.Id, u.Name, u.Email, u.PassWordHash, u.Role, u.Balance)
		if errForInsertingUserData != nil {
			fmt.Println("Something went wrong with inserting user data in database!")
		}
		http.Redirect(w, r, "/loginpage", http.StatusSeeOther)
	}
}

func getUserById(userId string) (User, error) {
	var u User
	err := Db.QueryRow("SELECT id, name, email, password, role, balance FROM users WHERE id = ?", userId).Scan(&u.Id, &u.Name, &u.Email, &u.PassWordHash, &u.Role, &u.Balance)
	if err != nil {
		return u, err
	}
	return u, nil
}
