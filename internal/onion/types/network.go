package types

type OnionNetwork interface {
	NextRand() int64
	GetSeed() int64
	GenerateId() uint32

	KnownServer(domain string) OnionNode
	// Routers() []OnionNode
	// Users() []OnionNode
	// Servers() []OnionNode

	Simulate()
}
