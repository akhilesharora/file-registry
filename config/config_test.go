package config

import (
	"os"
	"testing"
)

func TestLoadWithDefaultValues(t *testing.T) {
	clearEnvVars := []string{"ETH_NODE_URL", "PRIVATE_KEY", "CONTRACT_ADDRESS", "IPFS_NODE_URL", "PORT"}
	originalEnv := make(map[string]string)
	for _, env := range clearEnvVars {
		originalEnv[env] = os.Getenv(env)
		os.Unsetenv(env)
	}

	defer func() {
		// Restore original environment
		for k, v := range originalEnv {
			if v != "" {
				os.Setenv(k, v)
			} else {
				os.Unsetenv(k)
			}
		}
	}()

	cfg := Load()

	if cfg.EthNodeURL != "http://localhost:8545" {
		t.Errorf("expected default ETH_NODE_URL to be http://localhost:8545, got %s", cfg.EthNodeURL)
	}

	if cfg.PrivateKey != "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" {
		t.Errorf("expected default PRIVATE_KEY, got %s", cfg.PrivateKey)
	}

	if cfg.ContractAddress != "" {
		t.Errorf("expected default CONTRACT_ADDRESS to be empty, got %s", cfg.ContractAddress)
	}

	if cfg.IPFSNodeURL != "http://localhost:5001" {
		t.Errorf("expected default IPFS_NODE_URL to be http://localhost:5001, got %s", cfg.IPFSNodeURL)
	}

	if cfg.Port != "8090" {
		t.Errorf("expected default PORT to be 8090, got %s", cfg.Port)
	}
}

func TestLoadWithEnvValues(t *testing.T) {
	os.Setenv("ETH_NODE_URL", "http://custom-node:8545")
	os.Setenv("PRIVATE_KEY", "custom_private_key")
	os.Setenv("CONTRACT_ADDRESS", "0x1234567890abcdef")
	os.Setenv("IPFS_NODE_URL", "http://custom-ipfs:5001")
	os.Setenv("PORT", "9000")

	defer func() {
		// Clean up
		os.Unsetenv("ETH_NODE_URL")
		os.Unsetenv("PRIVATE_KEY")
		os.Unsetenv("CONTRACT_ADDRESS")
		os.Unsetenv("IPFS_NODE_URL")
		os.Unsetenv("PORT")
	}()

	cfg := Load()

	if cfg.EthNodeURL != "http://custom-node:8545" {
		t.Errorf("expected ETH_NODE_URL to be http://custom-node:8545, got %s", cfg.EthNodeURL)
	}

	if cfg.PrivateKey != "custom_private_key" {
		t.Errorf("expected PRIVATE_KEY to be custom_private_key, got %s", cfg.PrivateKey)
	}

	if cfg.ContractAddress != "0x1234567890abcdef" {
		t.Errorf("expected CONTRACT_ADDRESS to be 0x1234567890abcdef, got %s", cfg.ContractAddress)
	}

	if cfg.IPFSNodeURL != "http://custom-ipfs:5001" {
		t.Errorf("expected IPFS_NODE_URL to be http://custom-ipfs:5001, got %s", cfg.IPFSNodeURL)
	}

	if cfg.Port != "9000" {
		t.Errorf("expected PORT to be 9000, got %s", cfg.Port)
	}
}
