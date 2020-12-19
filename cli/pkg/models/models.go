package models

type Agent struct {
	AgentName    string   `json:"agent_name"`
	AgentVersion string   `json:"agent_version"`
	UUID         string   `json:"UUID"`
	PrimaryIP    string   `json:"primary_ip"`
	Hostname     string   `json:"hostname"`
	MAC          string   `json:"MAC"`
	AgentOS      string   `json:"agent_os"`
	OtherIPs     []string `json:"other_ips"`
}

type Action struct {
	ActionType   string `json:"action_type"`
	ActionCmd    string `json:"action_cmd"`
	ActionOutput string `json:"action_output"`
	UUID         string `json:"UUID"`
}

type AgentBeacon struct {
	Agent  *Agent  `json: "agent"`
	Action *Action `json: "action"`
	APIKEY string  `json: "apikey"`
}

type AgentBeaconModel struct {
	RegisterTime   string   `json:"register_time"`
	LastBeaconTime string   `json:"last_beacon_time"`
	ActionQueue    []string `json:"action_queue"`
	Actions        []string `json:"actions"`
}

type AgentModel struct {
	AgentObj *Agent            `json:"agent"`
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
