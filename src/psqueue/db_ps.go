package psqueue

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Problem struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Level  int    `json:"level"`
	URL    string `json:"url"`
	Status string `json:"status"`
	Memo   string `json:"memo"`
	User   string `json:"user"`
}

func GetProblemList(user string) []Problem {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(`SELECT id, title, level, url, status, memo FROM problem WHERE user = ?`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	problems := make([]Problem, 0)

	res, err := stmt.Query(user)

	for res.Next() {
		var problem Problem
		err = res.Scan(&problem.ID, &problem.Title, &problem.Level, &problem.URL, &problem.Status, &problem.Memo)
		problems = append(problems, problem)
	}
	return problems
}

func AddProblem(title string, level int, url string, status string, memo string) {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO problem(title, level, url, status, memo) values(?, ?, ?, ?, ?)")

	_, err = stmt.Exec(title, level, url, status, memo)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateProblem(pid int, title string, level int, url string, status string, memo string) {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE problem SET title = ?, level = ?, url = ?, status = ?, memo = ? WHERE id = ?")

	_, err = stmt.Exec(title, level, url, status, memo, pid)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteProblem(pid int) {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE problem WHERE id = ?")

	_, err = stmt.Exec(pid)
	if err != nil {
		log.Fatal(err)
	}
}
