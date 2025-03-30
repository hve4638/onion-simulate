package onion

import "hve/onion-simulate/internal/types"

type OnionNode interface {
	GetId() types.NodeId
	Receive(packet *Packet)
	GetNetwork() *OnionNetwork
}
