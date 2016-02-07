package ripego

const (
	LACNIC_WHOIS_SERVER = "whois.lacnic.net"
)

type Lacnic struct {
}

func (r Lacnic) Check(search string) (w WhoisInfo, err error) {
	whoisData, err := GetTcpContent(search, LACNIC_WHOIS_SERVER)

	if err != nil {
		return w, err
	}

	wi := WhoisInfo{}
	wi.Inetnum = ParseRPSLValue(whoisData, "inetnum", "inetnum")
	wi.Status = ParseRPSLValue(whoisData, "inetnum", "status")
	wi.Netname = ParseRPSLValue(whoisData, "inetnum", "ownerid")
	wi.AdminC = ParseRPSLValue(whoisData, "inetnum", "owner-c")
	wi.Country = ParseRPSLValue(whoisData, "inetnum", "country")
	wi.Descr = ParseRPSLValue(whoisData, "inetnum", "owner")
	wi.LastModified = ParseRPSLValue(whoisData, "inetnum", "changed")
	wi.MntBy = ParseRPSLValue(whoisData, "inetnum", "mnt-by")
	wi.MntLower = ParseRPSLValue(whoisData, "inetnum", "mnt-lower")
	wi.MntRoutes = ParseRPSLValue(whoisData, "inetnum", "mnt-routes")
	wi.Source = ParseRPSLValue(whoisData, "inetnum", "source")
	wi.TechC = ParseRPSLValue(whoisData, "inetnum", "tech-c")
	wi.Organization = ParseRPSLValue(whoisData, "inetnum", "owner")

	p := WhoisPerson{}
	p.Name = ParseRPSLValue(whoisData, "nic-hdl", "nic-hdl")
	p.AbuseMailbox = ParseRPSLValue(whoisData, "nic-hdl", "e-mail")
	p.Address = ParseRPSLValue(whoisData, "nic-hdl", "address")
	p.LastModified = ParseRPSLValue(whoisData, "nic-hdl", "changed")
	p.NicHdl = ParseRPSLValue(whoisData, "role", "nic-hdl")
	p.Phone = ParseRPSLValue(whoisData, "role", "phone")
	p.Source = ParseRPSLValue(whoisData, "p.Co", "source")

	wi.Person = p

	return wi, err
}
