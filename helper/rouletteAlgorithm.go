package helper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type BetData struct {
	BetAmount      string  `json:"betAmount"`
	SelectedColor  string  `json:"selectedColor"`
	WinningDegrees float64 `json:"winningDegrees"` // Define the WinningDegrees field
}

func calcSpin() float64 {
	int1 := rand.Intn(380)
	fmt.Println(int1)

	degreeRange := 360.0 / float64(38)

	degree1 := degreeRange * float64(int1)
	var degree float64
	degree = degree1 / 10

	degree += 1800
	fmt.Println(WinningColor(degree))
	return degree

}

var colorThatWon string

func WinningColor(degree float64) string {

	switch result := winningNumber(degree); result {
	case 2, 35, 4, 33, 6, 31, 8, 29, 10, 13, 24, 15, 22, 17, 20, 11, 26, 28:
		colorThatWon = "Black"

	case 14, 23, 16, 21, 18, 19, 12, 25, 27, 1, 36, 3, 34, 5, 32, 7, 30, 9:
		colorThatWon = "Purple"

	}

	return colorThatWon
}

func didUserWin2(selectedColor string, winnerColor string) bool {
	return selectedColor == winnerColor
}

func getUserBetValue(w http.ResponseWriter, r *http.Request, currUser User) (int, string) {
	var betAmount string
	var selectedColor string
	var betAmountInt int
	var winnerDegree float64

	if r.Method == "POST" {

		var requestData BetData

		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Failed to parse request", http.StatusBadRequest)

		}

		winnerDegree = requestData.WinningDegrees
		selectedColor = requestData.SelectedColor

		winnerNumber := winningNumber(winnerDegree)
		winnerColor := WinningColor(winnerNumber)
		fmt.Println("Number that won was: ", winnerNumber)

		if selectedColor == "black-button" {
			selectedColor = "Black"
		}
		if selectedColor == "green-button" {
			selectedColor = "Green"
		}
		if selectedColor == "purple-button" {
			selectedColor = "Purple"
		}
		fmt.Println("Winning color:", winnerColor)
		fmt.Println("User selected color:", selectedColor)
		betAmount = requestData.BetAmount
		betAmountInt, err := strconv.Atoi(betAmount)
		if err != nil {
			http.Error(w, "Invalid betAmount format", http.StatusBadRequest)

		}
		if selectedColor == winnerColor {
			fmt.Println("User won!")
		} else if selectedColor != winnerColor {
			fmt.Println("User lost!")
		}
		fmt.Println("Bet amount:", betAmountInt)
		var currentBalance int
		err = Db.QueryRow("SELECT balance FROM users WHERE id = ?", currUser.Id).Scan(&currentBalance)

		newBalance := currentBalance - betAmountInt
		if selectedColor == winnerColor {
			newBalance += 2 * betAmountInt
		}
		_, err = Db.Exec("UPDATE users SET balance = ? WHERE id = ?", newBalance, currUser.Id)

		if err != nil {
			fmt.Println("error")
		}

		response := struct {
			Message string `json:"message"`
		}{
			Message: "Bet submitted successfully",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	return betAmountInt, selectedColor
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func spinHandler(w http.ResponseWriter, r *http.Request) {
	//  calc degrees that won
	winningDegrees := calcSpin()

	// log test win degrees

	// send win degrees js
	w.Write([]byte(fmt.Sprintf("%.3f", winningDegrees)))
}

func winningNumber(number float64) float64 {
	var numberout float64
	constant := 360.0 / 38.0
	number -= 1800
	if number > 0 && number <= constant {
		numberout = 2
	}
	if number > constant && number <= 2*constant {
		numberout = 14
	}
	if number > 2*constant && number <= 3*constant {
		numberout = 35
	}
	if number > 3*constant && number <= 4*constant {
		numberout = 23
	}
	if number > 4*constant && number <= 5*constant {
		numberout = 4
	}
	if number > 5*constant && number <= 6*constant {
		numberout = 16
	}
	if number > 6*constant && number <= 7*constant {
		numberout = 33
	}
	if number > 7*constant && number <= 8*constant {
		numberout = 21
	}
	if number > 8*constant && number <= 9*constant {
		numberout = 6
	}
	if number > 9*constant && number <= 10*constant {
		numberout = 18
	}
	if number > 10*constant && number <= 11*constant {
		numberout = 31
	}
	if number > 11*constant && number <= 12*constant {
		numberout = 19
	}
	if number > 12*constant && number <= 13*constant {
		numberout = 8
	}
	if number > 13*constant && number <= 14*constant {
		numberout = 12
	}
	if number > 14*constant && number <= 15*constant {
		numberout = 29
	}
	if number > 15*constant && number <= 16*constant {
		numberout = 25
	}
	if number > 16*constant && number <= 17*constant {
		numberout = 10
	}
	if number > 17*constant && number <= 18*constant {
		numberout = 27
	}
	if number > 18*constant && number <= 19*constant {
		numberout = 00
	}
	if number > 19*constant && number <= 20*constant {
		numberout = 1
	}
	if number > 20*constant && number <= 21*constant {
		numberout = 13
	}
	if number > 21*constant && number <= 22*constant {
		numberout = 36
	}
	if number > 22*constant && number <= 23*constant {
		numberout = 24
	}
	if number > 23*constant && number <= 24*constant {
		numberout = 3
	}
	if number > 24*constant && number <= 25*constant {
		numberout = 15
	}
	if number > 25*constant && number <= 26*constant {
		numberout = 34
	}
	if number > 26*constant && number <= 27*constant {
		numberout = 22
	}
	if number > 27*constant && number <= 28*constant {
		numberout = 5
	}
	if number > 28*constant && number <= 29*constant {
		numberout = 17
	}
	if number > 29*constant && number <= 30*constant {
		numberout = 32
	}
	if number > 30*constant && number <= 31*constant {
		numberout = 20
	}
	if number > 31*constant && number <= 32*constant {
		numberout = 7
	}
	if number > 32*constant && number <= 33*constant {
		numberout = 11
	}
	if number > 33*constant && number <= 34*constant {
		numberout = 30
	}
	if number > 34*constant && number <= 35*constant {
		numberout = 26
	}
	if number > 35*constant && number <= 36*constant {
		numberout = 9
	}
	if number > 36*constant && number <= 37*constant {
		numberout = 28
	}
	if number > 37*constant && number <= 38*constant {
		numberout = 00
	}

	return numberout

}
