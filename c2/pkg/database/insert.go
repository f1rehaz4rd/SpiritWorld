package database

import (
	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/api"

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
		agent.OtherIPs,
		register.PublicIP)
	if err != nil {
		return false
	}
	model.mutex.Unlock()

	sqlStatment = `
	INSERT INTO agentbeacon (
		uuid,
		registertime,
		lastbeacon,
		actionqueue
	) VALUES ($1, $2, $3, $4)
	`

	_, err = model.db.Exec(sqlStatment,
		agent.UUID,
		register.RegisterTime,
		register.RegisterTime,
		"")
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
