package helper

import "fmt"

func mart() {

	var number float64
	var numberout float64
	if number >= 0 && number <= 4.86486486 {
		numberout = 0
	}
	if number > 4.86486486 && number <= 14.59459459459 {
		numberout = 2
	}
	if number > 14.59459459459 && number <= 24.32432432 {
		numberout = 14
	}
	if number > 24.32432432 && number <= 34.05405405 {
		numberout = 35
	}
	if number > 34.05405405 && number <= 43.78378378 {
		numberout = 23
	}
	if number > 43.78378378 && number <= 53.51351351 {
		numberout = 4
	}
	if number > 53.51351351 && number <= 63.24324324 {
		numberout = 16
	}
	if number > 63.24324324 && number <= 72.97297297 {
		numberout = 33
	}
	if number > 72.97297297 && number <= 82.7027027 {
		numberout = 21
	}
	if number > 82.7027027 && number <= 92.43243243 {
		numberout = 6
	}
	if number > 92.43243243 && number <= 102.1621622 {
		numberout = 18
	}
	if number > 102.1621622 && number <= 111.8918919 {
		numberout = 31
	}
	if number > 111.8918919 && number <= 121.6216216 {
		numberout = 19
	}
	if number > 121.6216216 && number <= 131.3513513 {
		numberout = 8
	}
	if number > 131.3513513 && number <= 141.0810811 {
		numberout = 12
	}
	if number > 141.0810811 && number <= 150.8108108 {
		numberout = 29
	}
	if number > 150.8108108 && number <= 160.5405405 {
		numberout = 25
	}
	if number > 160.5405405 && number <= 170.2702703 {
		numberout = 10
	}
	if number > 170.2702703 && number <= 180 {
		numberout = 27
	}
	if number > 180 && number <= 189.7297297 {
		numberout = 0
	}
	if number > 189.7297297 && number <= 199.4594595 {
		numberout = 1
	}
	if number > 199.4594595 && number <= 209.1891892 {
		numberout = 13
	}
	if number > 209.1891892 && number <= 218.9189189 {
		numberout = 36
	}
	if number > 218.9189189 && number <= 228.6486486 {
		numberout = 24
	}
	if number > 228.6486486 && number <= 238.3783784 {
		numberout = 3
	}
	if number > 238.3783784 && number <= 248.1081081 {
		numberout = 15
	}
	fmt.Println(numberout)
}

/*currUser := currentUser(r)
_, selectedColor := getUserBetValue(w, r, currUser)
fmt.Println("DOES IT WORK?", selectedColor)*/
