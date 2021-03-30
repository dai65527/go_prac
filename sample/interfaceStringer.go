package main

import "fmt"

type Stringer interface {
	String() string
}

type Oct int
type Hex int
type Float float64

func (oct Oct) String() string {
	return fmt.Sprintf("%o", int(oct))
}

func (hex Hex) String() string {
	return fmt.Sprintf("%x", int(hex))
}

func (flt Float) String() string {
	return fmt.Sprintf("%4.2f", float64(flt))
}

func main() {
	var n int = 42
	var oct Oct = 42
	var hex Hex = 42
	var flt Float = 42
	var stringer Stringer

	fmt.Println("no use stringer, n =", n)
	fmt.Println(oct.String())
	fmt.Println(hex.String())
	fmt.Println(flt.String())

	fmt.Println("use stringer, n =", n)
	stringer = oct
	fmt.Println(stringer.String())
	stringer = hex
	fmt.Println(stringer.String())
	stringer = flt
	fmt.Println(stringer.String())
}
