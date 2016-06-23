package ripego

import (
	"errors"

	"github.com/sebastianbroekhoven/go-get-ianawhois"
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

// IpLookup function for lagacy, not breaking stuff
func IpLookup(ipaddr string) (w WhoisInfo, err error) {
	if !isValidIp(ipaddr) {
		return w, errors.New("Invalid IPv4 address: " + ipaddr)
	}

	w, err = getNicProvider(ipaddr).Check(ipaddr)
	return w, err
}

// IPLookup function that returns IP information at provider and returns information.
func IPLookup(ipaddr string) (w WhoisInfo, err error) {
	if !isValidIp(ipaddr) {
		return w, errors.New("Invalid IPv4 address: " + ipaddr)
	}

	w, err = getNicProvider(ipaddr).Check(ipaddr)
	return w, err
}

// IPv4Lookup function that returns IP information at provider and returns information.
func IPv4Lookup(ipaddr string) (w WhoisInfo, err error) {
	if !isValidIp(ipaddr) {
		return w, errors.New("Invalid IPv4 address: " + ipaddr)
	}

	resp, err := whois.Query(ipaddr)
	if err != nil {
		return w, errors.New("Query failed for: " + ipaddr)
	}

	server, org := whois.Server(resp)

	if org == "afrinic" {
		w, err = AfrinicCheck(ipaddr)
	} else if org == "apnic" {
		w, err = ApnicCheck(ipaddr)
	} else if org == "arin" {
		w, err = ArinCheck(ipaddr)
	} else if org == "lacnic" {
		w, err = LacnicCheck(ipaddr)
	} else {
		w, err = RipeCheck(ipaddr)
	}

	println(server)
	// w, err = getNicProvider(ipaddr).Check(ipaddr)
	return w, err
}

// GetNicProvider function that search for the right provider for the lookup.
func getNicProvider(ipaddr string) Whois {

	var d = getNic["ripe"]

	for w := range getNic {
		if getNic[w].hasIP(ipaddr) {
			d = getNic[w]
			break
		}
	}

	return d
}

// Whois intercate containing the resulting infomration.
type Whois interface {
	Check(search string) (WhoisInfo, error)
	hasIP(ipaddr string) bool
}

// WhoisInfo struct with information on IP address range.
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

// WhoisPerson struct for Person information from provider.
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

// WhoisRoute struct for Route and Network information from provider.
type WhoisRoute struct {
	Route        string
	Descr        string
	Origin       string
	MntBy        string
	Created      string
	LastModified string
	Source       string
}
