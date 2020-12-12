package agents

type Agent struct {
	AgentName    string   `json:"agent_name"`
	AgentVersion string   `json:"agent_version"`
	UUID         string   `json:"UUID"`
	PrimaryIP    string   `json:"primary_ip"`
	Hostname     string   `json:"hostname"`
	MAC          string   `json:"MAC"`
	AgentOS      string   `json:"agent_os"`
	OtherIPs     []string `json:"other_ips"`
	APIKEY       string   `json:"API_KEY"`
}

type Action struct {
	ActionType   string `json:"action_type"`
	ActionOutput string `json:"action_output"`
}

type AgentBeacon struct {
	Agent  *Agent  `json: "agent"`
	Action *Action `json: "action"`
}
