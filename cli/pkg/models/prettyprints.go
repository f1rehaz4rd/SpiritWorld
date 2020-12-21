package models

import "fmt"

func PrintAgents(agents []AgentModel) {

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

	fmt.Println()
}

func PrintAgent(agent AgentModel) {
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
	fmt.Printf("\tAction Queue: \n\t%s\n\n", agent.Beacon.ActionQueue)
	fmt.Printf("\tActions: \n\t%s\n\n", agent.Beacon.Actions)
}

func PrintActions(actions []ActionModel) {
	fmt.Printf("%-36s | %-36s | %-13s | %-21s\n",
		"ACTION UUID",
		"AGENT UUID",
		"ACTION TYPE",
		"ACTION CMD")

	dashHolder := "------------------------------------------------"
	fmt.Printf("%-37s+%-37s+%-13s+%-21s\n",
		dashHolder[:37],
		dashHolder[:38],
		dashHolder[:15],
		dashHolder[:23])

	for i := 0; i < len(actions); i++ {
		fmt.Printf("%-36s | %-36s | %-13s | %s\n",
			actions[i].UUID,
			actions[i].AgentUUID,
			actions[i].ActionType,
			actions[i].ActionCmd)
	}

	fmt.Println()
}

func PrintAction(action ActionModel) {
	fmt.Printf("%s\n\n", action.UUID)
	fmt.Printf("\tAgent UUID: %s\n", action.AgentUUID)
	fmt.Printf("\tAction Type: %s\n", action.ActionType)
	fmt.Printf("\tAction Command: %s\n", action.ActionCmd)
	fmt.Printf("\tAction Output: \n%s\n", action.ActionOutput)
}

func PrintGroups(groups []GroupModel) {
	for i := 0; i < len(groups); i++ {
		fmt.Printf("%s\n", groups[i].GroupName)
		fmt.Printf("-------------------\n")
		for j := 0; j < len(groups[i].AgentsUUIDs); j++ {
			fmt.Printf("%s\n", groups[i].AgentsUUIDs[j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrintGroup(group GroupModel) {
	fmt.Printf("%s\n", group.GroupName)
	fmt.Printf("-------------------\n")
	for j := 0; j < len(group.AgentsUUIDs); j++ {
		fmt.Printf("%s\n", group.AgentsUUIDs[j])
	}
	fmt.Println()
}
