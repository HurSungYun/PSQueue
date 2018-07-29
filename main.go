package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Problem struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Level     int       `json:"level"`
	URL       string    `json:"url"`
	Status    string    `json:"status"`
	Memo      string    `json:"memo"`
	CreatedAt time.Time `json:"created_at"`
}

var tempProblemList []Problem

func init() {
	tempProblemList = make([]Problem, 3)
	tempProblemList[0] = Problem{1, "별찍기1", 3, "https://www.naver.com", "NOT_SOLVED", "류가 내준 문제1", time.Now()}
	tempProblemList[1] = Problem{2, "별찍기2", 1, "https://www.naver.com", "NOT_SOLVED", "류가 내준 문제2", time.Now()}
	tempProblemList[2] = Problem{3, "별찍기3", 2, "https://www.naver.com", "NOT_SOLVED", "류가 내준 문제3", time.Now()}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

// api handler

type QueueListResponse struct {
	Error    int       `json:"error"`
	Problems []Problem `json:"problem_list"`
}

func QueueListHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	_ = r.FormValue("user")
	// var req QueueListRequset
	// err := json.NewDecoder(r.Body).Decode(&req)
	// if err != nil {
	// 	log.Fatal(err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	resData := QueueListResponse{0, tempProblemList}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/queue/list", QueueListHandler)
	serveSingle("/bundle.js", "./public/bundle.js")
	if err := http.ListenAndServe(":15395", nil); err != nil {
		log.Fatal(err)
	}
}
