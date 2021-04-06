package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	fmt.Println("goroutine1 start")
	wg.Add(2) // wait for 2 routines
	go func() {
		defer func() {
			defer fmt.Println("goroutine1 done")
			wg.Done()
		}()
		time.Sleep(3 * time.Second)
	}()
	fmt.Println("goroutine2 start")
	go func() {
		defer func() {
			fmt.Println("goroutine2 done")
			wg.Done()
		}()
		time.Sleep(1 * time.Second)
	}()
	wg.Wait()
}
