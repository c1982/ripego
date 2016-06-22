package ripego

type apnic struct {
}

func (r apnic) Check(search string) (w WhoisInfo, err error) {
	whoisData, err := getTcpContent(search, apnic_whois_server)

	if err != nil {
		return w, err
	}

	wi := WhoisInfo{}
	wi.Inetnum = parseRPSLValue(whoisData, "inetnum", "inetnum")
	wi.Netname = parseRPSLValue(whoisData, "inetnum", "netname")
	wi.AdminC = parseRPSLValue(whoisData, "inetnum", "admin-c")
	wi.Country = parseRPSLValue(whoisData, "inetnum", "country")
	wi.Descr = parseRPSLValue(whoisData, "inetnum", "descr")
	wi.LastModified = parseRPSLValue(whoisData, "inetnum", "changed")
	wi.MntBy = parseRPSLValue(whoisData, "inetnum", "mnt-by")
	wi.MntLower = parseRPSLValue(whoisData, "inetnum", "mnt-lower")
	wi.MntRoutes = parseRPSLValue(whoisData, "inetnum", "mnt-routes")
	wi.Source = parseRPSLValue(whoisData, "inetnum", "source")
	wi.TechC = parseRPSLValue(whoisData, "inetnum", "tech-c")
	wi.Organization = parseRPSLValue(whoisData, "irt", "irt")

	p := WhoisPerson{}
	p.Name = parseRPSLValue(whoisData, "role", "role")
	p.AbuseMailbox = parseRPSLValue(whoisData, "irt", "abuse-mailbox")
	p.Address = parseRPSLValue(whoisData, "role", "address")
	p.LastModified = parseRPSLValue(whoisData, "role", "changed")
	p.MntBy = parseRPSLValue(whoisData, "role", "mnt-by")
	p.NicHdl = parseRPSLValue(whoisData, "role", "nic-hdl")
	p.Phone = parseRPSLValue(whoisData, "role", "phone")
	p.Source = parseRPSLValue(whoisData, "role", "source")

	wi.Person = p

	return wi, err
}

// hasIP function for derterming the right provider
func (r apnic) hasIP(ipaddr string) bool {
	//http://www.iana.org/assignments/ipv4-address-space/ipv4-address-space.xhtml
	ips := []string{"1", "14", "27", "36", "39", "42", "49", "49", "58", "59",
		"60", "61", "101", "103", "106", "110", "111", "112", "113", "114",
		"115", "116", "117", "118", "119", "120", "121", "122", "123", "124",
		"125", "126", "133", "150", "153", "163", "171", "175", "180", "182", "183",
		"202", "203", "210", "211", "218", "219", "220", "221", "222", "223"}

	return isProviderIP(ipaddr, ips)
}
