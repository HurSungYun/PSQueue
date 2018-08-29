package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"psqueue"
	"strconv"
)

var tempProblemList []psqueue.Problem

func init() {
	tempProblemList = make([]psqueue.Problem, 3)
	tempProblemList[0] = psqueue.Problem{1, "별찍기1", 3, "https://www.naver.com", "NOT_SOLVED", "류가 내준 문제1", "ethanhur"}
	tempProblemList[1] = psqueue.Problem{2, "별찍기2", 1, "https://www.naver.com", "NOT_SOLVED", "류가 내준 문제2", "ethanhur"}
	tempProblemList[2] = psqueue.Problem{3, "별찍기3", 2, "https://www.naver.com", "NOT_SOLVED", "류가 내준 문제3", "ethanhur"}
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
	Error    int               `json:"error"`
	Problems []psqueue.Problem `json:"problem_list"`
}

type SimpleResponse struct {
	Error int `json:"error"`
}

func queueListHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user := r.FormValue("user")

	pList := psqueue.GetProblemList(user)

	resData := QueueListResponse{0, pList}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}

func addProblemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	title := r.FormValue("title")
	level, err := strconv.Atoi(r.FormValue("level"))
	if err != nil {
		log.Fatal(err)
	}
	url := r.FormValue("url")
	status := r.FormValue("string")
	memo := r.FormValue("memo")

	psqueue.AddProblem(title, level, url, status, memo)
	resData := SimpleResponse{0}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}

func updateProblemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Fatal(err)
	}
	title := r.FormValue("title")
	level, err := strconv.Atoi(r.FormValue("level"))
	if err != nil {
		log.Fatal(err)
	}
	url := r.FormValue("url")
	status := r.FormValue("string")
	memo := r.FormValue("memo")

	psqueue.UpdateProblem(id, title, level, url, status, memo)

	resData := SimpleResponse{0}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}

func deleteProblemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Fatal(err)
	}

	psqueue.DeleteProblem(id)

	resData := SimpleResponse{0}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resData)
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/queue/list", queueListHandler)
	http.HandleFunc("/queue/add", addProblemHandler)
	http.HandleFunc("/queue/edit", updateProblemHandler)
	http.HandleFunc("/queue/delete", deleteProblemHandler)

	serveSingle("/bundle.js", "public/bundle.js")
	if err := http.ListenAndServe(":15395", nil); err != nil {
		log.Fatal(err)
	}
}
