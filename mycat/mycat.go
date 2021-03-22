package main

import (
	"bufio"
	"fmt"
	"os"
)

func displayFile(filename string, lineNo *int) {
	fs, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	scanner := bufio.NewScanner(fs)
	for scanner.Scan() {
		fmt.Printf("%4d: %s\n", *lineNo, scanner.Text())
		*lineNo++
	}
}

func main() {
	lineNo := 1
	for i := 1; i < len(os.Args); i++ {
		displayFile(os.Args[i], &lineNo)
	}
}
