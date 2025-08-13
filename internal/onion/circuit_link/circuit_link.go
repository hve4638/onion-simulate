package circuit_link

import "hve/onion-simulate/internal/onion/types"

type OnionNode = types.OnionNode
type NodeId = types.NodeId
type CircuitId = uint32

// bi-directional link between circuits
type CircuitLink struct {
	nextCircId CircuitId
	linked     map[NodeId]map[CircuitId]CircuitChain
}

type CircuitChain struct {
	// 서킷 ID
	CircId CircuitId
	// 연결된 노드
	Node OnionNode
}

func NewCircuitLink() CircuitLink {
	return CircuitLink{
		nextCircId: 0,
		linked:     make(map[NodeId]map[CircuitId]CircuitChain),
	}
}

// 새 서킷 링크 생성
func (cl *CircuitLink) Link(circId uint32, from OnionNode, to OnionNode) uint32 {
	newCircId := cl.nextCircId
	cl.nextCircId++

	fromChain := CircuitChain{newCircId, from}
	toChain := CircuitChain{circId, to}
	cl.addChain(fromChain, toChain)
	cl.addChain(toChain, fromChain)
	return newCircId
}

func (cl *CircuitLink) addChain(from CircuitChain, to CircuitChain) {
	forwardNodeId := from.Node.GetId()
	cl.linked[forwardNodeId][from.CircId] = to
}

func (cl *CircuitLink) Unlink(circId CircuitId, node OnionNode) bool {
	nodeId := node.GetId()
	forwardChains, ok := cl.linked[nodeId]
	if !ok {
		return false
	}

	backward := forwardChains[circId]
	backwardNodeId := backward.Node.GetId()
	backwardChains, ok := cl.linked[backwardNodeId]
	if !ok {
		return false
	}

	delete(forwardChains, circId)
	delete(backwardChains, backward.CircId)

	return true
}
