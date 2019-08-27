package main

import (
	"evangellion/db"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var templates = template.Must(template.ParseFiles("../templates/home.html"))

func loopAnimations() {
	files, err := ioutil.ReadDir("../assets/animations/valenberg")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
		path := filepath.Join("/home/evan/code/go/src/evangellion/", f.Name())
		fmt.Println(path)
	}
}

func main() {
	a, _ := db.SnagAnimation()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, *a)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// TODO hookup sqlite3 database with curated animations & music

// TODO generate random animation each reload
// TODO change animation every given time frame

// TODO play a song
// TODO play a song station

// TODO display animations & play music on each page (lofi with my fav animations to start)

// TODO add submit an animation to my database or song to me for review to put in queue
// TODO give me email if you want to submit a new one

// TODO host this on wifi, storage microelectronic (my raspberry pi probably once new one comes in)

// TODO fetch new animations & music in backend based on genre tags
// TODO filter animations & music based on mood

// TODO Donate to me with cash

// TODO use bitcoin to vote on next song out of 3

// TUI is white, LA is black
