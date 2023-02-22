package repo

import (
	"context"

	"github.com/karalef/shlink/model"
)

// Repository represents a repository.
type Repository interface {
	Store(ctx context.Context, short, origin string) error
	Get(ctx context.Context, short string) (*model.URL, error)
	Update(ctx context.Context, short string, viewsDelta int64) error
	Delete(ctx context.Context, short string) error
}
