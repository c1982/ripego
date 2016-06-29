package main

import (
	"fmt"
	"log"

	"github.com/sebastianbroekhoven/ripego"
)

func main() {

	w, err := ripego.IPv6Lookup("2001:4018:1404::1044")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inetnum: " + w.Inetnum)
	fmt.Println("Desc: " + w.Descr)
}
