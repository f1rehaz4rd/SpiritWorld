package database

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/agents"

	_ "github.com/lib/pq" // used to access postgres
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "redteam"
	PASSWORD = "changeme123"
	DBNAME   = "agents"
)

type DatabaseModel struct {
	db    *sql.DB
	mutex *sync.Mutex
}

type AgentBeaconModel struct {
	RegisterTime   string   `json:"register_time"`
	LastBeaconTime string   `json:"last_beacon_time"`
	ActionQueue    []string `json:"action_queue"`
	Actions        []string `json:"actions"`
}

type AgentModel struct {
	AgentObj *agents.Agent     `json:"agent"`
	Publicip string            `json:"public_ip"`
	Beacon   *AgentBeaconModel `json:"beacon"`
}

type ActionModel struct {
	ActionType   string `json:"action_type"`
	ActionCmd    string `json:"action_cmd"`
	ActionOutput string `json:"action_output"`
	UUID         string `json:"UUID"`
	AgentUUID    string `json:"agentUUID"`
}

type GroupModel struct {
	GroupName   string   `json:"group_name"`
	AgentsUUIDs []string `json:"agent_UUIDs`
}

func (model *DatabaseModel) Open() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME)

	var err error
	model.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = model.db.Ping()
	if err != nil {
		panic(err)
	}

	model.mutex = &sync.Mutex{}
}

func (model *DatabaseModel) Close() {
	model.db.Close()
}
