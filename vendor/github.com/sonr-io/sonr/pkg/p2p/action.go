package p2p

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/crypto"
	libp2pHost "github.com/libp2p/go-libp2p-core/host"

	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	"github.com/libp2p/go-libp2p-core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	ps "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-msgio"

	"google.golang.org/protobuf/proto"
)

// Connect connects with `peer.AddrInfo` if underlying Host is ready
func (hn *host) Connect(pi peer.AddrInfo) error {
	// Check if host is ready
	if err := hn.HasRouting(); err != nil {
		logger.Warn("Connect: Underlying host is not ready, failed to call Connect()")
		return err
	}

	// Call Underlying Host to Connect
	return hn.Host.Connect(hn.ctx, pi)
}

// HandlePeerFound is to be called when new  peer is found
func (hn *host) HandlePeerFound(pi peer.AddrInfo) {
	hn.mdnsPeerChan <- pi
}

// HasRouting returns no-error if the host is ready for connect
func (h *host) HasRouting() error {
	if h.IpfsDHT == nil || h.HostImpl == nil {
		return ErrRoutingNotSet
	}
	return nil
}

// Join wraps around PubSub.Join and returns topic. Checks wether the host is ready before joining.
func (hn *host) Join(topic string, opts ...ps.TopicOpt) (*ps.Topic, error) {
	// Check if PubSub is Set
	if hn.PubSub == nil {
		return nil, errors.New("Join: Pubsub has not been set on SNRHost")
	}

	// Check if topic is valid
	if topic == "" {
		return nil, errors.New("Join: Empty topic string provided to Join for host.Pubsub")
	}

	// Call Underlying Pubsub to Connect
	return hn.PubSub.Join(topic, opts...)
}

// NewStream opens a new stream to the peer with given peer id
func (n *host) NewStream(ctx context.Context, pid peer.ID, pids ...protocol.ID) (network.Stream, error) {
	return n.HostImpl.NewStream(ctx, pid, pids...)
}

// NewTopic creates a new topic
func (n *host) NewTopic(name string, opts ...ps.TopicOpt) (*ps.Topic, *ps.TopicEventHandler, *ps.Subscription, error) {
	// Check if PubSub is Set
	if n.PubSub == nil {
		return nil, nil, nil, errors.New("NewTopic: Pubsub has not been set on SNRHost")
	}

	// Call Underlying Pubsub to Connect
	t, err := n.Join(name, opts...)
	if err != nil {
		logger.Errorf("%s - NewTopic: Failed to create new topic", err)
		return nil, nil, nil, err
	}

	// Create Event Handler
	h, err := t.EventHandler()
	if err != nil {
		logger.Errorf("%s - NewTopic: Failed to create new topic event handler", err)
		return nil, nil, nil, err
	}

	// Create Subscriber
	s, err := t.Subscribe()
	if err != nil {
		logger.Errorf("%s - NewTopic: Failed to create new topic subscriber", err)
		return nil, nil, nil, err
	}
	return t, h, s, nil
}

// Router returns the host node Peer Routing Function
func (hn *host) Router(h libp2pHost.Host) (routing.PeerRouting, error) {
	// Create DHT
	kdht, err := dht.New(hn.ctx, h)
	if err != nil {
		hn.SetStatus(Status_FAIL)
		return nil, err
	}

	// Set Properties
	hn.IpfsDHT = kdht
	logger.Debug("Router: Host and DHT have been set for SNRNode")

	// Setup Properties
	return hn.IpfsDHT, nil
}

// SetStreamHandler sets the handler for a given protocol
func (n *host) SetStreamHandler(protocol protocol.ID, handler network.StreamHandler) {
	n.HostImpl.SetStreamHandler(protocol, handler)
}

// SendMessage writes a protobuf go data object to a network stream
func (h *host) SendMessage(id peer.ID, p protocol.ID, data proto.Message) error {
	err := h.HasRouting()
	if err != nil {
		return err
	}

	s, err := h.NewStream(h.ctx, id, p)
	if err != nil {
		logger.Errorf("%s - SendMessage: Failed to start stream", err)
		return err
	}
	defer s.Close()

	// marshall data to protobufs3 binary format
	bin, err := proto.Marshal(data)
	if err != nil {
		logger.Errorf("%s - SendMessage: Failed to marshal pb", err)
		return err
	}

	// Create Writer and write data to stream
	w := msgio.NewWriter(s)
	if err := w.WriteMsg(bin); err != nil {
		logger.Errorf("%s - SendMessage: Failed to write message to stream.", err)
		return err
	}
	return nil
}

// SignData signs an outgoing p2p message payload
func (n *host) SignData(data []byte) ([]byte, error) {
	// Get local node's private key
	//res, err := wallet.Sign(data)
	// if err != nil {
	// 	logger.Errorf("%s - SignData: Failed to get local host's private key", err)
	// 	return nil, err
	// }
	return nil, nil
}

// SignMessage signs an outgoing p2p message payload
func (n *host) SignMessage(message proto.Message) ([]byte, error) {
	data, err := proto.Marshal(message)
	if err != nil {
		logger.Errorf("%s - SignMessage: Failed to Sign Message", err)
		return nil, err
	}
	return n.SignData(data)
}

// // SignedMetadataToProto converts a SignedMetadata to a protobuf.
// func SignedMetadataToProto(m *wallet.SignedMetadata) *common.Metadata {
// 	return &common.Metadata{
// 		Timestamp: m.Timestamp,
// 		NodeId:    m.NodeId,
// 		PublicKey: m.PublicKey,
// 	}
// }

// Stat returns the host stat info
func (hn *host) Stat() (map[string]string, error) {
	// Return Host Stat
	return map[string]string{
		"ID":        hn.ID().String(),
		"Status":    hn.status.String(),
		"MultiAddr": hn.Addrs()[0].String(),
	}, nil
}

// Serve handles incoming peer Addr Info
func (hn *host) Serve() {
	for {
		select {
		case mdnsPI := <-hn.mdnsPeerChan:
			if err := hn.Connect(mdnsPI); err != nil {
				hn.Peerstore().ClearAddrs(mdnsPI.ID)
				continue
			}

		case dhtPI := <-hn.dhtPeerChan:
			if err := hn.Connect(dhtPI); err != nil {
				hn.Peerstore().ClearAddrs(dhtPI.ID)
				continue
			}
		case <-hn.ctx.Done():
			return
		}
	}
}

// VerifyData verifies incoming p2p message data integrity
func (n *host) VerifyData(data []byte, signature []byte, peerId peer.ID, pubKeyData []byte) bool {
	key, err := crypto.UnmarshalPublicKey(pubKeyData)
	if err != nil {
		logger.Errorf("%s - Failed to extract key from message key data", err)
		return false
	}

	// extract node id from the provided public key
	idFromKey, err := peer.IDFromPublicKey(key)
	if err != nil {
		logger.Errorf("%s - VerifyData: Failed to extract peer id from public key", err)
		return false
	}

	// verify that message author node id matches the provided node public key
	if idFromKey != peerId {
		logger.Errorf("%s - VerifyData: Node id and provided public key mismatch", err)
		return false
	}

	res, err := key.Verify(data, signature)
	if err != nil {
		logger.Errorf("%s - VerifyData: Error authenticating data", err)
		return false
	}
	return res
}
