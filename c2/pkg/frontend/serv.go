package frontend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/database"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/index.html")
}

func agentsPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend/agents.html")
}

func getallagents(w http.ResponseWriter, r *http.Request) {
	var db database.DatabaseModel
	db.Open()
	defer db.Close()

	agents, _ := db.AllAgents()

	for i := 0; i < len(agents); i++ {
		tmp, _ := json.Marshal(agents[i])
		fmt.Fprintf(w, string(tmp)+"\n")
	}
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/agents", agentsPage)
	http.HandleFunc("/getallagents", getallagents)
}

func StartHTTPServer() {
	setupRoutes()
	fmt.Println("Starting HTTP Server on 0.0.0.0:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
