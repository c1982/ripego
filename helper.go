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
	RPSL_LINE_PATTERN = `(.+):\W+(.+)`
)

func GetTcpContent(search string, host string) (s string, err error) {

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

func ParseRPSLValue(whoisText string, class string, section string) string {

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
				sectionValue = ParseRPSLine(line)
				break
			}
		}
	}

	return sectionValue
}

func ParseRPSLine(whoisLine string) string {

	rx, _ := regexp.Compile(RPSL_LINE_PATTERN)
	s := rx.FindAllStringSubmatch(whoisLine, -1)

	if len(s) >= 1 {
		return s[0][2]
	}

	return ""
}
