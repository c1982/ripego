package ripego

import (
	"bufio"
	"regexp"
	"strings"
)

const (
	RPSL_LINE_PATTERN = `(.+):\W+(.+)`
)

type Ripe struct {
}

func (r Ripe) Check(search string) (w WhoisInfo, err error) {
	content, err := GetTcpContent(search, "whois.ripe.net")
	w, err = r.ParseData(content)

	return w, err
}

func (r Ripe) ParseData(whoisData string) (w WhoisInfo, err error) {

	wi := WhoisInfo{}
	wi.AdminC = r.ReadValue(whoisData, "inetnum", "admin-c")
	wi.Country = r.ReadValue(whoisData, "inetnum", "country")
	w.Created = r.ReadValue(whoisData, "inetnum", "created")
	w.Descr = r.ReadValue(whoisData, "inetnum", "descr")
	w.Inetnum = r.ReadValue(whoisData, "inetnum", "inetnum")
	w.LastModified = r.ReadValue(whoisData, "inetnum", "last-modified")
	w.MntBy = r.ReadValue(whoisData, "inetnum", "mnt-by")
	w.MntLower = r.ReadValue(whoisData, "inetnum", "mnt-lower")
	w.MntRoutes = r.ReadValue(whoisData, "inetnum", "mnt-routes")
	w.Netname = r.ReadValue(whoisData, "inetnum", "netname")
	w.Source = r.ReadValue(whoisData, "inetnum", "source")
	w.TechC = r.ReadValue(whoisData, "inetnum", "tech-c")
	w.Organization = r.ReadValue(whoisData, "inetnum", "org")

	p := WhoisPerson{}
	p.Name = r.ReadValue(whoisData, "person", "person")
	p.AbuseMailbox = r.ReadValue(whoisData, "person", "abuse-mailbox")
	p.Address = r.ReadValue(whoisData, "person", "address")
	p.Created = r.ReadValue(whoisData, "person", "created")
	p.LastModified = r.ReadValue(whoisData, "person", "last-modified")
	p.MntBy = r.ReadValue(whoisData, "person", "mnt-by")
	p.NicHdl = r.ReadValue(whoisData, "person", "nic-hdl")
	p.Phone = r.ReadValue(whoisData, "person", "phone")
	p.Source = r.ReadValue(whoisData, "person", "source")

	rt := WhoisRoute{}
	rt.Origin = r.ReadValue(whoisData, "route", "origin")
	rt.Created = r.ReadValue(whoisData, "route", "created")
	rt.Descr = r.ReadValue(whoisData, "route", "descr")
	rt.LastModified = r.ReadValue(whoisData, "route", "last-modified")
	rt.Route = r.ReadValue(whoisData, "route", "route")
	rt.Source = r.ReadValue(whoisData, "route", "source")

	wi.Person = p
	wi.Route = rt

	return wi, err
}

func (r Ripe) ReadValue(whoisText string, class string, section string) string {

	var sectionValue = ""
	var hasIn = false

	sc := bufio.NewScanner(strings.NewReader(whoisText))

	for sc.Scan() {
		var line = sc.Text()
		if strings.HasPrefix(line, class) {
			if hasIn == false {
				hasIn = true
			}
		}

		if hasIn {
			if strings.HasPrefix(line, section) {
				sectionValue = r.readValueFromLine(line)
				break
			}
		}
	}

	return sectionValue
}

func (r Ripe) readValueFromLine(whoisLine string) string {

	rx, _ := regexp.Compile(RPSL_LINE_PATTERN)
	s := rx.FindAllStringSubmatch(whoisLine, -1)

	if len(s) >= 1 {
		return s[0][2]
	}

	return ""
}
