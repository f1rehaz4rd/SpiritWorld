package main

import (
	"fmt"
	"net"

	spiritworld "github.com/f1rehaz4rd/SpiritWorld/pkg/spiritworld"
)

const PORT = ":45123"
const maxBuf = 4096

// queue := make([]string, 0, 100)

func main() {

	//
	// Sets up the UDP Server
	//
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()

	buffer := make([]byte, maxBuf)

	for {

		//
		// Listening for UDP
		//
		fmt.Println("Listening for data ... ")
		n, remoteAddr, _ := connection.ReadFromUDP(buffer)

		data := string(buffer[0:n])
		fmt.Println("-> " + data)

		//
		// Send off to process the data
		//
		go spiritworld.ProcessData(data, remoteAddr)
	}

}
