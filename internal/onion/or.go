package onion

import (
	"hve/onion-simulate/internal/types"
	"log"
	"time"
)

type OnionRouter struct {
	id          types.NodeId
	packetCh    chan Packet
	circuitLink CircuitLink
	network     *OnionNetwork
}

func NewOR(id types.NodeId, network *OnionNetwork) *OnionRouter {
	return &OnionRouter{
		id:          id,
		packetCh:    make(chan Packet, 1),
		circuitLink: MakeCircuitLink(),
		network:     network,
	}
}

func (rt *OnionRouter) GetId() types.NodeId {
	return rt.id
}

func (rt *OnionRouter) Receive(packet *Packet) {
	rt.packetCh <- *packet
}

func (rt *OnionRouter) GetNetwork() *OnionNetwork {
	return rt.network
}

func (rt *OnionRouter) Simulate() {
	for {
		packet := <-rt.packetCh

		time.Sleep(time.Duration(packet.latency) * time.Millisecond)

		switch packet.cmd {
		case Create:
			pkt := Packet{
				cmd:    Created,
				circId: packet.circId,
				from:   rt,
				to:     packet.from,
			}
			pkt.Send()
		case Created:
			back := rt.circuitLink.Back[packet.circId]
			pkt := Packet{
				cmd:    Extended,
				circId: back.id,
				from:   rt,
				to:     back.node,
			}
			pkt.Send()
		case Extend:
			forward, ok := rt.circuitLink.Forward[packet.circId]
			if ok {
				// Relay Extend 전달
				pkt := Packet{
					cmd:     Extend,
					circId:  forward.id,
					from:    rt,
					to:      forward.node,
					payload: packet.payload,
				}
				pkt.Send()
			} else {
				// Create 전달
				payload, ok := packet.payload.(CreatePayload)
				if ok {
					circId := rt.circuitLink.Link(packet.circId, packet.from, payload.node)
					pkt := Packet{
						cmd:    Create,
						circId: circId,
						from:   rt,
						to:     payload.node,
					}
					pkt.Send()
				} else {
					log.Printf("[%d] Invalid cell : No payload in Extend Cell", rt.id)
				}
			}
		case Extended:
			result := packet.Relay(rt.circuitLink.Back)
			result.Send()
		case Relay:
			result := packet.Relay(rt.circuitLink.Back)
			result.Send()
		case Data:
		case Begin:
		case End:
		case Teardown:
		case Connected:
		}
	}
}
