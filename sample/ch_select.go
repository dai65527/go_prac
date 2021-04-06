package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() { ch1 <- "hoge" }()
	go func() { ch2 <- "fuga" }()
	for i := 0; i < 2; i++ {
		select {
		case v1 := <-ch1:
			fmt.Println(v1)
		case v2 := <-ch2:
			fmt.Println(v2)
		}
	}
}
