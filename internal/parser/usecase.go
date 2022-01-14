package parser

import "context"

type UseCase interface {
	Get(ctx context.Context) ([]map[string]map[string]float64, error)
}