package types

import (
	onion_log "hve/onion-simulate/internal/onion_log"
	"math/rand"
)

type SimulationSetting struct {
	globalRand *rand.Rand
	logs       onion_log.ConcurrentLog
	timer      *onion_log.Timer
	routines   map[string]Routines
	dns        map[string]*OnionNode

	Routers []*OnionNode
	Users   []*OnionNode
	Servers []*OnionNode
}
