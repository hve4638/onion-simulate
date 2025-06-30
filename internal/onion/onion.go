package onion

import (
	"fmt"
	onion_log "hve/onion-simulate/internal/onion_log"
	"hve/onion-simulate/internal/types"
	"math/rand"
)

type TimeInterval uint32

type OnionNetwork struct {
	globalRand *rand.Rand
	logs       onion_log.ConcurrentLog
	timer      *onion_log.Timer
	idGen      IdGenerator
	routines   map[string]Routines
	dns        map[string]*ServerOP

	Routers []*OnionRouter
	Users   []*UserOP
	Servers []*ServerOP
}

func NewOnionNetwork(size uint32, seed int64) OnionNetwork {
	timer := onion_log.NewTimer()

	return OnionNetwork{
		globalRand: rand.New(rand.NewSource(seed)),
		logs:       onion_log.NewConcurrentLog(int(size/10), &timer),
		timer:      &timer,
		idGen:      MakeIdCounter(size, int64(seed)),
		dns:        make(map[string]*ServerOP),
		routines:   make(map[string]Routines),

		Routers: make([]*OnionRouter, 0),
		Users:   make([]*UserOP, 0),
		Servers: make([]*ServerOP, 0),
	}
}

func (o *OnionNetwork) Simulate() {
	o.timer.Reset()

	for _, node := range o.Routers {
		go node.Simulate()
	}
	for _, node := range o.Servers {
		go node.Simulate()
	}
	for _, node := range o.Users {
		go node.Simulate()
	}
}

func (o *OnionNetwork) Init(config *types.Config) {
	o.addServers(config.Server)
	o.addRoutines(config.Routines)
	o.addUsers(config.User)
	o.addRouters(config.User)
}

func (o *OnionNetwork) addServers(serverConfig []types.ServerConfig) {
	proxies := make([]*ServerOP, 0)
	dns := make(map[string]*ServerOP)
	for _, data := range serverConfig {
		id := o.idGen.Next()
		op := MakeServerOP(types.NodeId(id), o)

		proxies = append(proxies, op)
		dns[data.Id] = op
	}

	o.dns = dns
	o.Servers = proxies
}

func (o *OnionNetwork) addRoutines(routines []types.RoutineConfig) {
	for _, config := range routines {
		o.addRoutine(config)
	}
}

func (o *OnionNetwork) addRoutine(config types.RoutineConfig) {
	endpoint, ok := o.dns[config.URL]
	if !ok {
		panicf("Initialize routine error: server not found: %s", config.URL)
	}
	minMs := config.Period.TimeRange.Min.Hour
	minMs = minMs*60 + config.Period.TimeRange.Min.Minute
	minMs = minMs*60 + config.Period.TimeRange.Min.Second
	minMs = minMs * 1000

	maxMs := config.Period.TimeRange.Max.Hour
	maxMs = maxMs*60 + config.Period.TimeRange.Max.Minute
	maxMs = maxMs*60 + config.Period.TimeRange.Max.Second
	maxMs = maxMs * 1000
	fmt.Printf("min: %d, max: %d\n", minMs, maxMs)
	fmt.Printf("config : %v\n", config)

	routine := Routine{
		Name:           config.Id,
		RepeatCount:    config.RepeatCountRange,
		RepeatInterval: config.RepeatIntervalRange,
		Period: Period{
			week: Week(config.Period.Week),
			time: types.Int64Range{Min: minMs, Max: maxMs},
		},
		Endpoint:         endpoint,
		CommunicateCount: config.CommunicateCountRange,
	}

	o.routines[config.Id] = append(o.routines[config.Id], routine)
}

func (o *OnionNetwork) addUsers(userConfig []types.UserConfig) {
	for _, data := range userConfig {
		id := types.NodeId(o.idGen.Next())
		op := NewUserOP(id, o)
		for _, routineName := range data.Routines {
			routine, ok := o.routines[routineName]
			if !ok {
				warnf("[Warning] routine not found: '%s'\n", routineName)
				println("[Warning] user: ", routine)
			}
			op.routines = append(op.routines, routine...)
		}

		o.Users = append(o.Users, op)
	}
}

func (o *OnionNetwork) addRouters(userConfig []types.UserConfig) {
	for o.idGen.HasNext() {
		id := types.NodeId(o.idGen.Next())
		op := NewOR(id, o)
		o.Routers = append(o.Routers, op)
	}
}

func (o *OnionNetwork) ShowOption() {
	fmt.Println("Routers count : ", len(o.Routers))
	fmt.Println("Servers count : ", len(o.Servers))
	fmt.Println("Users count   : ", len(o.Users))
}
