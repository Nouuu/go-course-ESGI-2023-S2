package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func initDB(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS books (" +
		"id INT AUTO_INCREMENT PRIMARY KEY, " +
		"title VARCHAR(255), " +
		"author VARCHAR(255), " +
		"date INT);")
	return err
}

func insert3Books(db *sql.DB) {
	_, err := db.Exec("INSERT INTO books (title, author, date) VALUES (?, ?, ?)",
		"Harry Potter", "JK", 1998)
	handleErr(err)
	_, err = db.Exec("INSERT INTO books (title, author, date) VALUES (?, ?, ?)",
		"Eragon", "Paolini", 2002)
	handleErr(err)
	_, err = db.Exec("INSERT INTO books (title, author, date) VALUES (?, ?, ?)",
		"Sorceleur", "sdkhfzojfoiuzfjz", 2005)
	handleErr(err)
}

type Books struct {
	id     int
	title  string
	author string
	date   int
}

func selectBooks(db *sql.DB, year int) []Books {
	rows, err := db.Query("SELECT id, title, author, date FROM books where date > ?", year)
	handleErr(err)
	books := []Books{}
	for rows.Next() {
		var id, date int
		var title, author string
		err = rows.Scan(&id, &title, &author, &date)
		handleErr(err)
		books = append(books, Books{
			id:     id,
			title:  title,
			author: author,
			date:   date,
		})
	}
	return books
}

func main() {

	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/mydatabase")
	handleErr(err)
	defer db.Close()
	err = initDB(db)
	handleErr(err)
	//insert3Books(db)
	books := selectBooks(db, 1900)
	fmt.Println(books)

}
