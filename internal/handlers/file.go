package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/akhilesharora/file-registry/internal/core/ports"
)

type FileHandler struct {
	service ports.FileService
}

func NewFileHandler(service ports.FileService) *FileHandler {
	return &FileHandler{service: service}
}

func (h *FileHandler) Register(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/", fs)

	mux.HandleFunc("/v1/files", h.enableCORS(h.handleFiles))
}

func (h *FileHandler) enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func (h *FileHandler) handleFiles(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.handleUpload(w, r)
	case http.MethodGet:
		h.handleGet(w, r)
	default:
		log.Printf("Method not allowed: %s", r.Method)
		writeError(w, http.StatusMethodNotAllowed, "Method not allowed")
	}
}

func (h *FileHandler) handleUpload(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling file upload request")

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Printf("Failed to parse form: %v", err)
		writeError(w, http.StatusBadRequest, "Failed to parse form")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Printf("No file provided: %v", err)
		writeError(w, http.StatusBadRequest, "No file provided")
		return
	}
	defer file.Close()

	filePath := r.FormValue("filePath")
	if filePath == "" {
		log.Printf("No filePath provided")
		writeError(w, http.StatusBadRequest, "filePath is required")
		return
	}

	log.Printf("Processing file upload - Name: %s, Size: %d bytes, Path: %s",
		header.Filename, header.Size, filePath)

	result, err := h.service.UploadFile(r.Context(), filePath, file)
	if err != nil {
		log.Printf("Upload failed: %v", err)
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Upload failed: %v", err))
		return
	}

	log.Printf("File uploaded successfully - Path: %s, CID: %s", result.Path, result.CID)
	writeJSON(w, http.StatusOK, Response{
		Success: true,
		Data: map[string]interface{}{
			"path": result.Path,
			"cid":  result.CID,
		},
		Message: fmt.Sprintf("File uploaded successfully. CID: %s", result.CID),
	})
}

func (h *FileHandler) handleGet(w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Query().Get("filePath")
	log.Printf("Handling get request for path: %s", filePath)

	if filePath == "" {
		log.Printf("No filePath provided in query")
		writeError(w, http.StatusBadRequest, "filePath is required")
		return
	}

	file, err := h.service.GetFile(r.Context(), filePath)
	if err != nil {
		log.Printf("Failed to get file: %v", err)
		writeError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get file: %v", err))
		return
	}

	log.Printf("File found - Path: %s, CID: %s", file.Path, file.CID)
	writeJSON(w, http.StatusOK, Response{
		Success: true,
		Data: map[string]interface{}{
			"path": file.Path,
			"cid":  file.CID,
		},
		Message: fmt.Sprintf("File found with CID: %s", file.CID),
	})
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error writing JSON response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	log.Printf("Error response: %s (Status: %d)", message, status)
	writeJSON(w, status, Response{
		Success: false,
		Error:   message,
	})
}
