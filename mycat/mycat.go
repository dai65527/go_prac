package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var nFlg = flag.Bool("n", false, "display line number")

func displayFile(filename string, lineNo *int) {
	fs, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fs.Close()
		return
	}

	scanner := bufio.NewScanner(fs)
	for scanner.Scan() {
		if *nFlg {
			fmt.Printf("%4d: ", *lineNo)
		}
		fmt.Printf("%s\n", scanner.Text())
		*lineNo++
	}

	fs.Close()
}

func main() {
	flag.Parse()
	lineNo := 1
	fmt.Println(flag.Args())
	for i := 0; i < len(flag.Args()); i++ {
		displayFile(flag.Args()[i], &lineNo)
	}
}
