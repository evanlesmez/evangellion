package db

import (
	"evangellion/models"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE animation (
    artist text,
	source []byte
);
`

func OpenAnimationDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "/home/evan/code/go/src/evangellion/db/evangellion.db")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}
func SnagAnimation() (*models.Animation, error) {
	db, err := OpenAnimationDb()
	defer db.Close()
	animations := []models.Animation{}
	err = db.Select(&animations, "SELECT * FROM animation ORDER BY RANDOM() LIMIT 1")
	return &animations[0], err
}

// func PopulateDB(db *sqlx.DB) {
// 	artist, animationPath := "valenberg", "/home/evan/code/go/src/evangellion/assets/animations/"
// 	files, err := ioutil.ReadDir(animationPath + artist)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	tx := db.MustBegin()
// 	for _, f := range files {
// 		fmt.Println(f.Name())
// 		path := filepath.Join(animationPath, f.Name())
// 		_, err = tx.NamedExec("INSERT INTO animation (artist, source) VALUES (:artist, :source)", &Animation{Source: , Artist: artist})
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	tx.Commit()
// }

func ResetAnimationTable() {
	db, err := OpenAnimationDb()
	defer db.Close()
	fmt.Println("reseting")
	statement, err := db.Prepare("DROP TABLE IF EXISTS animation")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()
	db.MustExec(schema)
}

// TODO make an autoincrement id

func SelecatAllAnimations(database *sqlx.DB) {
	animations := []models.Animation{}
	database.Select(&animations, "SELECT * FROM animation ORDER BY artist ASC")
}
