package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/f1rehaz4rd/SpiritWorld/agent/pkg/agents"
	"github.com/google/uuid"
)

const servAddr = "127.0.0.1:4321"
const initString = "REGISTER"

const AGENTNAME = "Spirit Agent"
const AGENTVERSION = "v 0.1"
const AGENTAPIKEY = "bc43b40c-3e5f-11eb-b378-0242ac130002"

var agent agents.Agent

func registerAgent() bool {

	buildAgent()

	register := agents.AgentBeacon{
		Agent: &agent,
		Action: &agents.Action{
			ActionType:   "register",
			ActionOutput: "",
		},
	}

	conn, err := net.Dial("tcp", servAddr)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer conn.Close()

	msg, _ := json.Marshal(register)
	conn.Write(msg)

	n, _ := conn.Read(msg)

	var action agents.Action
	if err := json.Unmarshal(msg[:n], &action); err != nil {
		fmt.Println("failed register")
		return false
	}

	if action.ActionType == "register" &&
		action.ActionOutput == "success" {
		return true
	}

	return false
}

func getIPs() []string {

	var ips []string

	ifaces, _ := net.Interfaces()
	for _, i := range ifaces {
		addrs, _ := i.Addrs()
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			ips = append(ips, ip.String())
		}
	}

	return ips
}

func getPrimaryIP(ips []string) string {

	primaryIP := ""

	for _, ip := range ips {

		if len(ip) > 7 {
			tmp := ip[0:7]
			if tmp == "192.168" {
				primaryIP = ip
			}
		}

		if len(ip) > 3 {
			if ip[0:3] == "10." {
				primaryIP = ip
			}
		}

	}

	return primaryIP
}

func ipsToString(ips []string) string {
	var ipString string

	for i := 0; i < len(ips); i++ {
		ipString = ipString + ips[i] + ","
	}

	ipString = ipString[:len(ipString)-1]

	return ipString
}

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}

func buildAgent() {

	// Get IP info
	ips := getIPs()
	primaryIP := getPrimaryIP(ips)

	// Get hostname info
	hostname, _ := os.Hostname()

	// Get MAC
	macAddr, _ := getMacAddr()

	agent = agents.Agent{
		AgentName:    AGENTNAME,
		AgentVersion: AGENTVERSION,
		UUID:         uuid.New().String(),
		PrimaryIP:    primaryIP,
		Hostname:     hostname,
		MAC:          macAddr[0],
		AgentOS:      runtime.GOOS,
		OtherIPs:     ipsToString(ips),
		APIKEY:       AGENTAPIKEY,
	}

}

func beacon() bool {
	beaconObj := agents.AgentBeacon{
		Agent: &agent,
		Action: &agents.Action{
			ActionType:   "beacon",
			ActionOutput: "",
		},
	}

	conn, err := net.Dial("tcp", servAddr)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer conn.Close()

	msg, _ := json.Marshal(beaconObj)
	conn.Write(msg)

	n, _ := conn.Read(msg)

	var action agents.Action
	if err := json.Unmarshal(msg[:n], &action); err != nil {
		fmt.Println("failed to beacon")
		return false
	}

	switch action.ActionType {
	case "beacon":
		return action.ActionOutput == "success"
	default:
		break
	}

	return false
}

func main() {

	if registerAgent() {
		fmt.Println("Sucessfully registered")
		for {
			if beacon() {
				fmt.Println("successful heartbeat")
			} else {
				fmt.Println("problem beaconing")
			}
			time.Sleep(3 * time.Second)
		}
	}

}
