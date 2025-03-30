package onion

func (op *UserOP) MakeCreatePacket(circId uint32, destNode OnionNode) Packet {
	op.circuit.Link(circId, nil, destNode)

	return Packet{
		cmd:    Create,
		circId: circId,
		from:   op,
		to:     destNode,
	}
}

func (op *UserOP) MakeExtendPacket(circId uint32, destNode OnionNode) Packet {
	foward := op.circuit.Forward[circId]

	return Packet{
		cmd:    Extend,
		circId: circId,
		from:   op,
		to:     foward.node,
		payload: CreatePayload{
			node: destNode,
		},
	}
}
