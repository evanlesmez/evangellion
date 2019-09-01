package main

import (
	"evangellion/db"
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan bool)
	go func() {
	}()
	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			a, _ := db.SnagAnimation()
			fmt.Println("Snag a new one: ", a.Source)

		}
	}
}
