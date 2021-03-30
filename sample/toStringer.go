package main

import (
	"fmt"
	"os"
)

type Stringer interface {
	String() string
}

type TypeError string

func (e TypeError) Error() string {
	return string(e)
}

func ToStringer(arg interface{}) (Stringer, error) {
	// 型アサーション
	if s, ok := arg.(Stringer); ok {
		return s, nil
	}
	return nil, TypeError("Cast error")
}

type MyInt int

func (v MyInt) String() string {
	return fmt.Sprintf("%d", v)
}

func main() {
	// case OK
	v := MyInt(42)
	fmt.Println("MyInt has string method")

	if s, err := ToStringer(v); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		fmt.Println(s.String())
	}

	// case Failure
	v2 := 42 // int
	fmt.Println("nomal int")

	if s, err := ToStringer(v2); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
	} else {
		fmt.Println(s.String())
	}
}
