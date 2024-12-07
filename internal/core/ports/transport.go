package ports

import (
	"context"
	"io"
)

type Transport interface {
	Upload(ctx context.Context, reader io.Reader) (string, error)
}
