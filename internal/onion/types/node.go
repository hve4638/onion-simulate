package types

type OnionNode interface {
	GetId() NodeId
	Receive(packet Packet)
	GetNetwork() *OnionNetwork
	NextHop(receivedNode OnionNode, circuitId uint32) (OnionNode, uint32)
	Simulate()
}
