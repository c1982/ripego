package ripego

import (
	_ "fmt"
)

const (
	ARIN_WHOIS_SERVER = "whois.arin.net"
)

type Arin struct {
}

func (r Arin) Check(search string) (w WhoisInfo, err error) {
	whoisData, err := GetTcpContent(search, ARIN_WHOIS_SERVER)

	if err != nil {
		return w, err
	}

	wi := WhoisInfo{}
	wi.Inetnum = ParseRPSLValue(whoisData, "NetRange", "NetRange")
	wi.Netname = ParseRPSLValue(whoisData, "NetRange", "NetName")
	wi.Organization = ParseRPSLValue(whoisData, "NetRange", "Organization")
	wi.Created = ParseRPSLValue(whoisData, "NetRange", "RegDate")
	wi.LastModified = ParseRPSLValue(whoisData, "NetRange", "Updated")
	wi.Status = ParseRPSLValue(whoisData, "NetRange", "NetType")

	rt := WhoisRoute{}
	rt.Origin = ParseRPSLValue(whoisData, "NetRange", "OriginAS")
	rt.Route = ParseRPSLValue(whoisData, "NetRange", "CIDR")

	wi.Route = rt

	return wi, err
}
