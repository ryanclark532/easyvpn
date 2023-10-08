package vpn

import (
	"easyvpn/src/utils"
	vpn_dtos "easyvpn/src/vpn/vpn-dtos"
	"fmt"
	"strings"
	"time"
)

func GetVpnServerStatus() (string, error) {
	running, err := utils.IsProcessRunning("openvpn")
	if err != nil {
		return "unknown", err
	}
	if !running {
		return "notRunning", nil
	}

	containsInit, err := utils.ContainsSequence("src/log/openvpn.log", "Initialization Sequence Completed")
	if containsInit {
		return "running", err
	}
	return "starting", nil
}

func VpnOperation(operation string) error {
	var err error
	fmt.Println(operation)
	switch operation {
	case "start":
		go utils.StartVPNServer()
		return nil
	case "stop":
		err = utils.StopVPNServer()
		return err
	case "restart":
		err = utils.StopVPNServer()
		go utils.StartVPNServer()
		return err
	}
	return nil
}

func GetActiveConnections() (*[]vpn_dtos.ServerConnection, error) {
	result, err := utils.TelnetCMD("status")
	if err != nil {
		return nil, err
	}
	formatted, err := formatServerConnection(result)
	if err != nil {
		return nil, err
	}
	fmt.Println(formatted)

	return formatted, nil

}

func formatServerConnection(output string) (*[]vpn_dtos.ServerConnection, error) {
	lines := strings.Split(output, "\r\n")

	headerLine := ""
	for i, line := range lines {
		if line == "Common Name,Real Address,Bytes Received,Bytes Sent,Connected Since" {
			headerLine = lines[i]
			break
		}
	}

	headerIndex := strings.Index(output, headerLine)

	dataLines := strings.Split(output[headerIndex:], "\r\n")[1:]

	var connections []vpn_dtos.ServerConnection

	for _, dataLine := range dataLines {
		fields := strings.Split(dataLine, ",")
		if len(fields) == 5 {
			commonName := fields[0]
			realAddress := fields[1]
			bytesReceived := fields[2]
			bytesSent := fields[3]
			connectedSince := fields[4]
			parsedTime, err := time.Parse(time.DateTime, connectedSince)
			if err != nil {
				return nil, err
			}

			y := vpn_dtos.ServerConnection{
				CommonName:     commonName,
				Address:        realAddress,
				BytesRec:       bytesReceived,
				BytesSent:      bytesSent,
				ConnectedSince: parsedTime,
			}
			connections = append(connections, y)
		}
	}

	return &connections, nil

}
