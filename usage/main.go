package main

import (
	"fmt"
	"log"

	"github.com/sebastianbroekhoven/ripego"
)

func main() {

	w, err := ripego.IPv6Lookup("2001:1af8:4101:1::1:1")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inetnum: " + w.Inetnum)
	fmt.Println("Source : " + w.Source)
	fmt.Println("Netname: " + w.Netname)
	fmt.Println("Descr  : " + w.Descr)
	fmt.Println("Country: " + w.Country)

}
