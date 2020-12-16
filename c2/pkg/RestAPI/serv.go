package RestAPI

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/database"

	"github.com/gorilla/mux"
)

var db database.DatabaseModel

func getAgents(w http.ResponseWriter, r *http.Request) {
	agents, _ := db.AllAgents()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(agents)
}

func getAgent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	agent, err := db.GetAgentByID(params["id"])
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprint(w, nil)
	}

	json.NewEncoder(w).Encode(agent)
}

func getActions(w http.ResponseWriter, r *http.Request) {
	actions, _ := db.AllActions()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actions)
}

func getAction(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	action, err := db.GetAgentByID(params["id"])

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Fprint(w, nil)
	}

	json.NewEncoder(w).Encode(action)
}

func createAction(w http.ResponseWriter, r *http.Request) {

}

func Start() {
	// Init Router
	r := mux.NewRouter()

	db.Open()
	defer db.Close()

	r.HandleFunc("/api/agents", getAgents).Methods("GET")
	r.HandleFunc("/api/agent/{id}", getAgent).Methods("GET")
	r.HandleFunc("/api/actions", getActions).Methods("GET")
	r.HandleFunc("/api/actions/{id}", getAction).Methods("GET")
	r.HandleFunc("/api/actions", createAction).Methods("POST")

	fmt.Println("Starting RestAPI on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
