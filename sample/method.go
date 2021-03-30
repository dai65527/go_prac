package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func (h Hex) Print() {
	fmt.Printf("%x\n", int(h))
}

func main() {
	var hex Hex = 42
	hex.Print()
	fmt.Println(hex.String())
}
