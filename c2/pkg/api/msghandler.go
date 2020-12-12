package api

import (
	"encoding/json"
	"fmt"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/agents"
)

func HandleMessage(msg []byte, publicIP string) bool {

	var agentBeacon agents.AgentBeacon

	if err := json.Unmarshal(msg, &agentBeacon); err != nil {
		fmt.Println(err)
		return false
	}

	switch agentBeacon.Action.ActionType {
	case "register":
		agentRegister(agentBeacon.Agent, publicIP)
		break
	case "beacon":
		agentBeaconing(agentBeacon.Agent, publicIP)
		break
	default:
		break
	}

	return true
}

func agentRegister(agent *agents.Agent, publicIP string) {

	// register := &RegisterAgent{
	// 	RegisterTime: time.Now(),
	// 	PublicIP:     publicIP,
	// 	Agent:        agent,
	// }

	fmt.Println("Registering: " + publicIP)

}

func agentBeaconing(agent *agents.Agent, publicIP string) {

	// beacon := &BeaconAgent{
	// 	BeaconTime: time.Now(),
	// 	PublicIP:   publicIP,
	// 	Agent:      agent,
	// }

	fmt.Println("Updating: " + publicIP)

}
