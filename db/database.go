package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// resetAnimationTable()
	database, err := sql.Open("sqlite3", "./evangellion.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS animation (id INTEGER PRIMARY KEY, artist TEXT, source TEXT)")
	statement.Exec()
	defer statement.Close()

	tx, err := database.Begin()
	if err != nil {
		log.Fatal(err)
	}

	statement, err = tx.Prepare("insert into animation( artist, source) values( ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	for index := 0; index < 50; index++ {
		// artist := "valenberg"
		// source := "ti"
		_, err = statement.Exec("test"+string(index), "test"+string(index))
	}

	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	rows, _ := database.Query("SELECT id, artist, source from animation")
	var id int
	var artist string
	var source string
	for rows.Next() {
		rows.Scan(&id, &artist, &source)
		fmt.Println(strconv.Itoa(id) + ": " + artist + " " + source)
	}
}

func resetAnimationTable() {
	fmt.Println("reseting")
	database, err := sql.Open("sqlite3", "./evangellion.db")
	statement, _ := database.Prepare("DROP TABLE IF EXISTS animation")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS animation (id INTEGER PRIMARY KEY, artist TEXT, source TEXT)")
	statement.Exec()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
}
