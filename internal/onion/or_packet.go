package onion

func (or *OnionRouter) MakeCreatePacket(circId uint32, destNode OnionNode) Packet {
	or.circuitLink.Link(circId, nil, destNode)

	return Packet{
		cmd:    Create,
		circId: circId,
		from:   or,
		to:     destNode,
	}
}

func (or *OnionRouter) MakeExtendPacket(circId uint32, destNode OnionNode) Packet {
	foward := or.circuitLink.Forward[circId]

	return Packet{
		cmd:    Extend,
		circId: circId,
		from:   or,
		to:     foward.node,
		payload: CreatePayload{
			node: destNode,
		},
	}
}
