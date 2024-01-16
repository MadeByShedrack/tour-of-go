package basics

import "fmt"

func Stacking() {
	fmt.Println("Counting....")

	for index := 0; index < 10; index++ {
		defer fmt.Printf("Index -> %v\n", index)
	}

	fmt.Println("Done....")
}
