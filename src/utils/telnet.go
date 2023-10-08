package utils

import (
	"strings"

	"github.com/ziutek/telnet"
)

func TelnetCMD(cmd string) (string, error) {
	conn, err := telnet.Dial("tcp", "localhost:7505")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	_, err = conn.Write([]byte(cmd + "\r\n"))
	if err != nil {
		return "", err
	}

	var resultBuilder strings.Builder

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return "", err
		}

		resultBuilder.Write(buf[:n])

		if strings.Contains(resultBuilder.String(), "END") {
			break
		}
	}

	return resultBuilder.String(), nil
}
