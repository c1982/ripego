package ripego

import (
	"errors"
)

var getNic = make(map[string]Whois)

const (
	afrinic_whois_server = "whois.afrinic.net"
	apnic_whois_server   = "whois.apnic.net"
	arin_whois_server    = "whois.arin.net"
	lacnic_whois_server  = "whois.lacnic.net"
	ripe_whois_server    = "whois.ripe.net"
)

func init() {
	getNic["afrinic"] = afrinic{}
	getNic["apnic"] = apnic{}
	getNic["arin"] = arin{}
	getNic["lacnic"] = lacnic{}
	getNic["ripe"] = ripe{}
}

func IpLookup(ipaddr string) (w WhoisInfo, err error) {

	if !isValidIp(ipaddr) {
		return w, errors.New("Invalid IPv4 address: " + ipaddr)
	}

	w, err = getNicProvider(ipaddr).Check(ipaddr)
	return w, err
}

func getNicProvider(ipaddr string) Whois {

	var d = getNic["ripe"]

	for w := range getNic {
		if getNic[w].hasIp(ipaddr) {
			d = getNic[w]
			break
		}
	}

	return d
}

type Whois interface {
	Check(search string) (WhoisInfo, error)
	hasIp(ipaddr string) bool
}

type WhoisInfo struct {
	Inetnum      string
	Netname      string
	Descr        string
	Country      string
	Organization string
	AdminC       string
	TechC        string
	MntLower     string
	Status       string
	MntBy        string
	Created      string
	LastModified string
	Source       string
	MntRoutes    string
	Person       WhoisPerson
	Route        WhoisRoute
}

type WhoisPerson struct {
	Name         string
	Address      string
	Phone        string
	AbuseMailbox string
	NicHdl       string
	MntBy        string
	Created      string
	LastModified string
	Source       string
}

type WhoisRoute struct {
	Route        string
	Descr        string
	Origin       string
	MntBy        string
	Created      string
	LastModified string
	Source       string
}
