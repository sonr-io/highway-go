package models

import (
	"context"
	"net/http"

	"github.com/sonr-io/core/node"
	"github.com/tendermint/starport/starport/pkg/cosmosclient"
	hw "go.buf.build/grpc/go/sonr-io/core/highway/v1"

	"google.golang.org/grpc"
)

// HighwayStub is the RPC Service for the Custodian Node.
type HighwayStub struct {
	hw.HighwayServer
	Host   node.HostImpl
	Cosmos cosmosclient.Client

	// Properties
	Ctx  context.Context
	Grpc *grpc.Server
	Http *http.Server

	// Configuration
}

//get
// no clear answer

//give
//did

//TODO this needs work, remove soon
type Jwt struct {
	Snr        string `json:"snr"`
	EthAddress string `json: "ethAddress"`
}
