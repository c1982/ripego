package ripego

type Providers struct {
	p map[interface{}]struct{}
}

type Whois interface {
	Check(search string) (WhoisInfo, error)
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
