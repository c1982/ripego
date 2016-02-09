package main

import (
	"fmt"
	"log"
	"ripego"
)

func main() {

	w, err := ripego.IpLookup("178.18.196.250")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inetnum: " + w.Inetnum)
	fmt.Println("Desc: " + w.Descr)
}
