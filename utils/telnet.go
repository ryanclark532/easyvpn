package utils

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)


func ConnectTelnet(address string) (net.Conn, error){ 
	return net.Dial("tcp", address)
}


func CommandTelnet( command string, conn net.Conn) error {
	writer := bufio.NewWriter(conn)
	_, err := writer.WriteString(command + "\n")
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}


func ReadTelnet(conn net.Conn)([]string, error){
	reader := bufio.NewReader(conn)
	ret := []string{}
	for {
		data, err := reader.ReadString('\n')
	if err == io.EOF {
 	   // Connection closed
  	  break
	} else if err != nil {
 	   fmt.Println("Error reading:", err)
  	  break
	}

		ret = append(ret, data)
		if(strings.TrimSpace(data) == "END"){
			break
		}
	}
	return ret, nil 
}