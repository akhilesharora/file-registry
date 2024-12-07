package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EthNodeURL      string
	PrivateKey      string
	ContractAddress string
	IPFSNodeURL     string
	Port            string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or error loading: %v", err)
	}

	config := Config{
		EthNodeURL:      getEnvOrDefault("ETH_NODE_URL", "http://localhost:8545"),
		PrivateKey:      getEnvOrDefault("PRIVATE_KEY", "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"),
		ContractAddress: getEnvOrDefault("CONTRACT_ADDRESS", ""),
		IPFSNodeURL:     getEnvOrDefault("IPFS_NODE_URL", "http://localhost:5001"),
		Port:            getEnvOrDefault("PORT", "8090"),
	}

	// Log loaded configuration (except private key)
	log.Printf("Loaded configuration: ETH_NODE_URL=%s, IPFS_NODE_URL=%s, CONTRACT_ADDRESS=%s, PORT=%s",
		config.EthNodeURL, config.IPFSNodeURL, config.ContractAddress, config.Port)

	return config
}

func (c Config) Validate() error {
	if c.PrivateKey == "" {
		return errors.New("PRIVATE_KEY environment variable is required")
	}
	if c.ContractAddress == "" {
		return errors.New("CONTRACT_ADDRESS environment variable is required")
	}
	return nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}
	log.Printf("Using default value for %s: %s", key, defaultValue)
	return defaultValue
}
