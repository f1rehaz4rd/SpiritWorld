package database

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/agents"

	_ "github.com/lib/pq"
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
	RegisterTime   string
	LastBeaconTime string
	Actionqueue    string
}

type AgentModel struct {
	AgentObj *agents.Agent
	Publicip string
	Beacon   *AgentBeaconModel
}

//
//
// DELETE THIS!!!!!!!!!!!!!!!!!!
//
//
type ActionModel struct {
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
