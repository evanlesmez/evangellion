package main

import (
	"evangellion/db"
	"log"
)

func main() {
	datbase, err := db.OpenDb()
	defer datbase.Close()
	if err != nil {
		log.Fatal(err)
	}

	db.BuildSchema(datbase)
	// db.BuildAnimationTable(datbase)
	// db.PopulateAnimations(datbase)

}
