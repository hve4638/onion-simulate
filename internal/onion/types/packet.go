package types

type Packet interface {
	// GetCmd() cmd.OnionCMD
	// GetLatency() TimeInterval
	// CircId  uint32
	// From    *OnionNode
	// To      *OnionNode
	// Payload any
	// Latency TimeInterval
}

type PacketType uint8

const (
	Create PacketType = iota
	Created
	RelayExtend
	RelayExtended
	RelayData
	RelayTruncate
	RelayTruncated
	// Teardown
	// Connected

	PlainData
)
