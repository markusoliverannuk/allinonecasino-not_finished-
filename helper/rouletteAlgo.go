package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func rouletteAlgorithmRandomNumber() {
	// Seed the random number generator with a unique seed (usually based on the current time)
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer between 1 and 100
	randomNumber := rand.Intn(100) + 1

	fmt.Println("Random Number:", randomNumber)
}
