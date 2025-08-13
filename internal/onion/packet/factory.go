package packet

import types "hve/onion-simulate/internal/onion/types"

func NewCreatePacket(srcNode types.OnionNode, targetNode types.OnionNode) FastPacket {
	return FastPacket{
		Type:    types.Create,
		From:    srcNode,
		To:      targetNode,
		Latency: types.TimeInterval(0),
	}
}

func NewExtendPacket(
	srcNode types.OnionNode,
	prevNode types.OnionNode,
	circId uint32,
	targetNode types.OnionNode,
) FastPacket {
	nextNode, nextCircId := srcNode.NextHop(prevNode, circId)

	return FastPacket{
		Type:                 types.RelayExtend,
		CircId:               nextCircId,
		From:                 srcNode,
		To:                   nextNode,
		DestinationForExtend: targetNode,
	}
}
