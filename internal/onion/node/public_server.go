package node

import (
	"hve/onion-simulate/internal/onion/packet"
	"hve/onion-simulate/internal/onion/types"
)

func MakePublicServer(id types.NodeId, network *types.OnionNetwork) PublicServer {
	return PublicServer{
		id:       id,
		packetCh: make(chan *packet.FastPacket, 32),
		circuit:  make([]uint32, 0),
		network:  network,
	}
}

type PublicServer struct {
	id       types.NodeId
	packetCh chan *packet.FastPacket
	circuit  []uint32
	network  *types.OnionNetwork
}

func (op *PublicServer) Simulate() {

}

func (op *PublicServer) GetId() types.NodeId {
	return op.id
}

func (op *PublicServer) Receive(packet *packet.FastPacket) {
	op.packetCh <- packet
}

func (op *PublicServer) GetNetwork() *types.OnionNetwork {
	return op.network
}
