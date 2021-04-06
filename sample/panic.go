package main

import "fmt"

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("PANIC in func")
}

func main() {
	f()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("PANIC in main")
}
