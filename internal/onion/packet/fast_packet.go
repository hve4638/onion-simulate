package packet

import (
	types "hve/onion-simulate/internal/onion/types"
)

type FastPacket struct {
	Type    types.PacketType
	CircId  uint32
	From    types.OnionNode
	To      types.OnionNode
	Latency types.TimeInterval

	DestinationForExtend types.OnionNode
}
