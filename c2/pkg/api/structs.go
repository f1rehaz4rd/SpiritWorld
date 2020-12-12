package api

import (
	"time"

	"github.com/f1rehaz4rd/SpiritWorld/c2/pkg/agents"
)

type RegisterAgent struct {
	RegisterTime time.Time     `json:"register_time"`
	PublicIP     string        `json:"public_ip"`
	Agent        *agents.Agent `json:"agent"`
}

type BeaconAgent struct {
	BeaconTime time.Time     `json:"beacon_time"`
	PublicIP   string        `json:"public_ip"`
	Agent      *agents.Agent `json:"agent"`
}
