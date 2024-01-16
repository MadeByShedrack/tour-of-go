package medium

import "fmt"

func Arrays() {
	var arrays [2]string
	arrays[0] = "Hello"
	arrays[1] = "World"
	fmt.Println(arrays[0], arrays[1])
	fmt.Println(arrays)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	for _, prime := range primes {
		fmt.Printf("Primes -> %v\n", prime)
	}

	var slices []int = primes[1:5]
	fmt.Printf("Slices -> %v\n", slices)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	for _, name := range names {
		fmt.Printf("The Beatles -> %v\n", name)
	}

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
