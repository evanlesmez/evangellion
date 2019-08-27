package main

import (
	"evangellion/db"
	"fmt"
)

func main() {
	a, _ := db.SnagAnimation()
	fmt.Println(a)
}
