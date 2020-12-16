package agents

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/f1rehaz4rd/SpiritWorld/cli/pkg/models"
)

func ListAgents() {
	response, _ := http.Get("http://127.0.0.1:8080/api/agents")
	data, _ := ioutil.ReadAll(response.Body)

	var agents []models.AgentModel

	json.Unmarshal(data, &agents)

	fmt.Printf("%-16s | %-15s | %-37s | %-21s | %s\n",
		"AGENT NAME",
		"PRIMARY IP",
		"AGENT ID",
		"PUBLIC IP",
		"BEACON LAST SEEN")

	dashHolder := "------------------------------------------------"
	fmt.Printf("%-16s+%-15s+%-37s+%-21s+%s\n",
		dashHolder[:17],
		dashHolder[:17],
		dashHolder[:39],
		dashHolder[:23],
		dashHolder[:38])

	for i := 0; i < len(agents); i++ {
		fmt.Printf("%-16s | %-15s | %-37s | %-21s | %s\n",
			agents[i].AgentObj.AgentName,
			agents[i].AgentObj.PrimaryIP,
			agents[i].AgentObj.UUID,
			agents[i].Publicip,
			agents[i].Beacon.LastBeaconTime)
	}

}

func GetAgent(id string) {
	response, _ := http.Get("http://127.0.0.1:8080/api/agent/" + id)
	data, _ := ioutil.ReadAll(response.Body)

	var agent models.AgentModel

	json.Unmarshal(data, &agent)

	if agent.AgentObj == nil {
		fmt.Println("Agent ID is invalid")
		return
	}

	fmt.Printf("%s - %s\n\n", agent.AgentObj.AgentName, agent.AgentObj.UUID)
	fmt.Printf("\tAgentVersion: %s\n", agent.AgentObj.AgentVersion)
	fmt.Printf("\tPrimary IP: %s\n", agent.AgentObj.PrimaryIP)
	fmt.Printf("\tHostname: %s\n", agent.AgentObj.Hostname)
	fmt.Printf("\tMAC: %s\n", agent.AgentObj.MAC)
	fmt.Printf("\tOS: %s\n", agent.AgentObj.AgentOS)
	fmt.Printf("\tOther IPs: %s\n", agent.AgentObj.OtherIPs)
	fmt.Printf("\tPublic IP: %s\n\n", agent.Publicip)
	fmt.Printf("\tRegister Time: %s\n", agent.Beacon.RegisterTime)
	fmt.Printf("\tLast Beacon Time: %s\n", agent.Beacon.LastBeaconTime)
	fmt.Printf("\tAction Queue: \n%s\n", agent.Beacon.Actionqueue)

}
