package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://localhost:8080/?p=dnakano")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	buf := bufio.NewScanner(res.Body)
	buf.Scan()
	fmt.Println(buf.Text())
}
