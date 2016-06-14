package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
)

func readCmd(c net.Conn) ([]byte, error) {
	bytesRead, err := bufio.NewReader(c).ReadBytes([]byte("\n")[0])
	if err != nil {
		c.Close()
		return make([]byte, 0), err
	}
	return bytesRead, nil
}

func main() {

	var cmdMap map[string]string
	var reply string

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// loop over connections (so that the connection can be reestablished
	// after a client exits/disconnects)
	for {
		// accept a new connection on the port we're listening on
		conn, _ := ln.Accept()
		// run loop forever (or until ctrl-c or a client disconnect is detected)
		for {
			// attempt to read a new message from the socket
			bytesRead, err := readCmd(conn)
			// if we didn't read anything, they break out of inner loop and
			// wait for a new connection
			if err != nil {
				break
			}
			err = json.Unmarshal(bytesRead, &cmdMap)
			if err != nil {
				reply = strings.ToUpper(fmt.Sprintf("%v", err))
			} else {
				// output message received
				fmt.Printf("Command Received: %+v\n", cmdMap)
				// sample process for string received
				reply = strings.ToUpper(string(bytesRead))
			}
			// send new string back to client
			conn.Write([]byte(reply + "\n"))
		}
	}
}
