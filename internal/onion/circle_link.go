package onion

type CircuitLink struct {
	nextCircId uint32
	Forward    map[uint32]CircuitChain
	Back       map[uint32]CircuitChain
}

type CircuitChain struct {
	id   uint32
	node OnionNode
}

func MakeCircuitLink() CircuitLink {
	return CircuitLink{
		nextCircId: 0,
		Forward:    make(map[uint32]CircuitChain),
		Back:       make(map[uint32]CircuitChain),
	}
}

func (link *CircuitLink) Link(circId uint32, from OnionNode, to OnionNode) uint32 {
	newCircId := link.nextCircId
	link.nextCircId++

	link.Forward[circId] = CircuitChain{newCircId, to}
	link.Back[newCircId] = CircuitChain{circId, from}
	return newCircId
}
