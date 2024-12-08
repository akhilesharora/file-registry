package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"testing"

	"github.com/akhilesharora/file-registry/config"
)

type Response struct {
	Success bool `json:"success"`
	Data    struct {
		Path string `json:"path"`
		CID  string `json:"cid"`
	} `json:"data"`
	Error string `json:"error,omitempty"`
}

func TestAPIEndpoints(t *testing.T) {
	cfg := config.Load()
	apiURL := "http://" + net.JoinHostPort("localhost", cfg.Port)

	content := []byte("test content")
	if err := os.WriteFile("test.txt", content, 0644); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test.txt")

	// Test Upload
	file, err := os.Open("test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := io.Copy(part, file); err != nil {
		t.Fatal(err)
	}
	if err := writer.WriteField("filePath", "/test.txt"); err != nil {
		t.Fatal(err)
	}
	writer.Close()

	resp, err := http.Post(apiURL+"/v1/files", writer.FormDataContentType(), body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var uploadResp Response
	if err := json.NewDecoder(resp.Body).Decode(&uploadResp); err != nil {
		t.Fatal(err)
	}

	if !uploadResp.Success {
		t.Fatalf("Upload failed: %s", uploadResp.Error)
	}

	// Test Get
	resp, err = http.Get(apiURL + "/v1/files?filePath=/test.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	var getResp Response
	if err := json.NewDecoder(resp.Body).Decode(&getResp); err != nil {
		t.Fatal(err)
	}

	if !getResp.Success {
		t.Fatalf("Get failed: %s", getResp.Error)
	}

	if getResp.Data.CID != uploadResp.Data.CID {
		t.Errorf("CID mismatch: got %s, want %s", getResp.Data.CID, uploadResp.Data.CID)
	}
}
