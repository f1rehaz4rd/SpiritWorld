package database

import (
	"database/sql"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/agents"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func (model *DatabaseModel) AllAgents() ([]AgentModel, error) {
	var allAgents []AgentModel

	sqlStatement := `SELECT * FROM agent;`

	model.mutex.Lock()
	rows, err := model.db.Query(sqlStatement)
	model.mutex.Unlock()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		var agentModel AgentModel
		var tmpAgent agents.Agent
		err = rows.Scan(
			&tmpAgent.UUID,
			&tmpAgent.AgentName,
			&tmpAgent.AgentVersion,
			&tmpAgent.PrimaryIP,
			&tmpAgent.Hostname,
			&tmpAgent.MAC,
			&tmpAgent.AgentOS,
			pq.Array(&tmpAgent.OtherIPs),
			&agentModel.Publicip)

		if err != nil {
			return nil, err
		}

		agentModel.AgentObj = &tmpAgent

		allAgents = append(allAgents, agentModel)

	}

	for i := 0; i < len(allAgents); i++ {
		sqlStatement = `SELECT * WHERE uuid=$1;`

		model.mutex.Lock()
		row := model.db.QueryRow(sqlStatement, allAgents[i].AgentObj.UUID)
		model.mutex.Unlock()

		var beaconModel AgentBeaconModel
		switch err := row.Scan(
			&beaconModel.RegisterTime,
			&beaconModel.LastBeaconTime,
			pq.Array(&beaconModel.ActionQueue),
			pq.Array(&beaconModel.Actions)); err {
		case sql.ErrNoRows:
			break
		case nil:
			allAgents[i].Beacon = &beaconModel
			break
		default:
			return nil, err
		}

	}

	return allAgents, nil
}

func (model *DatabaseModel) GetAgentByID(id string) (AgentModel, error) {
	var agent AgentModel

	sqlStatement := `SELECT * FROM agent WHERE uuid=$1;`

	model.mutex.Lock()
	row := model.db.QueryRow(sqlStatement, id)
	model.mutex.Unlock()

	var tmpAgent agents.Agent

	switch err := row.Scan(
		&tmpAgent.UUID,
		&tmpAgent.AgentName,
		&tmpAgent.AgentVersion,
		&tmpAgent.PrimaryIP,
		&tmpAgent.Hostname,
		&tmpAgent.MAC,
		&tmpAgent.AgentOS,
		pq.Array(&tmpAgent.OtherIPs),
		&agent.Publicip); err {
	case sql.ErrNoRows:
		break
	case nil:
		agent.AgentObj = &tmpAgent
	default:
		return agent, err
	}

	sqlStatement = `SELECT registertime,
		 lastbeacon,
		 actionqueue,
		 actions
		 FROM agentbeacon WHERE uuid=$1;`

	row = model.db.QueryRow(sqlStatement, id)
	var beaconModel AgentBeaconModel
	switch err := row.Scan(
		&beaconModel.RegisterTime,
		&beaconModel.LastBeaconTime,
		pq.Array(&beaconModel.ActionQueue),
		pq.Array(&beaconModel.Actions)); err {
	case sql.ErrNoRows:
		break
	case nil:
		agent.Beacon = &beaconModel
		break
	default:
		return agent, err
	}

	return agent, nil
}

func (model *DatabaseModel) AllActions() ([]ActionModel, error) {
	var allActions []ActionModel

	sqlStatement := `SELECT * FROM actions;`

	model.mutex.Lock()
	rows, err := model.db.Query(sqlStatement)
	model.mutex.Unlock()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var actionModel ActionModel
	for rows.Next() {

		err = rows.Scan(
			&actionModel.UUID,
			&actionModel.AgentUUID,
			&actionModel.ActionCmd,
			&actionModel.ActionOutput)

		if err != nil {
			return nil, err
		}

		allActions = append(allActions, actionModel)

	}

	return allActions, nil
}

func (model *DatabaseModel) GetActionByID(id string) (ActionModel, error) {

	sqlStatement := `SELECT * FROM actions WHERE uuid=$1;`

	model.mutex.Lock()
	row := model.db.QueryRow(sqlStatement, id)
	model.mutex.Unlock()

	var actionModel ActionModel
	switch err := row.Scan(
		&actionModel.UUID,
		&actionModel.AgentUUID,
		&actionModel.ActionCmd,
		&actionModel.ActionOutput); err {
	case sql.ErrNoRows:
		break
	case nil:
		break
	default:
		return actionModel, err
	}

	return actionModel, nil
}
