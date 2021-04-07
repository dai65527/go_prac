package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "よんひゃくよん　のっとふぁうんど", 404)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
