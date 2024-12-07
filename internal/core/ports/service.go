package ports

import (
	"context"
	"io"

	"github.com/akhilesharora/file-registry/internal/core/domain"
)

type FileService interface {
	UploadFile(ctx context.Context, path string, content io.Reader) (*domain.File, error)
	GetFile(ctx context.Context, path string) (*domain.File, error)
}
