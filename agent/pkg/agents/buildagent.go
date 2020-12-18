package agents

import (
	"net"
	"os"
	"runtime"

	"github.com/google/uuid"
)

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

func BuildAgent() Agent {

	// Get IP info
	ips := getIPs()
	primaryIP := getPrimaryIP(ips)

	// Get hostname info
	hostname, _ := os.Hostname()

	// Get MAC
	macAddr, _ := getMacAddr()

	agent := Agent{
		AgentName:    AGENTNAME,
		AgentVersion: AGENTVERSION,
		UUID:         uuid.New().String(),
		PrimaryIP:    primaryIP,
		Hostname:     hostname,
		MAC:          macAddr[0],
		AgentOS:      runtime.GOOS,
		OtherIPs:     ips,
	}

	return agent
}
