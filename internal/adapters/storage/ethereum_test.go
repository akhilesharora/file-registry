package storage

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

type MockContract struct {
	SaveFunc func(opts *bind.TransactOpts, filePath string, cid string) (*types.Transaction, error)
	GetFunc  func(opts *bind.CallOpts, filePath string) (string, error)
}

func (m *MockContract) Save(opts *bind.TransactOpts, filePath string, cid string) (*types.Transaction, error) {
	if m.SaveFunc != nil {
		return m.SaveFunc(opts, filePath, cid)
	}
	return types.NewTx(&types.LegacyTx{}), nil
}

func (m *MockContract) Get(opts *bind.CallOpts, filePath string) (string, error) {
	if m.GetFunc != nil {
		return m.GetFunc(opts, filePath)
	}
	return "", nil
}

func TestEthereumStorage(t *testing.T) {
	mockContract := &MockContract{
		SaveFunc: func(opts *bind.TransactOpts, filePath string, cid string) (*types.Transaction, error) {
			return types.NewTx(&types.LegacyTx{}), nil
		},
		GetFunc: func(opts *bind.CallOpts, filePath string) (string, error) {
			return "QmTestCID123", nil
		},
	}

	storage := &EthereumStorage{
		contract: mockContract,
		auth:     &bind.TransactOpts{},
	}

	// Test Save
	err := storage.Save(context.Background(), "test.txt", "QmTestCID123")
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Test Get
	cid, err := storage.Get(context.Background(), "test.txt")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if cid != "QmTestCID123" {
		t.Errorf("Expected CID %s, got %s", "QmTestCID123", cid)
	}
}
