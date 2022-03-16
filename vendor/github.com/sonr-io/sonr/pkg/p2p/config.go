package p2p

import (
	"context"
	"fmt"
	"time"

	"github.com/libp2p/go-libp2p"
	connmgr "github.com/libp2p/go-libp2p-connmgr"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	dscl "github.com/libp2p/go-libp2p-core/discovery"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

// HostStatus is the status of the host
type HostStatus int

// SNRHostStatus Definitions
const (
	Status_IDLE       HostStatus = iota // Host is idle, default state
	Status_STANDBY                      // Host is standby, waiting for connection
	Status_CONNECTING                   // Host is connecting
	Status_READY                        // Host is ready
	Status_FAIL                         // Host failed to connect
	Status_CLOSED                       // Host is closed
)

// Equals returns true if given SNRHostStatus matches this one
func (s HostStatus) Equals(other HostStatus) bool {
	return s == other
}

// IsNotIdle returns true if the SNRHostStatus != Status_IDLE
func (s HostStatus) IsNotIdle() bool {
	return s != Status_IDLE
}

// IsStandby returns true if the SNRHostStatus == Status_STANDBY
func (s HostStatus) IsStandby() bool {
	return s == Status_STANDBY
}

// IsReady returns true if the SNRHostStatus == Status_READY
func (s HostStatus) IsReady() bool {
	return s == Status_READY
}

// IsConnecting returns true if the SNRHostStatus == Status_CONNECTING
func (s HostStatus) IsConnecting() bool {
	return s == Status_CONNECTING
}

// IsFail returns true if the SNRHostStatus == Status_FAIL
func (s HostStatus) IsFail() bool {
	return s == Status_FAIL
}

// IsClosed returns true if the SNRHostStatus == Status_CLOSED
func (s HostStatus) IsClosed() bool {
	return s == Status_CLOSED
}

// String returns the string representation of the SNRHostStatus
func (s HostStatus) String() string {
	switch s {
	case Status_IDLE:
		return "IDLE"
	case Status_STANDBY:
		return "STANDBY"
	case Status_CONNECTING:
		return "CONNECTING"
	case Status_READY:
		return "READY"
	case Status_FAIL:
		return "FAIL"
	case Status_CLOSED:
		return "CLOSED"
	}
	return "UNKNOWN"
}

var (
	bootstrapAddrStrs = []string{
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
		"/ip4/104.131.131.82/tcp/4001/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
		"/ip4/104.131.131.82/udp/4001/quic/p2p/QmaCpDMGvV2BGHeYERUEnRQAwe3N8SzbUtfsmvsqQLuvuJ",
	}
	addrStoreTTL = time.Minute * 5
)

// Role is the type of the node (Client, Highway)
type Role int

const (
	// StubMode_LIB is the Node utilized by Mobile and Web Clients
	Role_UNSPECIFIED Role = iota

	// StubMode_CLI is the Node utilized by CLI Clients
	Role_TEST

	// Role_MOTOR is for a Motor Node
	Role_MOTOR

	// Role_HIGHWAY is for a Highway Node
	Role_HIGHWAY
)

// Motor returns true if the node has a client.
func (m Role) IsMotor() bool {
	return m == Role_MOTOR
}

// Highway returns true if the node has a highway stub.
func (m Role) IsHighway() bool {
	return m == Role_HIGHWAY
}

// Prefix returns golog prefix for the node.
func (m Role) Prefix() string {
	var name string
	switch m {
	case Role_HIGHWAY:
		name = "highway"
	case Role_MOTOR:
		name = "motor"
	case Role_TEST:
		name = "test"
	default:
		name = "unknown"
	}
	return fmt.Sprintf("[SONR.%s] ", name)
}

// Option is a function that modifies the node options.
type Option func(*options)

// WithConnOptions sets the connection manager options. Defaults are (lowWater: 15, highWater: 40, gracePeriod: 5m)
func WithConnOptions(low int, hi int, grace time.Duration) Option {
	return func(o *options) {
		o.LowWater = low
		o.HighWater = hi
		o.GracePeriod = grace
	}
}

// WithRendevouz sets the rendevouz for the host. Default is 5 seconds.
func WithRendevouz(r string) Option {
	return func(o *options) {
		o.Rendezvous = r
	}
}

// WithTTL sets the ttl for the host. Default is 2 minutes.
func WithTTL(ttl time.Duration) Option {
	return func(o *options) {
		o.TTL = dscl.TTL(ttl)
	}
}

// options is a collection of options for the node.
type options struct {
	// Host
	BootstrapPeers []peer.AddrInfo
	LowWater       int
	HighWater      int
	GracePeriod    time.Duration
	MultiAddrs     []multiaddr.Multiaddr
	Rendezvous     string
	Interval       time.Duration
	TTL            dscl.Option
}

// defaultOptions returns the default options
func defaultOptions(r Role) *options {
	// Create Bootstrapper List
	var bootstrappers []multiaddr.Multiaddr
	for _, s := range bootstrapAddrStrs {
		ma, err := multiaddr.NewMultiaddr(s)
		if err != nil {
			continue
		}
		bootstrappers = append(bootstrappers, ma)
	}

	// Create Address Info List
	ds := make([]peer.AddrInfo, 0, len(bootstrappers))
	for i := range bootstrappers {
		info, err := peer.AddrInfoFromP2pAddr(bootstrappers[i])
		if err != nil {
			continue
		}
		ds = append(ds, *info)
	}

	return &options{
		LowWater:       200,
		HighWater:      400,
		GracePeriod:    time.Second * 20,
		Rendezvous:     "/sonr/rendevouz/0.9.2",
		MultiAddrs:     make([]multiaddr.Multiaddr, 0),
		Interval:       time.Second * 5,
		BootstrapPeers: ds,
		TTL:            dscl.TTL(time.Minute * 2),
	}
}

// Apply applies the host options and returns new SNRHost
func (opts *options) Apply(ctx context.Context, privKey crypto.PrivKey, options ...Option) (*host, error) {
	// Iterate over the options.
	var err error
	for _, opt := range options {
		opt(opts)
	}

	// Create the host.
	hn := &host{
		ctx:          ctx,
		status:       Status_IDLE,
		mdnsPeerChan: make(chan peer.AddrInfo),
	}

	// Start Host
	hn.Host, err = libp2p.New(
		libp2p.Identity(privKey),
		libp2p.ConnectionManager(connmgr.NewConnManager(
			opts.LowWater,    // Lowwater
			opts.HighWater,   // HighWater,
			opts.GracePeriod, // GracePeriod
		)),
		libp2p.DefaultListenAddrs,
		libp2p.Routing(hn.Router),
		libp2p.EnableAutoRelay(),
	)
	if err != nil {
		logger.Errorf("%s - NewHost: Failed to create libp2p host", err)
		return nil, err
	}
	hn.SetStatus(Status_CONNECTING)

	// Bootstrap DHT
	if err := hn.Bootstrap(context.Background()); err != nil {
		logger.Errorf("%s - Failed to Bootstrap KDHT to Host", err)
		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Connect to Bootstrap Nodes
	for _, pi := range opts.BootstrapPeers {
		if err := hn.Connect(pi); err != nil {
			continue
		} else {
			break
		}
	}

	// Initialize Discovery for DHT
	if err := hn.createDHTDiscovery(opts); err != nil {
		logger.Fatal("Could not start DHT Discovery", err)
		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Set the private key.
	return hn, nil
}
