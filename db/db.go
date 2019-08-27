package db

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

func openAnimationDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "/home/evan/code/go/src/evangellion/db/evangellion.db")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
func SnagAnimation() (*Animation, error) {
	db, err := openAnimationDb()
	defer db.Close()
	animations := []Animation{}
	err = db.Select(&animations, "SELECT * FROM animation ORDER BY RANDOM() LIMIT 1")
	return &animations[0], err
}
func OpenAnimationTable() (*sqlx.DB, error) {
	database, err := sqlx.Open("sqlite3", "./evangellion.db")
	if err != nil {
		log.Fatal(err)
	}
	return database, err
}
func PopulateDB(db *sqlx.DB) {
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

func ResetAnimationTable(database *sqlx.DB) {
	fmt.Println("reseting")
	statement, err := database.Prepare("DROP TABLE IF EXISTS animation")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	database.MustExec(schema)
}

// TODO make an autoincrement id

func SelecatAllAnimations(database *sqlx.DB) {
	animations := []Animation{}
	database.Select(&animations, "SELECT * FROM animation ORDER BY artist ASC")
}
