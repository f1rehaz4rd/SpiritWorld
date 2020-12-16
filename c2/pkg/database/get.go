package database

import (
	"database/sql"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/agents"

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
		var uuid,
			agentname,
			agentversion,
			primaryip,
			hostname,
			mac,
			agentos,
			otherips,
			publicip string

		err = rows.Scan(
			&uuid,
			&agentname,
			&agentversion,
			&primaryip,
			&hostname,
			&mac,
			&agentos,
			&otherips,
			&publicip)

		if err != nil {
			return nil, err
		}

		tmpAgent := &agents.Agent{
			AgentName:    agentname,
			AgentVersion: agentversion,
			UUID:         uuid,
			PrimaryIP:    primaryip,
			Hostname:     hostname,
			MAC:          mac,
			AgentOS:      agentos,
			OtherIPs:     otherips,
		}

		tmp := AgentModel{
			AgentObj: tmpAgent,
			Publicip: publicip,
		}

		allAgents = append(allAgents, tmp)

	}

	for i := 0; i < len(allAgents); i++ {
		sqlStatement = `SELECT registertime, lastbeacon, actionqueue FROM agentbeacon WHERE uuid=$1;`

		model.mutex.Lock()
		row := model.db.QueryRow(sqlStatement, allAgents[i].AgentObj.UUID)
		model.mutex.Unlock()

		var registertime, lastbeacon, actionqueue string
		switch err := row.Scan(&registertime, &lastbeacon, &actionqueue); err {
		case sql.ErrNoRows:
			break
		case nil:
			tmp := &AgentBeaconModel{
				RegisterTime:   registertime,
				LastBeaconTime: lastbeacon,
				Actionqueue:    actionqueue,
			}
			allAgents[i].Beacon = tmp
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

	var uuid,
		agentname,
		agentversion,
		primaryip,
		hostname,
		mac,
		agentos,
		otherips,
		publicip string

	switch err := row.Scan(
		&uuid,
		&agentname,
		&agentversion,
		&primaryip,
		&hostname,
		&mac,
		&agentos,
		&otherips,
		&publicip); err {
	case sql.ErrNoRows:
		break
	case nil:
		tmpAgent := &agents.Agent{
			AgentName:    agentname,
			AgentVersion: agentversion,
			UUID:         uuid,
			PrimaryIP:    primaryip,
			Hostname:     hostname,
			MAC:          mac,
			AgentOS:      agentos,
			OtherIPs:     otherips,
		}

		agent = AgentModel{
			AgentObj: tmpAgent,
			Publicip: publicip,
		}

	default:
		return agent, err
	}

	sqlStatement = `SELECT registertime, lastbeacon, actionqueue FROM agentbeacon WHERE uuid=$1;`

	row = model.db.QueryRow(sqlStatement, id)

	var registertime, lastbeacon, actionqueue string
	switch err := row.Scan(&registertime, &lastbeacon, &actionqueue); err {
	case sql.ErrNoRows:
		break
	case nil:
		tmp := &AgentBeaconModel{
			RegisterTime:   registertime,
			LastBeaconTime: lastbeacon,
			Actionqueue:    actionqueue,
		}
		agent.Beacon = tmp
		break
	default:
		return agent, err
	}

	return agent, nil
}

func (model *DatabaseModel) AllActions() ([]agents.Action, error) {
	var allActions []agents.Action

	sqlStatement := `SELECT * FROM actions;`

	model.mutex.Lock()
	rows, err := model.db.Query(sqlStatement)
	model.mutex.Unlock()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var uuid,
			actiontype,
			actioncmd,
			actionresponse string

		err = rows.Scan(
			&uuid,
			&actiontype,
			&actioncmd,
			&actionresponse)

		if err != nil {
			return nil, err
		}

		tmpAction := agents.Action{
			UUID:         uuid,
			ActionType:   actiontype,
			ActionCmd:    actioncmd,
			ActionOutput: actionresponse,
		}

		allActions = append(allActions, tmpAction)

	}

	return allActions, nil
}

func (model *DatabaseModel) GetActionByID(id string) (agents.Action, error) {
	var action agents.Action

	sqlStatement := `SELECT * FROM actions WHERE uuid=$1;`

	model.mutex.Lock()
	row := model.db.QueryRow(sqlStatement, id)
	model.mutex.Unlock()

	var uuid,
		actiontype,
		actioncmd,
		actionresponse string

	switch err := row.Scan(
		&uuid,
		&actiontype,
		&actioncmd,
		&actionresponse); err {
	case sql.ErrNoRows:
		break
	case nil:
		action = agents.Action{
			UUID:         uuid,
			ActionType:   actiontype,
			ActionCmd:    actioncmd,
			ActionOutput: actionresponse,
		}
		break
	default:
		return action, err
	}

	return action, nil
}
