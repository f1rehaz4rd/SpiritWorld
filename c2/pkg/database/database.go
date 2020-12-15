package database

import (
	"database/sql"
	"fmt"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/agents"
	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/api"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "redteam"
	password = "changeme123"
	dbname   = "agents"
)

type DatabaseModel struct {
	db *sql.DB
}

type AgentBeaconModel struct {
	registerTime   string
	lastBeaconTime string
	actionqueue    string
}

type AgentModel struct {
	agentObj *agents.Agent
	publicip string
	beacon   *AgentBeaconModel
}

func (model *DatabaseModel) Open() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	model.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = model.db.Ping()
	if err != nil {
		panic(err)
	}

}

func (model *DatabaseModel) Close() {
	model.db.Close()
}

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
	_, err := model.db.Exec(sqlStatment,
		beacon.Agent.UUID,
		beacon.BeaconTime)
	if err != nil {
		return false
	}

	return true
}

func (model *DatabaseModel) AllAgents() ([]AgentModel, error) {
	var allAgents []AgentModel

	sqlStatement := `SELECT * FROM agent;`

	rows, err := model.db.Query(sqlStatement)
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
			agentObj: tmpAgent,
			publicip: publicip,
		}

		allAgents = append(allAgents, tmp)

	}

	for i := 0; i < len(allAgents); i++ {
		sqlStatement = `SELECT registertime, lastbeacon, actionqueue FROM agentbeacon WHERE uuid=$1;`

		row := model.db.QueryRow(sqlStatement, allAgents[i].agentObj.UUID)

		var registertime, lastbeacon, actionqueue string
		switch err := row.Scan(&registertime, &lastbeacon, &actionqueue); err {
		case sql.ErrNoRows:
			return nil, err
		case nil:
			tmp := &AgentBeaconModel{
				registerTime:   registertime,
				lastBeaconTime: lastbeacon,
				actionqueue:    actionqueue,
			}
			allAgents[i].beacon = tmp
			break
		default:
			return nil, err
		}

	}

	return allAgents, nil
}
