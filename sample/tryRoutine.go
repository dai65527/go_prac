package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("main done")
	fmt.Println("goroutine1 start")
	go func() {
		defer fmt.Println("goroutine1 done")
		time.Sleep(3 * time.Second)
	}()
	fmt.Println("goroutine2 start")
	go func() {
		defer fmt.Println("goroutine2 done")
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(5 * time.Second)
}
