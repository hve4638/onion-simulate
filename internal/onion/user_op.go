package onion

import (
	"hve/onion-simulate/internal/types"
	"math/rand"
)

type UserOP struct {
	id       types.NodeId
	packetCh chan Packet
	circuit  CircuitLink
	routines Routines
	network  *OnionNetwork
	rand     *rand.Rand

	circuitChs    map[uint32](*chan Packet)
	nextCircuitId uint32
}

func NewUserOP(id types.NodeId, network *OnionNetwork) *UserOP {
	return &UserOP{
		id:         id,
		packetCh:   make(chan Packet, 32),
		circuit:    MakeCircuitLink(),
		circuitChs: make(map[uint32](*chan Packet)),
		routines:   make([]Routine, 0),
		rand:       rand.New(rand.NewSource(network.globalRand.Int63())),
		network:    network,
	}
}

func (op *UserOP) GetId() types.NodeId {
	return op.id
}

func (op *UserOP) Receive(packet *Packet) {
	op.packetCh <- *packet
}

func (op *UserOP) GetNetwork() *OnionNetwork {
	return op.network
}
