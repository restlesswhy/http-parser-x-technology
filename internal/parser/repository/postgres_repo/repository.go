package postgresrepo

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser/models"
)

const (
	dataTable = "some_data"
)

type ParserRepository struct {
	db *sqlx.DB
}

func NewParserRepository(db *sqlx.DB) *ParserRepository {
	return &ParserRepository{
		db: db,
	}
}

func (r *ParserRepository) Get(ctx context.Context) ([]models.Data, error) {
	var data []models.Data

	query := fmt.Sprintf("SELECT symbol, price, volume, last_trade FROM %s", 
	dataTable)
	if err := r.db.Select(&data, query); err != nil {
		return nil, err
	}

	return data, nil
}

func (r *ParserRepository) Update(ctx context.Context, data models.Data) error {
	return nil
}

func (r *ParserRepository) Create(ctx context.Context, data models.Data) error {
	query := fmt.Sprintf(`INSERT INTO %s (symbol, price, volume, last_trade) 
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (symbol) DO UPDATE 
	  SET price = $2, 
		  volume = $3,
		  last_trade = $4;`, dataTable)
	_, err := r.db.ExecContext(ctx, query, data.Symbol, data.Price, data.Volume, data.LastTrade)
	if err != nil {
		return errors.Wrap(err, "Create.Exec")
	}
	

	return nil
}

	
	