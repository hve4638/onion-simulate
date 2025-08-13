package simulator

import (
	"fmt"
	id_generator "hve/onion-simulate/internal/onion/id_generator"
	"hve/onion-simulate/internal/onion/node"
	"hve/onion-simulate/internal/onion/routine"
	"hve/onion-simulate/internal/onion/types"
	onion_log "hve/onion-simulate/internal/onion_log"
	"math/rand"
)

type OnionNetworkSimulator struct {
	seed       int64
	globalRand rand.Rand
	logs       onion_log.ConcurrentLog
	timer      *onion_log.Timer
	idGen      id_generator.IdGenerator
	routines   map[string]routine.Routine
	// dns        map[string]*ServerOP

	routerCount int

	routers []*node.OnionRouter
	users   []*node.OnionUser
	servers []*any
}

func (o *OnionNetworkSimulator) NextRand() int64 {
	return o.globalRand.Int63()
}

func (o *OnionNetworkSimulator) GetSeed() int64 {
	return o.seed
}

func (o *OnionNetworkSimulator) GenerateId() uint32 {
	return o.idGen.Next()
}

func (o *OnionNetworkSimulator) Simulate() {
	o.timer.Reset()

	for _, node := range o.users {
		go node.Simulate()
	}
	for _, node := range o.routers {
		go node.Simulate()
	}
	// for _, node := range o.servers {
	// 	go node.Simulate()
	// }
}

// func (o *OnionNetworkSimulator) Init(config *types.Config) {
// 	// o.addServers(config.Server)
// 	// o.addRoutines(config.Routines)
// 	o.addUsers(config.User)
// 	o.addRouters(config.User)
// }

func (o *OnionNetworkSimulator) KnownServer(domain string) types.OnionNode {
	return nil
}

// func (o *OnionNetworkSimulator) addServers(serverConfig []types.ServerConfig) {
// 	proxies := make([]*ServerOP, 0)
// 	dns := make(map[string]*ServerOP)
// 	for _, data := range serverConfig {
// 		id := o.idGen.Next()
// 		op := MakeServerOP(types.NodeId(id), o)

// 		proxies = append(proxies, op)
// 		dns[data.Id] = op
// 	}

// 	o.dns = dns
// 	o.Servers = proxies
// }

// func (o *OnionNetworkSimulator) addRoutines(routines []types.RoutineConfig) {
// 	for _, config := range routines {
// 		o.addRoutine(config)
// 	}
// }

func (o *OnionNetworkSimulator) addRoutine(config types.RoutineConfig) {
	// endpoint, ok := o.dns[config.URL]
	// if !ok {
	// 	panicf("Initialize routine error: server not found: %s", config.URL)
	// }
	// minMs := config.Period.TimeRange.Min.Hour
	// minMs = minMs*60 + config.Period.TimeRange.Min.Minute
	// minMs = minMs*60 + config.Period.TimeRange.Min.Second
	// minMs = minMs * 1000

	// maxMs := config.Period.TimeRange.Max.Hour
	// maxMs = maxMs*60 + config.Period.TimeRange.Max.Minute
	// maxMs = maxMs*60 + config.Period.TimeRange.Max.Second
	// maxMs = maxMs * 1000
	// fmt.Printf("min: %d, max: %d\n", minMs, maxMs)
	// fmt.Printf("config : %v\n", config)

	// routine := Routine{
	// 	Name:           config.Id,
	// 	RepeatCount:    config.RepeatCountRange,
	// 	RepeatInterval: config.RepeatIntervalRange,
	// 	Period: Period{
	// 		week: Week(config.Period.Week),
	// 		time: types.Int64Range{Min: minMs, Max: maxMs},
	// 	},
	// 	Endpoint:         endpoint,
	// 	CommunicateCount: config.CommunicateCountRange,
	// }

	// o.routines[config.Id] = append(o.routines[config.Id], routine)
}

// func (o *OnionNetworkSimulator) addUsers(userConfig []types.UserConfig) {
// 	for _, data := range userConfig {
// 		id := types.NodeId(o.idGen.Next())
// 		op := node.NewOnionUser(id, o, o.NextRand())
// 		// for _, routineName := range data.Routines {
// 		// 	routine, ok := o.routines[routineName]
// 		// 	if !ok {
// 		// 		warnf("[Warning] routine not found: '%s'\n", routineName)
// 		// 		println("[Warning] user: ", routine)
// 		// 	}
// 		// 	op.routines = append(op.routines, routine...)
// 		// }

// 		o.users = append(o.users, op)
// 	}
// }

// func (o *OnionNetworkSimulator) addRouters(userConfig []types.UserConfig) {
// 	for o.idGen.HasNext() {
// 		id := types.NodeId(o.idGen.Next())
// 		op := node.NewOnionRouter(id, o)
// 		o.routers = append(o.routers, op)
// 	}
// }

func (o *OnionNetworkSimulator) ShowOption() {
	fmt.Println("Routers count : ", len(o.routers))
	fmt.Println("Servers count : ", len(o.servers))
	fmt.Println("Users count   : ", len(o.users))
}
