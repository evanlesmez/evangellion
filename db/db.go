package db

import (
	"evangellion/models"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var Schema = `
CREATE TABLE IF NOT EXISTS animation (
    artist text,
	source []byte
);
CREATE TABLE IF NOT EXISTS songs (
	artist text,
	source []byte,
	vibe text
)
`

// TODO music schema https://www.free-stock-music.com/
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func OpenDb() (*sqlx.DB, error) {
	// Remember to defer close
	db, err := sqlx.Open("sqlite3", "/home/evan/code/go/src/evangellion/db/evangellion.db")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
func SnagAnimation() (*models.Animation, error) {
	db, err := OpenDb()
	defer db.Close()
	animations := []models.Animation{}
	err = db.Select(&animations, "SELECT * FROM animation ORDER BY RANDOM() LIMIT 1")
	return &animations[0], err
}

// TODO should only be able to add blobs not already in there
// TODO should be able to loop through entire animations directory
func PopulateAnimations(db *sqlx.DB) {
	artist, animationPath := "kirokazepixel", "/home/evan/code/go/src/evangellion/assets/animations/"
	files, err := ioutil.ReadDir(animationPath + artist)
	if err != nil {
		log.Fatal(err)
	}
	tx := db.MustBegin()
	for _, f := range files {
		dat, err := ioutil.ReadFile(animationPath + artist + "/" + f.Name())
		tx := db.MustBegin()
		_, err = tx.NamedExec("INSERT INTO animation (artist, source) VALUES (:artist, :source)", &models.Animation{Source: dat, Artist: artist})
		check(err)
		tx.Commit()
	}
	tx.Commit()
}

func BuildSchema(db *sqlx.DB) {
	fmt.Println("Building tabls in the DB")
	db.MustExec(Schema)
}

// func BuildAnimationTable(db *sqlx.DB) {
// 	fmt.Println("building animation table")
// 	statement, err := db.Prepare("DROP TABLE IF EXISTS animation")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	statement.Exec()
// 	db.MustExec(Schema)
// }

// TODO make an autoincrement id

func SelectAllAnimations(database *sqlx.DB) {
	animations := []models.Animation{}
	database.Select(&animations, "SELECT * FROM animation ORDER BY artist ASC")
}
