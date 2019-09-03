package main

import (
	"evangellion/db"
	"evangellion/models"
	"fmt"
	"io/ioutil"
	"log"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func TestRead() {
	dat, err := ioutil.ReadFile("/home/evan/code/go/src/evangellion/assets/animations/valenberg/valenberg.gif")
	check(err)
	fmt.Print(string(dat))
}

// Test adding  a new gif to db
func TestAdd() {
	db, err := db.OpenAnimationDb()
	check(err)
	defer db.Close()

	dat, err := ioutil.ReadFile("/home/evan/code/go/src/evangellion/assets/animations/valenberg/valenberg.gif")
	tx := db.MustBegin()
	_, err = tx.NamedExec("INSERT INTO animation (artist, source) VALUES (:artist, :source)", &models.Animation{Source: dat, Artist: "valenberg"})
	check(err)
	tx.Commit()
}
func main() {

}
