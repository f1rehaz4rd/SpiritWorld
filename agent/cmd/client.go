package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/f1rehaz4rd/SpiritWorld/agent/pkg/agents"
	"github.com/google/uuid"
)

const servAddr = "127.0.0.1:4321"
const initString = "REGISTER"

const AGENTNAME = "Spirit Test"
const AGENTVERSION = "v 0.1"
const AGENTAPIKEY = "ASDFASDFASDFASDF"

var agent agents.Agent

func registerAgent() {

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
		return
	}

	msg, _ := json.Marshal(register)
	conn.Write(msg)

	conn.Close()
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
		OtherIPs:     ips,
		APIKEY:       AGENTAPIKEY,
	}

}

func main() {

	registerAgent()

}
