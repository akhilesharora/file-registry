package services

import (
	"context"
	"io"

	"github.com/akhilesharora/file-registry/internal/core/domain"
	"github.com/akhilesharora/file-registry/internal/core/ports"
)

type fileService struct {
	storage   ports.Storage
	transport ports.Transport
}

func NewFileService(storage ports.Storage, transport ports.Transport) ports.FileService {
	return &fileService{
		storage:   storage,
		transport: transport,
	}
}

func (s *fileService) UploadFile(ctx context.Context, path string, content io.Reader) (*domain.File, error) {
	cid, err := s.transport.Upload(ctx, content)
	if err != nil {
		return nil, domain.ErrUploadFailed
	}

	file, err := domain.NewFile(path, cid)
	if err != nil {
		return nil, err
	}

	if err := s.storage.Save(ctx, file.Path, file.CID); err != nil {
		return nil, err
	}

	return file, nil
}

func (s *fileService) GetFile(ctx context.Context, path string) (*domain.File, error) {
	cid, err := s.storage.Get(ctx, path)
	if err != nil {
		return nil, domain.ErrFileNotFound
	}

	file, err := domain.NewFile(path, cid)
	if err != nil {
		return nil, err
	}

	return file, nil
}
