package services

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// mockStorage conforms to ports.Storage
type mockStorage struct {
	mock.Mock
}

func (m *mockStorage) Save(ctx context.Context, path string, cid string) error {
	args := m.Called(ctx, path, cid)
	return args.Error(0)
}

func (m *mockStorage) Get(ctx context.Context, path string) (string, error) {
	args := m.Called(ctx, path)
	return args.String(0), args.Error(1)
}

// mockTransport conforms to ports.Transport
type mockTransport struct {
	mock.Mock
}

func (m *mockTransport) Upload(ctx context.Context, reader io.Reader) (string, error) {
	args := m.Called(ctx, reader)
	return args.String(0), args.Error(1)
}

func TestFileService_UploadFile(t *testing.T) {
	mockStorage := new(mockStorage)
	mockTransport := new(mockTransport)
	service := NewFileService(mockStorage, mockTransport)

	ctx := context.Background()
	content := strings.NewReader("test content")
	path := "/test.txt"
	expectedCID := "QmTest123"

	mockTransport.On("Upload", ctx, content).Return(expectedCID, nil)
	mockStorage.On("Save", ctx, path, expectedCID).Return(nil)

	file, err := service.UploadFile(ctx, path, content)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, path, file.Path)
	assert.Equal(t, expectedCID, file.CID)

	mockStorage.AssertExpectations(t)
	mockTransport.AssertExpectations(t)
}

func TestFileService_GetFile(t *testing.T) {
	mockStorage := new(mockStorage)
	mockTransport := new(mockTransport) // not needed directly here, but included for completeness
	service := NewFileService(mockStorage, mockTransport)

	ctx := context.Background()
	path := "/test.txt"
	expectedCID := "QmTest123"

	mockStorage.On("Get", ctx, path).Return(expectedCID, nil)

	file, err := service.GetFile(ctx, path)
	assert.NoError(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, path, file.Path)
	assert.Equal(t, expectedCID, file.CID)

	mockStorage.AssertExpectations(t)
}
