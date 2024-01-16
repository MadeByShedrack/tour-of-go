package basics

import (
	"fmt"
	"time"
)

func Conditionals() {
	fmt.Println("When's saturday? ")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	myBirthYear(2010, 2024)

	myTime := time.Now()

	switch {
	case myTime.Hour() < 12:
		fmt.Println("Good morning")
	case myTime.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}
}

func myBirthYear(birthYear, year int) {
	if currentAge := year - birthYear; currentAge >= 18 {
		fmt.Println("You are of age. welcome")
	} else {
		fmt.Println("You are too young for this")
	}
}
