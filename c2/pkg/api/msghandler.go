package api

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/agents"
)

func HandleMessage(msg []byte, conn net.Conn) bool {

	var agentBeacon agents.AgentBeacon

	if err := json.Unmarshal(msg, &agentBeacon); err != nil {
		fmt.Println(err)
		return false
	}

	switch agentBeacon.Action.ActionType {
	case "register":
		return agentRegister(agentBeacon.Agent, conn)
	case "beacon":
		return agentBeaconing(agentBeacon.Agent, conn)
	default:
		break
	}

	return true
}

func agentRegister(agent *agents.Agent, conn net.Conn) bool {

	// register := &RegisterAgent{
	// 	RegisterTime: time.Now(),
	// 	PublicIP:     conn.RemoteAddr().String(),
	// 	Agent:        agent,
	// }

	fmt.Println("Registering: " + conn.RemoteAddr().String())

	resp := &agents.Action{
		ActionType:   "register",
		ActionOutput: "success",
	}

	msg, _ := json.Marshal(resp)
	_, err := conn.Write(msg)
	if err != nil {
		return false
	}

	return true
}

func agentBeaconing(agent *agents.Agent, conn net.Conn) bool {

	// beacon := &BeaconAgent{
	// 	BeaconTime: time.Now(),
	// 	PublicIP:   conn.RemoteAddr().String(),
	// 	Agent:      agent,
	// }

	fmt.Println("Updating: " + conn.RemoteAddr().String())

	resp := &agents.Action{
		ActionType:   "beacon",
		ActionOutput: "success",
	}

	msg, _ := json.Marshal(resp)
	_, err := conn.Write(msg)
	if err != nil {
		return false
	}

	return true
}
