package ripego

type ripe struct {
}

func (r ripe) Check(search string) (w WhoisInfo, err error) {
	whoisData, err := getTcpContent(search, ripe_whois_server)

	if err != nil {
		return w, err
	}

	wi := WhoisInfo{}
	wi.Inetnum = parseRPSLValue(whoisData, "inetnum", "inetnum")
	wi.Netname = parseRPSLValue(whoisData, "inetnum", "netname")
	wi.AdminC = parseRPSLValue(whoisData, "inetnum", "admin-c")
	wi.Country = parseRPSLValue(whoisData, "inetnum", "country")
	wi.Created = parseRPSLValue(whoisData, "inetnum", "created")
	wi.Descr = parseRPSLValue(whoisData, "inetnum", "descr")
	wi.LastModified = parseRPSLValue(whoisData, "inetnum", "last-modified")
	wi.MntBy = parseRPSLValue(whoisData, "inetnum", "mnt-by")
	wi.MntLower = parseRPSLValue(whoisData, "inetnum", "mnt-lower")
	wi.MntRoutes = parseRPSLValue(whoisData, "inetnum", "mnt-routes")
	wi.Source = parseRPSLValue(whoisData, "inetnum", "source")
	wi.TechC = parseRPSLValue(whoisData, "inetnum", "tech-c")
	wi.Organization = parseRPSLValue(whoisData, "inetnum", "org")
	wi.Status = parseRPSLValue(whoisData, "inetnum", "status")

	p := WhoisPerson{}
	p.Name = parseRPSLValue(whoisData, "person", "person")
	p.AbuseMailbox = parseRPSLValue(whoisData, "person", "abuse-mailbox")
	p.Address = parseRPSLValue(whoisData, "person", "address")
	p.Created = parseRPSLValue(whoisData, "person", "created")
	p.LastModified = parseRPSLValue(whoisData, "person", "last-modified")
	p.MntBy = parseRPSLValue(whoisData, "person", "mnt-by")
	p.NicHdl = parseRPSLValue(whoisData, "person", "nic-hdl")
	p.Phone = parseRPSLValue(whoisData, "person", "phone")
	p.Source = parseRPSLValue(whoisData, "person", "source")

	rt := WhoisRoute{}
	rt.Origin = parseRPSLValue(whoisData, "route", "origin")
	rt.Created = parseRPSLValue(whoisData, "route", "created")
	rt.Descr = parseRPSLValue(whoisData, "route", "descr")
	rt.LastModified = parseRPSLValue(whoisData, "route", "last-modified")
	rt.Route = parseRPSLValue(whoisData, "route", "route")
	rt.Source = parseRPSLValue(whoisData, "route", "source")

	wi.Person = p
	wi.Route = rt

	return wi, err
}

func (r ripe) hasIp(ipaddr string) bool {
	//http://www.iana.org/assignments/ipv4-address-space/ipv4-address-space.xhtml
	ips := []string{"2", "5", "25", "31", "37", "46", "51", "53", "57", "62",
		"77", "78", "79", "80", "81", "82", "83", "84", "85", "86", "87", "88", "89",
		"90", "91", "92", "93", "94", "95", "109", "141", "145", "151", "176", "178",
		"185", "188", "193", "194", "195", "212", "213", "217"}

	return isProviderIP(ipaddr, ips)
}
