package transport

import (
	"context"
	"fmt"
	"io"
	"log"

	shell "github.com/ipfs/go-ipfs-api"
)

type ipfsShell interface {
	Add(r io.Reader, options ...shell.AddOpts) (string, error)
}

type IPFSTransport struct {
	sh ipfsShell
}

func NewIPFSTransport(nodeURL string) (*IPFSTransport, error) {
	ipfsShell := shell.NewShell(nodeURL)
	// Verify connection
	version, _, err := ipfsShell.Version()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to IPFS node: %w", err)
	}
	log.Printf("Connected to IPFS node. Version: %s", version)

	// Create IPFS shell
	return &IPFSTransport{
		sh: shell.NewShell(nodeURL),
	}, nil
}

func (t *IPFSTransport) Upload(ctx context.Context, reader io.Reader) (string, error) {
	log.Printf("Attempting to upload file to IPFS")
	// IPFS doesn't natively support context cancellation
	cid, err := t.sh.Add(reader)
	if err != nil {
		log.Printf("IPFS upload failed: %v", err)
		return "", fmt.Errorf("failed to upload to IPFS: %w", err)
	}

	log.Printf("Successfully uploaded to IPFS. CID: %s", cid)
	return cid, nil
}
