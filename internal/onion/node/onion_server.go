package node

import (
	cl "hve/onion-simulate/internal/onion/circuit_link"
	"hve/onion-simulate/internal/onion/packet"
	"hve/onion-simulate/internal/onion/routine"
	"hve/onion-simulate/internal/onion/types"
	"math/rand"
)

type OnionServer struct {
	id          types.NodeId
	packetCh    chan *packet.FastPacket
	circuitLink cl.CircuitLink
	routines    routine.Routine
	network     types.OnionNetwork
	rand        rand.Rand

	circuitChs    map[uint32](*chan Packet)
	nextCircuitId uint32
}

func NewOnionServer(id types.NodeId, network types.OnionNetwork, seed int64) OnionServer {
	return OnionServer{
		id:            id,
		packetCh:      make(chan *packet.FastPacket, 32),
		circuitChs:    make(map[uint32](*chan Packet)),
		routines:      make([]routine.Routine, 0),
		rand:          *rand.New(rand.NewSource(seed)),
		network:       network,
		nextCircuitId: 0,
	}
}

func (op *OnionServer) Simulate() {

}
