# ripego

Bu paket IPv4 adresini ilgili kordinasyon merkezini tespit ederek Whois sorgulaması yapar ve IP hakkında bilgiyi getirir. Desteklediği kordinasyon merkezleri. ripe, arin, apnic, afrinic, lacnic

### Yükleme

```bash
$ go get github.com/c1982/ripego
```

### Örnek Kullanım

```golang
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
```