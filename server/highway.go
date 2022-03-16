package highway

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kataras/golog"
	"github.com/sonr-io/highway-go/config"
	controller "github.com/sonr-io/highway-go/controllers"
	db "github.com/sonr-io/highway-go/database"
	"github.com/sonr-io/highway-go/models"
	"github.com/sonr-io/highway-go/pkg/client"
	service "github.com/sonr-io/highway-go/services"
	webAuth "github.com/sonr-io/highway-go/webAuthN"
	"google.golang.org/grpc/credentials"

	hw "go.buf.build/grpc/go/sonr-io/highway/v1"

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
	logger                 = golog.Default.Child("grpc/highway")
	ErrEmptyQueue          = errors.New("no items in Transfer Queue")
	ErrInvalidQuery        = errors.New("no SName or PeerID provided")
	ErrMissingParam        = errors.New("paramater is missing")
	ErrProtocolsNotSet     = errors.New("node Protocol has not been initialized")
	ErrMethodUnimplemented = errors.New("method is not implemented")
)

// Start starts the RPC Service.
func Start(ctx context.Context, cnfg *config.SonrConfig) error {
	r := mux.NewRouter()

	// DB setup
	DB, err := db.Connect(cnfg.MongoUri, cnfg.MongoCollectionName, cnfg.MongoDbName)
	if err != nil {
		logger.Errorf("dtabase connection failed")
	}

	//web authn setup
	auth, err := webAuth.New(cnfg, DB)
	if err != nil {
		return err
	}

	// Create the cosmos listener.
	l, err := net.Listen(verifyAddress(cnfg))
	if err != nil {
		return err
	}

	logger.Infof("Network: " + l.Addr().Network())
	logger.Infof("Address: " + l.Addr().String())

	cosmos, err := client.NewClient(context.Background(), l.Addr().String(), "test", "unimplemented-password")
	if err != nil {
		return err
	}

	// Get TLS config if TLS is enabled
	var stub *models.HighwayStub
	credentials, err := loadTLSCredentials()
	if err != nil {
		logger.Infof("Error loading TLS credentials: ", err)

		// If TLS is not enabled, create a new listener.
		// Create the RPC Service
		stub = &models.HighwayStub{
			Host:     nil,
			Ctx:      ctx,
			Grpc:     grpc.NewServer(),
			Cosmos:   cosmos.Client,
			Listener: l,
		}
		hw.RegisterHighwayServer(stub.Grpc, stub)
		//reflection.RegisterReflection(stub.grpc)
		logger.Infof("Starting RPC Service on %s", l.Addr().String())
	} else {
		// Create the RPC Service
		stub = &models.HighwayStub{
			Host:     nil,
			Ctx:      ctx,
			Grpc:     grpc.NewServer(grpc.Creds(credentials)),
			Cosmos:   cosmos.Client,
			Listener: l,
		}
		hw.RegisterHighwayServer(stub.Grpc, stub)
		//reflection.RegisterReflection(stub.grpc)
		logger.Infof("Starting RPC Service on %s", l.Addr().String())
	}

	httpCtrl, err := controller.New(*DB, cnfg, stub, auth)
	if err != nil {
		return nil
	}
	service.AddHandlers(r, httpCtrl)
	httpAddr := cnfg.HttpPort

	// Check if files exists then start http listener
	if fileExists(PEM_CERT_FILE) && fileExists(PEM_KEY_FILE) {
		logger.Infof("Using TLS")
		return http.ListenAndServeTLS(":"+httpAddr, PEM_CERT_FILE, PEM_KEY_FILE, r)
	} else {
		logger.Warn("Using insecure HTTP")
		return http.ListenAndServe(":"+httpAddr, r)
	}
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
