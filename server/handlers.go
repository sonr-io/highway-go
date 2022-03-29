package server

import (
	"context"
	"net"
	"net/http"
	"time"

	"log"

	"github.com/duo-labs/webauthn/webauthn"
	"github.com/gorilla/mux"
	"github.com/sonr-io/highway-go/config"
	"github.com/sonr-io/highway-go/controller"
	"github.com/sonr-io/highway-go/session"
)

// Timeout is the number of seconds to attempt a graceful shutdown, or
// for timing out read/write operations
const Timeout = 5 * time.Second

// Option is an option that sets a particular value for the server
type Option func(*Server)

// Server is a configurable HTTP server that implements a demo of the WebAuthn
// specification
type Server struct {
	server   *http.Server
	config   *config.Config
	webauthn *webauthn.WebAuthn
	store    *session.Store
	Ctrl     *controller.Controller
}

// NewServer returns a new instance of a Server configured with the provided
// configuration
func NewServer(ctrl *controller.Controller, config *config.Config, opts ...Option) (*Server, error) {
	addr := net.JoinHostPort(config.HostAddress, config.HostPort)
	defaultServer := &http.Server{
		Addr:         addr,
		ReadTimeout:  Timeout,
		WriteTimeout: Timeout,
	}
	defaultStore, err := session.NewStore()
	if err != nil {
		return nil, err
	}
	defaultWebAuthn, _ := webauthn.New(&webauthn.Config{
		RPDisplayName: config.RelyingParty,
		RPID:          config.RelyingParty,
		RPOrigin:      config.RPOrigin,
	})
	ws := &Server{
		config:   config,
		server:   defaultServer,
		store:    defaultStore,
		webauthn: defaultWebAuthn,
		Ctrl:     ctrl,
	}
	for _, opt := range opts {
		opt(ws)
	}
	ws.registerRoutes()
	return ws, nil
}

// WithWebAuthn sets the webauthn configuration for the server
func WithWebAuthn(w *webauthn.WebAuthn) Option {
	return func(ws *Server) {
		ws.webauthn = w
	}
}

// Start starts the underlying HTTP server
func (ws *Server) Start() error {
	log.Printf("Starting webauthn server at %s", ws.server.Addr)
	return ws.server.ListenAndServe()
}

// Shutdown attempts to gracefully shutdown the underlying HTTP server.
func (ws *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	return ws.server.Shutdown(ctx)
}

func (ws *Server) registerRoutes() {
	router := mux.NewRouter()
	// Unauthenticated handlers for registering a new credential and logging in.
	router.HandleFunc("/", ws.Login)
	router.HandleFunc("/makeCredential/{name}", ws.RequestNewCredential).Methods("GET")
	router.HandleFunc("/makeCredential", ws.MakeNewCredential).Methods("POST")
	router.HandleFunc("/assertion/{name}", ws.GetAssertion).Methods("GET")
	router.HandleFunc("/assertion", ws.MakeAssertion).Methods("POST")
	router.HandleFunc("/user/{name}/exists", ws.UserExists).Methods("GET")
	router.HandleFunc("/user/{name}/credentials", ws.GetCredentials).Methods("GET")

	//helper handlers
	router.HandleFunc("/check/name/{name}", ws.CheckName).Methods("GET")
	router.HandleFunc("/health", ws.HealthHandler).Methods("GET")

	// Authenticated handlers for viewing credentials after logging in
	router.HandleFunc("/dashboard", ws.LoginRequired(ws.Index))
	//router.HandleFunc("/register/name/{name}", ws.RegisterName).Methods("POST")

	//stripe
	router.HandleFunc("/create/payment/intent/{name}", ws.CreatePaymentIntent).Methods("POST")
	router.HandleFunc("/stripe/webhook", ws.StripeWebhook).Methods("POST")

	//pages
	router.HandleFunc("/checkout", ws.CheckoutPage)
	router.HandleFunc("/payment", ws.PaymentPage)
	router.HandleFunc("/roadmap", ws.RoadMapPage)

	// Static file serving
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	ws.server.Handler = router
}
