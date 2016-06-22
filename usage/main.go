package main

import (
	"fmt"
	"log"

	"github.com/c1982/ripego"
)

func main() {

	w, err := ripego.IPLookup("178.18.196.250")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inetnum: " + w.Inetnum)
	fmt.Println("Desc: " + w.Descr)
}
