package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func getFortune() string {
	switch rand.Int() % 10 {
	case 0:
		return "大吉"
	case 1, 2:
		return "中吉"
	case 3, 4, 5, 6:
		return "吉"
	case 7, 8:
		return "凶"
	default:
		return "大凶"
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, getFortune())
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
