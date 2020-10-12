package spiritworldtcp

import (
	"fmt"
	"net"
	"sync"
)

const BUF_LEN = 4096

const STOP_STRING = "No output"
const BEACON_RECV = "BEACON"
const BEACON_RESPONSE = "ROGER"
const BEACON_DONE = "DONE"

const DELIMETER = ">="

var cmdQueue []string

var mutex = &sync.Mutex{}

func QueueCMD(cmd string) {
	mutex.Lock()
	cmdQueue = append(cmdQueue, cmd)
	mutex.Unlock()
}

func handleConnection(client net.Conn) {
	// fmt.Printf("Server %s\n", client.RemoteAddr().String())

	var buf [BUF_LEN]byte
	_, err := client.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	// data := string(buf[:n])
	// fmt.Println(data)

	// Send response : ROGER or a CMD
	mutex.Lock()
	if len(cmdQueue) == 0 {
		client.Write([]byte(string(BEACON_RESPONSE)))
	} else {
		sendString := cmdQueue[0]
		cmdQueue = cmdQueue[1:]
		client.Write([]byte(string(sendString)))
	}
	mutex.Unlock()

	// Read the output
	_, err = client.Read(buf[:])
	if err != nil {
		fmt.Println(err)
		return
	}
	// data = string(buf[:n])
	// fmt.Println(data)

	// Cuts the beacon off
	client.Write([]byte(string(BEACON_DONE)))

}

func StartServer(port int) {
	PORT := ":" + fmt.Sprint(port)

	listener, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	//fmt.Println("Server Starting ...")

	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(client)
	}

}
