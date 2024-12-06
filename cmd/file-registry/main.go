// cmd/file-registry/main.go
package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/akhilesharora/file-registry/config"
	"github.com/akhilesharora/file-registry/internal/adapters/storage"
	"github.com/akhilesharora/file-registry/internal/adapters/transport"
	services "github.com/akhilesharora/file-registry/internal/core/service"
	"github.com/akhilesharora/file-registry/internal/handlers"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	// Load config
	cfg := config.Load()
	log.Printf("Private key loaded: %s", cfg)

	// Setup services
	ethClient, auth, err := storage.SetupEthereum(&cfg)
	if err != nil {
		log.Fatalf("Failed to setup ethereum: %v", err)
	}

	// Initialize components
	ethStorage, err := storage.NewEthereumStorage(
		ethClient,
		common.HexToAddress(cfg.ContractAddress),
		auth,
	)
	if err != nil {
		log.Fatalf("Failed to setup storage: %v", err)
	}

	ipfsTransport, err := transport.NewIPFSTransport(cfg.IPFSNodeURL)
	if err != nil {
		log.Fatalf("Failed to setup IPFS: %v", err)
	}
	fileService := services.NewFileService(ethStorage, ipfsTransport)
	fileHandler := handlers.NewFileHandler(fileService)

	// Setup and start HTTP server
	mux := http.NewServeMux()
	fileHandler.Register(mux)

	server := startServer(mux, cfg.Port)
	waitForShutdown(server)
}

func startServer(handler http.Handler, port string) *http.Server {
	addr := net.JoinHostPort("0.0.0.0", port)

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		log.Printf("Server starting at http://%s", addr)
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server error: %v", err)
		}
	}()

	return srv
}

func waitForShutdown(srv *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
}
