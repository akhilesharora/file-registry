package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/akhilesharora/file-registry/config"
	"github.com/akhilesharora/file-registry/contracts/fileregistry"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	simulated "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Contract interface {
	Save(opts *bind.TransactOpts, filePath string, cid string) (*types.Transaction, error)
	Get(opts *bind.CallOpts, filePath string) (string, error)
}

type EthereumStorage struct {
	backend  bind.ContractBackend
	contract Contract
	auth     *bind.TransactOpts
}

func NewEthereumStorage(backend bind.ContractBackend, contractAddr common.Address, auth *bind.TransactOpts) (*EthereumStorage, error) {
	log.Printf("Initializing Ethereum storage with contract address: %s", contractAddr.Hex())

	fileRegistry, err := fileregistry.NewFileRegistry(contractAddr, backend)
	if err != nil {
		log.Printf("Failed to instantiate contract: %v", err)
		return nil, fmt.Errorf("failed to instantiate contract: %w", err)
	}

	return &EthereumStorage{
		backend:  backend,
		auth:     auth,
		contract: fileRegistry,
	}, nil
}

func (s *EthereumStorage) Save(ctx context.Context, path string, cid string) error {
	log.Printf("Attempting to save to contract - Path: %s, CID: %s", path, cid)

	if path == "" || cid == "" {
		return fmt.Errorf("path and CID cannot be empty")
	}

	s.auth.Context = ctx
	tx, err := s.contract.Save(s.auth, path, cid)
	if err != nil {
		log.Printf("Failed to save to contract: %v", err)
		return fmt.Errorf("failed to save to contract: %w", err)
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())

	if simBackend, ok := s.backend.(*simulated.SimulatedBackend); ok {
		log.Printf("Using simulated backend, committing transaction")
		simBackend.Commit()
		return nil
	}

	// For real backend, wait for mining
	deployBackend, ok := s.backend.(bind.DeployBackend)
	if !ok || deployBackend == nil {
		log.Printf("No deploy backend present, skipping WaitMined")
		return nil
	}

	receipt, err := bind.WaitMined(ctx, deployBackend, tx)
	if err != nil {
		log.Printf("Failed to wait for transaction to be mined: %v", err)
		return fmt.Errorf("transaction failed to be mined: %w", err)
	}

	if receipt.Status == 0 {
		log.Printf("Transaction failed: %s", tx.Hash().Hex())
		return fmt.Errorf("transaction failed")
	}

	log.Printf("Successfully saved to contract. Transaction hash: %s", tx.Hash().Hex())
	return nil
}

func (s *EthereumStorage) Get(ctx context.Context, path string) (string, error) {
	log.Printf("Attempting to get from contract - Path: %s", path)

	if path == "" {
		return "", fmt.Errorf("path cannot be empty")
	}

	cid, err := s.contract.Get(&bind.CallOpts{Context: ctx}, path)
	if err != nil {
		log.Printf("Failed to get from contract: %v", err)
		return "", fmt.Errorf("failed to get from contract: %w", err)
	}

	if cid == "" {
		log.Printf("No CID found for path: %s", path)
		return "", fmt.Errorf("file not found: %s", path)
	}

	log.Printf("Successfully retrieved CID: %s for path: %s", cid, path)
	return cid, nil
}

func SetupEthereum(cfg *config.Config) (*ethclient.Client, *bind.TransactOpts, error) {
	client, err := ethclient.Dial(cfg.EthNodeURL)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}
	// Verify connection by getting the latest block number
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to retrieve block number: %w", err)
	}

	log.Printf("Connected to Ethereum node. Latest block number: %d", blockNumber)

	// Get the private key
	privateKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("invalid private key: %w", err)
	}

	// Get the chain ID
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Create the transactor
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		client.Close()
		return nil, nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	return client, auth, nil
}
