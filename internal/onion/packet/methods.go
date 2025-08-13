package packet

import (
	types "hve/onion-simulate/internal/onion/types"
)

func (pkt *FastPacket) GetLatency() types.TimeInterval {
	return pkt.Latency
}

// Create 패킷에 대한 Created 응답 생성
func (pkt *FastPacket) MakeCreated() FastPacket {
	return FastPacket{
		CircId:  pkt.CircId,
		From:    pkt.To,
		To:      pkt.From,
		Latency: pkt.Latency,
	}
}

// Relay 데이터를 바탕으로 하는 Extend 패킷 생성
func (pkt *FastPacket) MakeCreate() FastPacket {
	return FastPacket{
		CircId:  pkt.CircId,
		From:    pkt.To,
		To:      pkt.From,
		Latency: pkt.Latency,
	}
}

func (pkt *FastPacket) MakeExtended() FastPacket {
	return FastPacket{
		CircId:  pkt.CircId,
		From:    pkt.To,
		To:      pkt.From,
		Latency: pkt.Latency,
	}
}
