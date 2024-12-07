package transport

import (
	"context"
	"io"
	"strings"
	"testing"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockShell struct {
	mock.Mock
}

func (m *mockShell) Add(r io.Reader, options ...shell.AddOpts) (string, error) {
	args := m.Called(r)
	return args.String(0), args.Error(1)
}

func TestIPFSTransport_Upload(t *testing.T) {
	t.Run("successful upload", func(t *testing.T) {
		mockSh := new(mockShell)
		transport := &IPFSTransport{sh: mockSh}

		content := "test content"
		reader := strings.NewReader(content)
		expectedCID := "QmTest123"

		mockSh.On("Add", reader).Return(expectedCID, nil)

		cid, err := transport.Upload(context.Background(), reader)
		assert.NoError(t, err)
		assert.Equal(t, expectedCID, cid)

		mockSh.AssertExpectations(t)
	})

	t.Run("failed upload", func(t *testing.T) {
		mockSh := new(mockShell)
		transport := &IPFSTransport{sh: mockSh}

		content := "test content"
		reader := strings.NewReader(content)

		mockSh.On("Add", reader).Return("", assert.AnError)

		cid, err := transport.Upload(context.Background(), reader)
		assert.Error(t, err)
		assert.Empty(t, cid)

		mockSh.AssertExpectations(t)
	})
}
