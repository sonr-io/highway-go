package p2p

import (
	"context"
	"errors"

	"github.com/kataras/golog"
	"github.com/libp2p/go-libp2p-core/crypto"
	libp2pHost "github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	common "github.com/sonr-io/sonr/x/registry/types"

	"google.golang.org/protobuf/proto"

	ps "github.com/libp2p/go-libp2p-pubsub"
)

var (
	logger              = golog.Default.Child("pkg/p2p")
	ErrMissingParam     = errors.New("Paramater is missing.")
	ErrRoutingNotSet    = errors.New("DHT and Host have not been set by Routing Function")
	ErrListenerRequired = errors.New("Listener was not Provided")
	ErrMDNSInvalidConn  = errors.New("Invalid Connection, cannot begin MDNS Service")
)

type HostImpl interface {
	// AuthenticateMessage authenticates a message
	AuthenticateMessage(msg proto.Message, metadata *common.Metadata) bool

	// Close closes the node
	Close()

	// Connect to a peer
	Connect(pi peer.AddrInfo) error

	// Did returns the DID of the node
	Did() string

	// HasRouting returns true if the node has routing
	HasRouting() error

	// HostID returns the ID of the Host
	HostID() peer.ID

	// Join subsrcibes to a topic
	Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error)

	// NewStream opens a new stream to a peer
	NewStream(ctx context.Context, pid peer.ID, pids ...protocol.ID) (network.Stream, error)

	// NewTopic creates a new pubsub topic with event handler and subscription
	NewTopic(topic string, opts ...ps.TopicOpt) (*ps.Topic, *ps.TopicEventHandler, *ps.Subscription, error)

	// NeedsWait checks if state is Resumed or Paused and blocks channel if needed
	NeedsWait()

	// Pause tells all of goroutines to pause execution
	Pause()

	// Ping sends a ping to a peer to check if it is alive
	Ping(did string) error

	// Publish publishes a message to a topic
	Publish(topic string, msg proto.Message, metadata *common.Metadata) error

	// Resume tells all of goroutines to resume execution
	Resume()

	// Role returns the role of the node
	Role() Role

	// Router returns the routing.Router
	Router(h libp2pHost.Host) (routing.PeerRouting, error)

	// SendMessage sends a message to a peer
	SendMessage(id peer.ID, p protocol.ID, data proto.Message) error

	// SetStreamHandler sets the handler for a protocol
	SetStreamHandler(protocol protocol.ID, handler network.StreamHandler)

	// SignData signs the data with the private key
	SignData(data []byte) ([]byte, error)

	// SignMessage signs a message with the node's private key
	SignMessage(message proto.Message) ([]byte, error)

	// VerifyData verifies the data signature
	VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool
}

type host struct {
	// Standard Node Implementation
	HostImpl
	libp2pHost.Host
	mode Role

	// Host and context
	connection common.Connection

	privKey      crypto.PrivKey
	mdnsPeerChan chan peer.AddrInfo
	dhtPeerChan  <-chan peer.AddrInfo

	// Properties
	ctx    context.Context
	pubKey crypto.PubKey
	did    string
	*dht.IpfsDHT
	*ps.PubSub

	// State
	flag   uint64
	Chn    chan bool
	status HostStatus
}

// NewHost Initializes a libp2p host to be used with the Sonr Highway and Motor nodes
func NewHost(ctx context.Context, privKey crypto.PrivKey, options ...Option) (HostImpl, error) {
	// Initialize DHT
	opts := defaultOptions(Role_MOTOR)
	node, err := opts.Apply(ctx, privKey, options...)
	if err != nil {
		return nil, err
	}

	// Initialize Discovery for MDNS
	node.createMdnsDiscovery(opts)
	node.SetStatus(Status_READY)
	go node.Serve()

	// Open Store with profileBuf
	return node, nil
}

func (n *host) Did() string {
	return n.did
}

// HostID returns the ID of the Host
func (n *host) HostID() peer.ID {
	return n.Host.ID()
}

// Ping sends a ping to the peer
func (n *host) Ping(pid string) error {
	return nil
}

// Publish publishes a message to the network
func (n *host) Publish(t string, message proto.Message, metadata *common.Metadata) error {
	return nil
}

// Role returns the role of the node
func (n *host) Role() Role {
	return n.mode
}
