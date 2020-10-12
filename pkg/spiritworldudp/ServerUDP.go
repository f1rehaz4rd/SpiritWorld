package spiritworldudp

import (
	"fmt"
	"net"
	"os"
)

const BEACON_INCOMING = "Beacon"
const BEACON_RESPONSE = "Roger"
const MAXBUF = 4096

type clientMsg struct {
	msg   string
	raddr net.UDPAddr
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}

func ProcessData(data []byte, addr *net.UDPAddr) {

}

func handleClient(conn *net.UDPConn) {

}

func ServerSetup(serverAddr string, port int) {

	//
	// Sets up the UDP Server
	//
	udpAddr := net.UDPAddr{
		Port: port,
		IP:   net.ParseIP(serverAddr),
	}

	conn, err := net.ListenUDP("udp", &udpAddr)
	checkError(err)

	fmt.Println("Listening for data ... ")
	for {
		handleClient(conn)
	}
}
