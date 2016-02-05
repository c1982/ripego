package ripego

import (
	_ "fmt"
)

type Ripe struct {
}

func (r Ripe) Check(search string) (w WhoisInfo, err error) {

	content, err := GetTcpContent(search, "whois.ripe.net")

	w, err = r.ParseData(content)

	return w, err
}

func (r Ripe) ParseData(whoisData string) (w WhoisInfo, err error) {

	//RSPL Parser

	return w, err
}
