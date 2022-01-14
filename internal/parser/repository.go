package parser

import (
	"context"

	"github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser/models"
)
type Repository interface {
	Get(ctx context.Context) ([]models.Data, error)
	Create(data models.Data) (error)
}