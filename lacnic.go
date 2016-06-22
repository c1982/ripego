package ripego

type lacnic struct {
}

func (r lacnic) Check(search string) (w WhoisInfo, err error) {
	whoisData, err := getTcpContent(search, lacnic_whois_server)

	if err != nil {
		return w, err
	}

	wi := WhoisInfo{}
	wi.Inetnum = parseRPSLValue(whoisData, "inetnum", "inetnum")
	wi.Status = parseRPSLValue(whoisData, "inetnum", "status")
	wi.Netname = parseRPSLValue(whoisData, "inetnum", "ownerid")
	wi.AdminC = parseRPSLValue(whoisData, "inetnum", "owner-c")
	wi.Country = parseRPSLValue(whoisData, "inetnum", "country")
	wi.Descr = parseRPSLValue(whoisData, "inetnum", "owner")
	wi.LastModified = parseRPSLValue(whoisData, "inetnum", "changed")
	wi.MntBy = parseRPSLValue(whoisData, "inetnum", "mnt-by")
	wi.MntLower = parseRPSLValue(whoisData, "inetnum", "mnt-lower")
	wi.MntRoutes = parseRPSLValue(whoisData, "inetnum", "mnt-routes")
	wi.Source = parseRPSLValue(whoisData, "inetnum", "source")
	wi.TechC = parseRPSLValue(whoisData, "inetnum", "tech-c")
	wi.Organization = parseRPSLValue(whoisData, "inetnum", "owner")

	p := WhoisPerson{}
	p.Name = parseRPSLValue(whoisData, "nic-hdl", "nic-hdl")
	p.AbuseMailbox = parseRPSLValue(whoisData, "nic-hdl", "e-mail")
	p.Address = parseRPSLValue(whoisData, "nic-hdl", "address")
	p.LastModified = parseRPSLValue(whoisData, "nic-hdl", "changed")
	p.NicHdl = parseRPSLValue(whoisData, "role", "nic-hdl")
	p.Phone = parseRPSLValue(whoisData, "role", "phone")
	p.Source = parseRPSLValue(whoisData, "p.Co", "source")

	wi.Person = p

	return wi, err
}

// hasIP function for derterming the right provider
func (r lacnic) hasIP(ipaddr string) bool {
	//http://www.iana.org/assignments/ipv4-address-space/ipv4-address-space.xhtml
	ips := []string{"177", "179", "181", "186", "187", "189", "190", "191", "200", "201"}

	return isProviderIP(ipaddr, ips)
}
