package agents

const AGENTNAME = "Spirit Agent"
const AGENTVERSION = "v 0.1"
const AGENTAPIKEY = "bc43b40c-3e5f-11eb-b378-0242ac130002"

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
