package ripego

const (
	APNIC_WHOIS_SERVER = "whois.apnic.net"
)

type Apnic struct {
}

func (r Apnic) Check(search string) (w WhoisInfo, err error) {
	whoisData, err := GetTcpContent(search, APNIC_WHOIS_SERVER)

	if err != nil {
		return w, err
	}

	wi := WhoisInfo{}
	wi.Inetnum = ParseRPSLValue(whoisData, "inetnum", "inetnum")
	wi.Netname = ParseRPSLValue(whoisData, "inetnum", "netname")
	wi.AdminC = ParseRPSLValue(whoisData, "inetnum", "admin-c")
	wi.Country = ParseRPSLValue(whoisData, "inetnum", "country")
	wi.Descr = ParseRPSLValue(whoisData, "inetnum", "descr")
	wi.LastModified = ParseRPSLValue(whoisData, "inetnum", "changed")
	wi.MntBy = ParseRPSLValue(whoisData, "inetnum", "mnt-by")
	wi.MntLower = ParseRPSLValue(whoisData, "inetnum", "mnt-lower")
	wi.MntRoutes = ParseRPSLValue(whoisData, "inetnum", "mnt-routes")
	wi.Source = ParseRPSLValue(whoisData, "inetnum", "source")
	wi.TechC = ParseRPSLValue(whoisData, "inetnum", "tech-c")
	wi.Organization = ParseRPSLValue(whoisData, "irt", "irt")

	p := WhoisPerson{}
	p.Name = ParseRPSLValue(whoisData, "role", "role")
	p.AbuseMailbox = ParseRPSLValue(whoisData, "irt", "abuse-mailbox")
	p.Address = ParseRPSLValue(whoisData, "role", "address")
	p.LastModified = ParseRPSLValue(whoisData, "role", "changed")
	p.MntBy = ParseRPSLValue(whoisData, "role", "mnt-by")
	p.NicHdl = ParseRPSLValue(whoisData, "role", "nic-hdl")
	p.Phone = ParseRPSLValue(whoisData, "role", "phone")
	p.Source = ParseRPSLValue(whoisData, "role", "source")

	wi.Person = p

	return wi, err
}
