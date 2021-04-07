package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

var tmpl = template.Must(template.New("omikuji").Parse("<html><body>{{.Name}}さんの運勢は<b>{{.Omikuji}}</b>です</body></html>"))

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

type Result struct {
	Name    string
	Omikuji string
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("p")
	if len(name) == 0 {
		name = "Guest"
	}
	result := Result{Name: name, Omikuji: getFortune()}
	tmpl.Execute(w, result)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
