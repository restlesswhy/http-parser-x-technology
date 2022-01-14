package parser

import (
	"context"

	"github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser/models"
)
type Repository interface {
	Get(ctx context.Context) ([]models.Data, error)
	Update(ctx context.Context, data models.Data) error
	Create(ctx context.Context, data models.Data) (error)
}