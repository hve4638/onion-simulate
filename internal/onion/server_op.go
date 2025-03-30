package onion

import "hve/onion-simulate/internal/types"

func MakeServerOP(id types.NodeId, network *OnionNetwork) *ServerOP {
	return &ServerOP{
		id:       id,
		packetCh: make(chan Packet, 32),
		circuit:  make([]uint32, 0),
		routines: make([]Routine, 0),
		network:  network,
	}
}

type ServerOP struct {
	id       types.NodeId
	packetCh chan Packet
	circuit  []uint32
	routines Routines
	network  *OnionNetwork
}

func (op *ServerOP) Simulate() {

}

func (op *ServerOP) GetId() types.NodeId {
	return op.id
}

func (op *ServerOP) Receive(packet *Packet) {
	op.packetCh <- *packet
}

func (op *ServerOP) GetNetwork() *OnionNetwork {
	return op.network
}
