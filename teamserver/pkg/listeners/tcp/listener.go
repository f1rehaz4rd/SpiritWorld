package tcp

import (
	"fmt"
	"net"
	"os"

	"github.com/f1rehaz4rd/SpiritWorld/teamserver/pkg/database"
)

// SERVADDR is the server port number
var SERVADDR = "0.0.0.0:"

var db database.DatabaseModel

//
// handleConnection from the TCP listener
//
func handleConnection(conn net.Conn) {
	defer conn.Close()

	msg := make([]byte, 4096)
	n, _ := conn.Read(msg)

	if !HandleMessage(msg[:n], conn) {
		fmt.Println("Error with connection: " + conn.RemoteAddr().String())
	}

}

//
// StartListener starts a tcp4 listener on the set port.
//
func StartListener() {
	SERVADDR += os.Getenv("TCP_LISTENER_PORT")

	listener, err := net.Listen("tcp", SERVADDR)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Starting TCP listener on " + SERVADDR)
	defer listener.Close()

	db.Open()
	defer db.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn)
	}

}
