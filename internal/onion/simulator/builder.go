package simulator

import (
	id_generator "hve/onion-simulate/internal/onion/id_generator"
	"hve/onion-simulate/internal/onion/node"
	"hve/onion-simulate/internal/onion/types"
)

func (o *OnionNetworkSimulator) setConfig(config types.Config) {
	o.initIdGenerator(config)
	o.initUserAmount(config)
	o.SetRouterAmount(config.Node.Amount, config.Node.AmountExit)
	o.SetServerAmount(len(config.Server))
	o.SetExitNodeAmount(config.Node.AmountExit)

	// for _, server := range config.Server {
	// 	id := types.NodeId(o.idGen.Next())
	// 	op := node.NewOnionServer(id, o, o.NextRand(), server.Domain)
	// 	o.servers = append(o.servers, &op)
	// }
}

func (o *OnionNetworkSimulator) initIdGenerator(config types.Config) {
	routerCount := uint32(config.Node.Amount + config.Node.AmountExit)
	serverCount := uint32(config.Node.Amount + config.Node.AmountExit)
	userCount := 0

	for _, user := range config.User {
		userCount += user.Amount
	}

	o.idGen = id_generator.NewIdGenerator(routerCount, o.seed)
}

func (o *OnionNetworkSimulator) initUserAmount(data types.Config) {
	o.users = make([]*node.OnionUser, data.Node.Amount)

	for _, data := range o.userConfig {
		id := types.NodeId(o.idGen.Next())
		op := node.NewOnionUser(id, o, o.NextRand())
		// for _, routineName := range data.Routines {
		// 	routine, ok := o.routines[routineName]
		// 	if !ok {
		// 		warnf("[Warning] routine not found: '%s'\n", routineName)
		// 		println("[Warning] user: ", routine)
		// 	}
		// 	op.routines = append(op.routines, routine...)
		// }

		o.users = append(o.users, op)
	}
}

func (o *OnionNetworkSimulator) SetRouterAmount(amount int, exitAmount int) {
	o.routers = make([]*node.OnionRouter, amount)
}

func (o *OnionNetworkSimulator) SetServerAmount(amount int) {

}

func (o *OnionNetworkSimulator) SetExitNodeAmount(amount int) {

}
