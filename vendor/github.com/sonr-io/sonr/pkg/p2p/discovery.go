package p2p

import (
	dsc "github.com/libp2p/go-libp2p-discovery"
	psub "github.com/libp2p/go-libp2p-pubsub"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	rt "github.com/sonr-io/sonr/x/registry/types"
)

// createMdnsDiscovery is a Helper Method to initialize the MDNS Discovery
func (hn *host) createMdnsDiscovery(opts *options) {
	// Verify if MDNS is Enabled
	if hn.connection == rt.Connection_CONNECTION_OFFLINE {
		logger.Errorf("%s - Failed to Start MDNS Discovery ", ErrMDNSInvalidConn)
		return
	}

	// Create MDNS Service
	ser := mdns.NewMdnsService(hn.Host, opts.Rendezvous, hn)
	err := ser.Start()
	if err != nil {
		logger.Errorf("%s - Failed to Start MDNS Discovery ", err)
		return
	}
}

// createDHTDiscovery is a Helper Method to initialize the DHT Discovery
func (hn *host) createDHTDiscovery(opts *options) error {
	// Set Routing Discovery, Find Peers
	var err error
	routingDiscovery := dsc.NewRoutingDiscovery(hn.IpfsDHT)
	dsc.Advertise(hn.ctx, routingDiscovery, opts.Rendezvous, opts.TTL)

	// Create Pub Sub
	hn.PubSub, err = psub.NewGossipSub(hn.ctx, hn.Host, psub.WithDiscovery(routingDiscovery))
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Errorf("%s - Failed to Create new Gossip Sub", err)
		return err
	}

	// Handle DHT Peers
	hn.dhtPeerChan, err = routingDiscovery.FindPeers(hn.ctx, opts.Rendezvous, opts.TTL)
	if err != nil {
		hn.SetStatus(Status_FAIL)
		logger.Errorf("%s - Failed to create FindPeers Discovery channel", err)
		return err
	}
	hn.SetStatus(Status_READY)
	return nil
}
