package main

import (
	"log"
	"net/http"
	"html/template"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("templates/index.html")
    t.Execute(w, nil)
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func main() {
	http.HandleFunc("/", mainHandler)
	serveSingle("/bundle.js", "./public/bundle.js")
	if err := http.ListenAndServe(":15395", nil); err != nil {
		log.Fatal(err)
	}
}
