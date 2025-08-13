package node

import "hve/onion-simulate/internal/onion/types"

type DirectoryServer struct {
	network *types.OnionNetwork
}

func NewDirectoryServer(network *types.OnionNetwork) *DirectoryServer {
	return &DirectoryServer{
		network: network,
	}
}

func (ds *DirectoryServer) Simulate() {

}
