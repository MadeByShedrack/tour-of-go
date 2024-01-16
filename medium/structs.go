package medium

import "fmt"

type vertex struct {
	x int
	y int
}

type authentication struct {
	publicKey  string
	privateKey string
	owner      string
	isEncrypt  bool
}

func (auth authentication) verifyKey() {
	if len(auth.privateKey) < 8 {
		fmt.Println("Private key is too short")
	} else if len(auth.privateKey) == 8 || len(auth.privateKey) >= 8 {
		fmt.Println("Your private key is secure")
	} else {
		fmt.Println("No keys assigned")
	}
}

type numbers struct {
	num1, num2 int
}

var (
	n1 = numbers{1, 2}
	n2 = numbers{num1: 12}
	n3 = numbers{}
	n4 = &numbers{1, 2}
)

func Structs() {
	fmt.Println(vertex{1, 2})

	auth := authentication{
		publicKey:  "0xereuejjfjfjrk",
		privateKey: "0xahdhereysjskxneyemds",
		owner:      "@Madebyshedrack",
		isEncrypt:  true,
	}
	auth.verifyKey()
	fmt.Printf("Public Key -> %v\n", auth.publicKey)
	fmt.Printf("Private Key -> %v\n", auth.privateKey)

	myVertex := vertex{12, 14}
	p := &myVertex
	p.x = 1e9
	fmt.Println(myVertex)

	fmt.Println(n1, n2, n3, n4)
}
