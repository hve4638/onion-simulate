package onion

type Packet struct {
	cmd     OnionCMD
	circId  uint32
	from    OnionNode
	to      OnionNode
	payload Payload
	latency TimeInterval
}

type Payload interface{}

type CreatePayload struct {
	node OnionNode
}

func (pkt *Packet) Send() {
	network := pkt.to.GetNetwork()
	network.logs.Add(uint32(pkt.from.GetId()), uint32(pkt.to.GetId()), network.timer.Now())

	pkt.to.Receive(pkt)
}

func (pkt *Packet) Relay(chain map[uint32]CircuitChain) Packet {
	next := chain[pkt.circId]

	return Packet{
		cmd:     Relay,
		circId:  next.id,
		from:    pkt.to,
		to:      next.node,
		payload: pkt,
	}
}
