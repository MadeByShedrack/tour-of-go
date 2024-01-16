package medium

import "fmt"

func Pointers() {
	index, num := 42, 2701

	p := &index
	fmt.Printf("Pointer value: %v\n", *p)
	*p = 21
	fmt.Printf("Another value: %v\n", index)

	p = &num
	*p = *p / 37
	fmt.Println("Num ->", num)
}
