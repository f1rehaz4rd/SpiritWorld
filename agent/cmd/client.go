package main

import (
	"fmt"
	"time"

	"github.com/f1rehaz4rd/SpiritWorld/agent/pkg/agents"
	"github.com/f1rehaz4rd/SpiritWorld/agent/pkg/beacons/tcp"
)

func main() {

	agent := agents.BuildAgent()

	if tcp.RegisterAgent(&agent) {
		fmt.Println("Sucessfully registered")
		for {
			check := tcp.Beacon(&agent)
			if check {
				fmt.Println("successful heartbeat")
			} else {
				fmt.Println("problem beaconing")
			}

			time.Sleep(3 * time.Second)
		}
	}

}
