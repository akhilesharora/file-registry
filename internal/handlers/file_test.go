package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/akhilesharora/file-registry/internal/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockFileService struct {
	mock.Mock
}

func (m *mockFileService) UploadFile(ctx context.Context, path string, content io.Reader) (*domain.File, error) {
	args := m.Called(ctx, path, content)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.File), args.Error(1)
}

func (m *mockFileService) GetFile(ctx context.Context, path string) (*domain.File, error) {
	args := m.Called(ctx, path)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.File), args.Error(1)
}

func TestFileHandler_HandleUpload(t *testing.T) {
	mockService := new(mockFileService)
	handler := NewFileHandler(mockService)

	t.Run("successful upload", func(t *testing.T) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", "test.txt")
		assert.NoError(t, err)

		content := "test content"
		_, err = io.Copy(part, strings.NewReader(content))
		assert.NoError(t, err)

		err = writer.WriteField("filePath", "/test.txt")
		assert.NoError(t, err)

		writer.Close()

		req := httptest.NewRequest(http.MethodPost, "/v1/files", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		rr := httptest.NewRecorder()

		expectedFile := &domain.File{
			Path: "/test.txt",
			CID:  "QmTest123",
		}
		mockService.On("UploadFile", mock.Anything, "/test.txt", mock.Anything).Return(expectedFile, nil)

		handler.handleUpload(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response Response
		err = json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.True(t, response.Success)

		data := response.Data.(map[string]interface{})
		assert.Equal(t, expectedFile.Path, data["path"])
		assert.Equal(t, expectedFile.CID, data["cid"])

		mockService.AssertExpectations(t)
	})
}

func TestFileHandler_HandleGet(t *testing.T) {
	mockService := new(mockFileService)
	handler := NewFileHandler(mockService)

	t.Run("successful get", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/files?filePath=/test.txt", nil)
		rr := httptest.NewRecorder()

		expectedFile := &domain.File{
			Path: "/test.txt",
			CID:  "QmTest123",
		}
		mockService.On("GetFile", mock.Anything, "/test.txt").Return(expectedFile, nil)

		handler.handleGet(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)

		var response Response
		err := json.NewDecoder(rr.Body).Decode(&response)
		assert.NoError(t, err)
		assert.True(t, response.Success)

		data := response.Data.(map[string]interface{})
		assert.Equal(t, expectedFile.Path, data["path"])
		assert.Equal(t, expectedFile.CID, data["cid"])

		mockService.AssertExpectations(t)
	})
}
