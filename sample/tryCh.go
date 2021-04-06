package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print("> ")
		str, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(str)
	}
}
