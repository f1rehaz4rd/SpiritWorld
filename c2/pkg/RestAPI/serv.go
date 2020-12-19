package RestAPI

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/database"
	"github.com/google/uuid"

	"github.com/gorilla/mux"
)

var db database.DatabaseModel

func getAgents(w http.ResponseWriter, r *http.Request) {
	agents, err := db.AllAgents()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(agents)
	}
}

func getAgent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	agent, err := db.GetAgentByID(params["id"])
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(agent)
	}
}

func getActions(w http.ResponseWriter, r *http.Request) {
	actions, err := db.AllActions()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(actions)
	}
}

func getAction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	action, err := db.GetActionByID(params["id"])

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(action)
	}
}

func getGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := db.AllGroups()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(groups)
	}
}

func getGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	group, err := db.GetGroupByID(params["name"])
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(group)
	}
}

func createAction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	action := database.ActionModel{
		ActionType: params["type"],
		ActionCmd:  params["cmd"],
		UUID:       uuid.New().String(),
		AgentUUID:  params["id"],
	}

	w.Header().Set("Content-Type", "application/json")
	if !db.InsertAction(action) {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(action)
	}
}

func createGroupAction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	check := db.InsertGroupAction(params["name"], params["type"], params["cmd"])
	if !check {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(check)
	}
}

func createGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	group := database.GroupModel{
		GroupName: params["name"],
	}

	w.Header().Set("Content-Type", "application/json")
	if !db.CreateGroup(group) {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(group)
	}
}

func addToGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	check := db.InsertIntoGroup(params["name"], params["id"])
	if !check {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(check)
	}
}

func removeFromGroup(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	check := db.RemoveFromGroup(params["name"], params["id"])
	if !check {
		fmt.Fprint(w, nil)
	} else {
		json.NewEncoder(w).Encode(check)
	}
}

func Start() {
	// Init Router
	r := mux.NewRouter()

	db.Open()
	defer db.Close()

	r.HandleFunc("/api/agents", getAgents).Methods("GET")
	r.HandleFunc("/api/agent/{id}", getAgent).Methods("GET")
	r.HandleFunc("/api/actions", getActions).Methods("GET")
	r.HandleFunc("/api/action/{id}", getAction).Methods("GET")
	r.HandleFunc("/api/groups", getGroups).Methods("GET")
	r.HandleFunc("/api/group/{name}", getGroup).Methods("GET")
	r.HandleFunc("/api/createaction/{id}/{type}/{cmd}", createAction).Methods("POST")
	r.HandleFunc("/api/creategroupaction/{name}/{type}/{cmd}", createGroupAction).Methods("POST")
	r.HandleFunc("/api/creategroup/{name}", createGroup).Methods("POST")
	r.HandleFunc("/api/addtogroup/{name}/{id}", addToGroup).Methods("POST")
	r.HandleFunc("/api/removefromgroup/{name}/{id}", removeFromGroup).Methods("POST")

	fmt.Println("Starting RestAPI on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
