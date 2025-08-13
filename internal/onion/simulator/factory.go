package simulator

import (
	"hve/onion-simulate/internal/onion/node"
	"hve/onion-simulate/internal/onion/routine"
	onion_log "hve/onion-simulate/internal/onion_log"
	"math/rand"
)

// uint32(config.Node.Amount),
func NewOnionSimulator(seed int64) *OnionNetworkSimulator {
	timer := onion_log.NewTimer()

	simulator := OnionNetworkSimulator{
		seed:       seed,
		globalRand: *rand.New(rand.NewSource(seed)),
		// logs:       onion_log.NewConcurrentLog(int(size/10), &timer),
		timer: &timer,
		// idGen:    id_generator.NewIdGenerator(size, int64(seed)),
		routines: make(map[string]routine.Routine),

		routers: make([]*node.OnionRouter, 0),
		users:   make([]*node.OnionUser, 0),
		servers: make([]*any, 0),
	}

	return &simulator
}
