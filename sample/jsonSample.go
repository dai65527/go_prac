package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var buf bytes.Buffer
	p := Person{Name: "dnakano", Age: 26}
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Print(buf.String())
}
