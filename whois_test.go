package ripego

import "testing"

func TestRipeWhois(t *testing.T) {

	who := Whois(new(Ripe))
	info, err := who.Check("178.18.196.250")

	if err != nil {
		t.Error(err)
	}

	t.Log(info.Inetnum)
	t.Log(info.Netname)
	t.Log(info.Person.Name)
	t.Log(info.Route.Origin)
}
