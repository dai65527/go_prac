package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var nFlg = flag.Bool("n", false, "display line number")

func displayLines(fs *os.File, lineNo *int) {
	scanner := bufio.NewScanner(fs)
	for scanner.Scan() {
		if *nFlg {
			fmt.Printf("%6d ", *lineNo)
		}
		fmt.Printf("%s\n", scanner.Text())
		*lineNo++
	}
}

func displayFile(filename string, lineNo *int) {
	fs, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		displayLines(fs, lineNo)
	}
	fs.Close()
}

func main() {
	flag.Parse()

	lineNo := 1
	if len(flag.Args()) == 0 {
		displayLines(os.Stdin, &lineNo)
	} else {
		for i := 0; i < len(flag.Args()); i++ {
			displayFile(flag.Args()[i], &lineNo)
		}
	}
}
