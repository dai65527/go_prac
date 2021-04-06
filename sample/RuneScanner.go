package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"unicode/utf8"
)

type RuneScanner struct {
	reader io.Reader
	buf    [32]byte
}

func (scanner *RuneScanner) Scan() (rune, error) {
	// read
	n, err := scanner.reader.Read(scanner.buf[:]) // result will be stored
	if err != nil {
		return 0, err
	}

	// decode
	r, size := utf8.DecodeRune(scanner.buf[:n])
	if r == utf8.RuneError {
		return 0, errors.New("RuneError")
	}

	// unread buf
	scanner.reader = io.MultiReader(bytes.NewReader(scanner.buf[size:n]), scanner.reader)
	return r, nil
}

func NewRuneScanner(r io.Reader) *RuneScanner {
	return &RuneScanner{reader: r}
}

func main() {
	scanner := NewRuneScanner(strings.NewReader("Hello, 世界👨🏻‍💻"))
	for {
		r, err := scanner.Scan()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%c\n", r)
	}
}
