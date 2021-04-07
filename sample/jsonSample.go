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
	// encode json
	var buf bytes.Buffer
	p := Person{Name: "dnakano", Age: 26}
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Print(buf.String())

	// decode json
	var p2 Person
	dec := json.NewDecoder(&buf)
	if err := dec.Decode(&p2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p2)
}
