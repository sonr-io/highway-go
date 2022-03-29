package main

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	n "github.com/sonr-io/core/node"
	"github.com/sonr-io/highway-go/config"
	"github.com/sonr-io/highway-go/controller"
	db "github.com/sonr-io/highway-go/database"
	"github.com/sonr-io/highway-go/logger"
	log "github.com/sonr-io/highway-go/logger"
	"github.com/sonr-io/highway-go/models"
	"github.com/sonr-io/highway-go/server"
	"google.golang.org/grpc/credentials"

	"google.golang.org/grpc"
)

const (
	// PEM_CERT_FILE is the path to the certificate file.
	PEM_CERT_FILE = "cert.pem"

	// PEM_KEY_FILE is the file containing the private key.
	PEM_KEY_FILE = "key.pem"
)

// Error Definitions
var (
	//logger                 = golog.Default.Child("grpc/highway")
	ErrEmptyQueue          = errors.New("no items in Transfer Queue")
	ErrInvalidQuery        = errors.New("no SName or PeerID provided")
	ErrMissingParam        = errors.New("paramater is missing")
	ErrProtocolsNotSet     = errors.New("node Protocol has not been initialized")
	ErrMethodUnimplemented = errors.New("method is not implemented")
)

func main() {

	var ctx context.Context

	highwayConfig, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	authConfig := &config.Config{
		HostAddress:  highwayConfig.HighwayAddress,
		HostPort:     highwayConfig.HttpPort,
		DBName:       highwayConfig.SqlName,
		DBPath:       highwayConfig.SqlPath,
		RelyingParty: highwayConfig.RelyingParty,
		RPOrigin:     highwayConfig.RPOrigin + highwayConfig.RPPort,
	}

	err = log.Setup(authConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Create the cosmos listener.
	l, err := net.Listen(verifyAddress(highwayConfig))
	if err != nil {
		log.Fatal(err)
	}

	cosmos, err := n.NewCosmos(context.Background(), l.Addr().String(), "test", "unimplemented-password")
	if err != nil {
		log.Fatal(err)
	}
	// Get TLS config if TLS is enabled
	var stub *models.HighwayStub
	credentials, err := loadTLSCredentials()

	if err != nil {
		//logger.Infof("Error loading TLS credentials: ", err)

		// If TLS is not enabled, create a new listener.
		// Create the RPC Service
		stub = &models.HighwayStub{
			Host:   nil,
			Ctx:    ctx,
			Grpc:   grpc.NewServer(),
			Cosmos: cosmos.Client,
		}
		// hw.RegisterHighwayServer(stub.Grpc, stub)
		//reflection.RegisterReflection(stub.grpc)
		//logger.Infof("Starting RPC Service on %s", l.Addr().String())
	} else {
		// Create the RPC Service
		stub = &models.HighwayStub{
			Host:   nil,
			Ctx:    ctx,
			Grpc:   grpc.NewServer(grpc.Creds(credentials)),
			Cosmos: cosmos.Client,
		}
		// hw.RegisterHighwayServer(stub.Grpc, stub)
		//reflection.RegisterReflection(stub.grpc)
		//logger.Infof("Starting RPC Service on %s", l.Addr().String())
	}

	DB, err := db.Connect(highwayConfig.MongoUri, highwayConfig.MongoCollectionName, highwayConfig.MongoDbName)
	if err != nil {
		logger.Errorf("database connection failed")
	}

	ctrl, err := controller.New(DB, highwayConfig, stub)
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.NewServer(ctrl, authConfig)
	if err != nil {
		log.Fatal(err)
	}
	go server.Start()

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGINT)

	<-c
	log.Info("Shutting down...")
	server.Shutdown()
}

// verifyAddress verifies the address is valid.
func verifyAddress(cnfg *config.SonrConfig) (string, string) {
	// Define Variables
	var network string
	var address string
	var port int
	var err error

	// Set Network
	if cnfg.HighwayNetwork != "tcp" {
		network = "tcp"
	}
	logger.Infof("Network: %s", network)

	//set port
	port, err = strconv.Atoi(cnfg.GrpcPort)
	if err != nil {
		return "", err.Error()
	}
	logger.Infof("Using port %d", port)

	// Set Address
	if !strings.Contains(cnfg.HighwayAddress, ":") {
		address = fmt.Sprintf("%s:%d", cnfg.HighwayAddress, port)
	}
	logger.Infof("Address: %s", address)

	return network, address
}

// Helper function to see if file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair(PEM_CERT_FILE, PEM_KEY_FILE)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
