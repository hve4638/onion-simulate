package node

import (
	cl "hve/onion-simulate/internal/onion/circuit_link"
	"hve/onion-simulate/internal/onion/routine"
	"hve/onion-simulate/internal/onion/types"
	"log"
	"math/rand"
)

type OnionUser struct {
	id          types.NodeId
	packetCh    chan *FastPacket
	circuitLink cl.CircuitLink
	routines    routine.Routine
	network     types.OnionNetwork
	rand        rand.Rand

	circuitChs    map[uint32](*chan Packet)
	nextCircuitId uint32
}

func NewOnionUser(id types.NodeId, network types.OnionNetwork, seed int64) OnionUser {
	return OnionUser{
		id:       id,
		packetCh: make(chan *FastPacket, 32),

		// circuit:    MakeCircuitLink(),
		circuitChs: make(map[uint32](*chan Packet)),
		routines:   make([]routine.Routine, 0),
		rand:       *rand.New(rand.NewSource(seed)),
		network:    network,
	}
}

func (op *OnionUser) GetId() types.NodeId {
	return op.id
}

func (op *OnionUser) Send(packet *FastPacket) {
	if packet.To == nil {
		log.Printf("[%d] Packet has no destination node", op.id)
		return
	}

	packet.To.Receive(packet)
}
func (op *OnionUser) Receive(packet *FastPacket) {
	op.packetCh <- packet
}

func (op *OnionUser) GetNetwork() *types.OnionNetwork {
	return &op.network
}

func (op *OnionUser) Simulate() {
	for {
		// time.Sleep(time.Duration(packet.Latency) * time.Millisecond)
		pkt := <-op.packetCh

		switch pkt.Type {
		case types.Create:
			res := pkt.MakeCreated()
			op.Send(&res)
			break
		case types.Created:
			res := pkt.MakeExtended()
			op.Send(&res)
			break
		case types.RelayExtend:
			res := pkt.MakeExtended()
			op.Send(&res)
			break
		case types.RelayExtended:
			res := pkt.MakeCreate()
			op.Send(&res)
			break
		default:
			log.Println("unknown packet type")
		}
	}
}
