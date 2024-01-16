package basics

import (
	"errors"
	"fmt"
)

func Defers() {
	// defer fmt.Println("World")

	// fmt.Println("Hello")
	skipNames("Shedrack", "Abel")

	num1 := 30
	num2 := 0

	result, err := handleErrors(num1, num2)

	if err == nil {
		fmt.Printf("Result -> %v\n", result)
	} else {
		fmt.Println(err.Error())
	}
}

func skipNames(first, last string) {
	defer fmt.Printf("First Name: %v\n", first)

	fmt.Printf("Last Name: %v\n", last)
}

func handleErrors(numberOne, numberTwo int) (int, error) {

	if numberTwo == 0 {
		return -1, errors.New("divide by zero erros")
	}

	return numberOne + numberTwo, nil
}
