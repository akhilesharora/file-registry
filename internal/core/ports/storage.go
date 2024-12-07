package ports

import (
	"context"
)

type Storage interface {
	Save(ctx context.Context, path string, cid string) error
	Get(ctx context.Context, path string) (string, error)
}
