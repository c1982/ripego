package ripego

import (
	"bufio"
	"io/ioutil"
	"net"
	"regexp"
	"strings"
	"time"
)

const (
	rpsl_line_pattern = `(.+):\W+(.+)`
)

func getTcpContent(search string, host string) (s string, err error) {

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, "43"), time.Second*28)
	defer conn.Close()

	if err != nil {
		return s, err
	}

	conn.Write([]byte(search + "\r\n"))

	buffer, err := ioutil.ReadAll(conn)

	if err != nil {
		return s, err
	}

	s = string(buffer[:])

	return s, err
}

func parseRPSLValue(whoisText string, class string, section string) string {

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
				sectionValue = parseRPSLine(line)
				break
			}
		}
	}

	return sectionValue
}

func parseRPSLine(whoisLine string) string {

	rx, _ := regexp.Compile(rpsl_line_pattern)
	s := rx.FindAllStringSubmatch(whoisLine, -1)

	if len(s) >= 1 {
		return s[0][2]
	}

	return ""
}

func isProviderIP(ipaddr string, ips []string) bool {

	hasip := false
	octet := firstOctec(ipaddr)

	for i := range ips {
		if octet == ips[i] {
			hasip = true
			break
		}
	}

	return hasip
}

func firstOctec(ipaddr string) string {
	return strings.Split(ipaddr, ".")[0]
}

func isValidIp(ipaddr string) bool {
	ip := net.ParseIP(ipaddr)

	return ip.To4() != nil
}

func isValidIPv6(ipaddr string) bool {
	ip := net.ParseIP(ipaddr)

	return ip.To16() != nil
}
