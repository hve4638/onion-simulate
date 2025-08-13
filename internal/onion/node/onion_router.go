package node

import (
	cl "hve/onion-simulate/internal/onion/circuit_link"
	"hve/onion-simulate/internal/onion/types"
	"log"
)

type OnionRouter struct {
	id          types.NodeId
	packetCh    chan *FastPacket
	circuitLink cl.CircuitLink
	setting     SimulationSetting
}

func NewOnionRouter(id types.NodeId, setting SimulationSetting) *OnionRouter {
	return &OnionRouter{
		id:          id,
		packetCh:    make(chan *FastPacket, 1),
		circuitLink: cl.NewCircuitLink(),
		setting:     setting,
	}
}

func (rt *OnionRouter) GetId() types.NodeId {
	return rt.id
}

func (rt *OnionRouter) Receive(packet *FastPacket) {
	rt.packetCh <- packet
}

func (rt *OnionRouter) Send(packet *FastPacket) {
	if packet.To == nil {
		log.Printf("[%d] Packet has no destination node", rt.id)
		return
	}

	// packet.To.
	packet.To.Receive(packet)
}

func (rt *OnionRouter) Simulate() {
	for {
		// time.Sleep(time.Duration(packet.Latency) * time.Millisecond)
		pkt := <-rt.packetCh

		switch pkt.Type {
		case types.Create:
			res := pkt.MakeCreated()
			rt.Send(&res)
		case types.Created:
			res := pkt.MakeExtended()
			rt.Send(&res)
		case types.RelayExtend:
			res := pkt.MakeExtended()
			rt.Send(&res)
		case types.RelayExtended:
			res := pkt.MakeCreate()
			rt.Send(&res)
		default:
			log.Println("unknown packet type")
		}
	}
}

// func (rt *OnionRouter) RelayForward(circId uint32, destNode types.OnionNode) packet.FastPacket {
// 	rt.circuitLink.Link(circId, nil, destNode)

// 	return packet.FastPacket{
// 		cmd:     types.Create,
// 		circId:  circId,
// 		from:    rt,
// 		to:      destNode,
// 		payload: CreatePayload{},
// 		latency: 0,
// 	}
// }

// func (rt *OnionRouter) RelayBack(pkt *Packet) packet.FastPacket {
// 	return packet.FastPacket{
// 		cmd:     types.Relay,
// 		circId:  next.id,
// 		from:    pkt.to,
// 		to:      next.node,
// 		payload: pkt,
// 	}
// }
