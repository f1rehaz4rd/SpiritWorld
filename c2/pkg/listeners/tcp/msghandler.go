package tcp

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/agents"
	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/api"
	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/database"
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

	register := api.RegisterAgent{
		RegisterTime: time.Now(),
		PublicIP:     conn.RemoteAddr().String(),
		Agent:        agent,
	}

	var db database.DatabaseModel
	db.Open()
	defer db.Close()

	if !db.InsertAgent(register) {
		fmt.Println("Failed to register: " + conn.RemoteAddr().String())
		return false
	}

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

	beacon := api.BeaconAgent{
		BeaconTime: time.Now(),
		PublicIP:   conn.RemoteAddr().String(),
		Agent:      agent,
	}

	var db database.DatabaseModel
	db.Open()
	defer db.Close()

	if !db.UpdateAgent(beacon) {
		fmt.Println("Failed to update: " + conn.RemoteAddr().String())
		return false
	}

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
