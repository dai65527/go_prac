package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		ch1 <- "hogehoge"
		str := <-ch2
		fmt.Println(str)
	}()
	go func() {
		defer wg.Done()
		ch2 <- "fugafuga"
		str := <-ch1
		fmt.Println(str)
	}()
	wg.Wait()
}
