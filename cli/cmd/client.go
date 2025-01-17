package main

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/f1rehaz4rd/SpiritWorld/cli/pkg/agents"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "ListAgents", Description: "List all of the agents"},
		{Text: "GetAgent", Description: "<AGENT ID> | Get more information about an agent"},
		{Text: "ListActions", Description: "List all the actions"},
		{Text: "GetAction", Description: "<ACTION ID> | Get more information about an action"},
		{Text: "CreateAction", Description: "<AGENT ID> <ACTION TYPE> <COMMAND> | Create action for an agent"},
		{Text: "CreateGroupAction", Description: "<GROUP NAME> <ACTION TYPE> <COMMAND> | Create action for an agent"},
		{Text: "ListGroups", Description: "List all of the groups"},
		{Text: "GetGroup", Description: "<GROUP NAME> | Get more information about the group"},
		{Text: "CreateGroup", Description: "<GROUP NAME> | Create a group"},
		{Text: "AddToGroup", Description: "<GROUP NAME> <AGENT ID> | Add agent to group"},
		{Text: "RemoveFromGroup", Description: "<GROUP NAME> <AGENT ID> | Remove agent from group"},
		// {Text: "help", Description: "Show all the commands"},
		{Text: "exit", Description: "Exits the cli"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Println("Welcome to Spirit World Command and Control")
	fmt.Println("Use ctrl+l to clear the terminal window")

	exitFlag := false
	for {
		t := prompt.Input("spiritcli$ ", completer)

		tmp := strings.Split(t, " ")

		if len(tmp) == 0 {
			continue
		}

		cmd := tmp[0]

		switch cmd {
		case "ListAgents":
			agents.ListAgents()
		case "GetAgent":
			if len(tmp) < 2 {
				fmt.Println("Incorrect arguments.")
				continue
			}
			agents.GetAgent(tmp[1])
		case "ListActions":
			agents.ListActions()
		case "GetAction":
			if len(tmp) < 2 {
				fmt.Println("Incorrect arguments.")
				continue
			}
			agents.GetAction(tmp[1])
		case "CreateAction":
			if len(tmp) < 4 {
				fmt.Println("Incorrect arguments.")
				continue
			}
			agents.CreateAction(tmp[1], tmp[2], strings.Join(tmp[3:], " "))
		case "CreateGroupAction":
			if len(tmp) < 4 {
				fmt.Println("Incorrect arguments.")
				continue
			}
			agents.CreateGroupAction(tmp[1], tmp[2], strings.Join(tmp[3:], " "))
		case "ListGroups":
			agents.ListGroups()
		case "GetGroup":
			if len(tmp) < 2 {
				fmt.Println("Incorrect arguments.")
				continue
			}
			agents.GetGroup(strings.Join(tmp[1:], " "))
		case "CreateGroup":
			if len(tmp) < 2 {
				fmt.Println("Incorrect arguments.")
				continue
			}
			agents.CreateGroup(strings.Join(tmp[1:], " "))
		case "AddToGroup":
			if len(tmp) < 3 {
				fmt.Println("Incorrect arguments.")
				continue
			}
			agents.AddToGroup(tmp[1], tmp[2])
		case "RemoveFromGroup":
			if len(tmp) < 3 {
				fmt.Println("Incorrect arguments.")
				continue
			}
			agents.RemoveFromGroup(tmp[1], tmp[2])
		case "exit":
			exitFlag = true
		default:
			fmt.Println("Invalid command")
		}

		if exitFlag {
			break
		}
	}
}
