package ripego

type arin struct {
}

func (r arin) Check(search string) (w WhoisInfo, err error) {
	whoisData, err := getTcpContent(search, arin_whois_server)

	if err != nil {
		return w, err
	}

	wi := WhoisInfo{}
	wi.Inetnum = parseRPSLValue(whoisData, "NetRange", "NetRange")
	wi.Netname = parseRPSLValue(whoisData, "NetRange", "NetName")
	wi.Organization = parseRPSLValue(whoisData, "NetRange", "Organization")
	wi.Created = parseRPSLValue(whoisData, "NetRange", "RegDate")
	wi.LastModified = parseRPSLValue(whoisData, "NetRange", "Updated")
	wi.Status = parseRPSLValue(whoisData, "NetRange", "NetType")

	rt := WhoisRoute{}
	rt.Origin = parseRPSLValue(whoisData, "NetRange", "OriginAS")
	rt.Route = parseRPSLValue(whoisData, "NetRange", "CIDR")

	wi.Route = rt

	return wi, err
}

// hasIP function for derterming the right provider
func (r arin) hasIP(ipaddr string) bool {
	//http://www.iana.org/assignments/ipv4-address-space/ipv4-address-space.xhtml
	ips := []string{"3", "4", "6", "7", "8", "9", "11", "12", "13", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24",
		"26", "28", "29", "30", "32", "33", "34", "35", "38", "40", "44", "45", "47", "48", "50", "52", "54", "55", "56", "63",
		"64", "65", "66", "67", "68", "69", "70", "71", "72", "73", "74", "75", "76", "96", "97", "98", "99", "100", "104", "107",
		"108", "128", "129", "130", "131", "132", "134", "135", "136", "137", "138", "139", "140", "142", "143", "144", "146", "147",
		"148", "149", "152", "155", "156", "157", "158", "159", "160", "161", "162", "164", "165", "166", "167", "168", "169",
		"170", "172", "173", "174", "184", "192", "198", "199", "204", "205", "206", "207", "208", "209", "214", "215", "216"}

	return isProviderIP(ipaddr, ips)
}
