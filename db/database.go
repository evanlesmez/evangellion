package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE animation (
    artist text,
	source text
);
`

// Animation struct
type Animation struct {
	Source string `db:"source"`
	Artist string `db:"artist"`
}

func main() {

	database, err := sqlx.Open("sqlite3", "./evangellion.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	resetAnimationTable(database)

	populateDB(database)

	animations := []Animation{}
	err = database.Select(&animations, "SELECT * FROM animation ORDER BY artist ASC")
	fmt.Println(err, animations)
}

func populateDB(db *sqlx.DB) {
	artist, animationPath := "valenberg", "/home/evan/code/go/src/evangellion/assets/animations/"
	files, err := ioutil.ReadDir(animationPath + artist)
	if err != nil {
		log.Fatal(err)
	}
	tx := db.MustBegin()
	for _, f := range files {
		fmt.Println(f.Name())
		path := filepath.Join(animationPath, f.Name())
		_, err = tx.NamedExec("INSERT INTO animation (artist, source) VALUES (:artist, :source)", &Animation{Source: path, Artist: artist})
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func resetAnimationTable(database *sqlx.DB) {
	fmt.Println("reseting")
	statement, err := database.Prepare("DROP TABLE IF EXISTS animation")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	database.MustExec(schema)
}
