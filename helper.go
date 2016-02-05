package ripego

import (
	"io/ioutil"
	"net"
	"time"
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
