package usecase

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/restlesswhy/rest/http-parsing-x-technology/config"
	"github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser"
	"github.com/restlesswhy/rest/http-parsing-x-technology/internal/parser/models"
	"github.com/restlesswhy/rest/http-parsing-x-technology/pkg/logger"
)

type ParserUseCase struct {
	parserRepo parser.Repository
	cfg          *config.Config
}

func NewParserUseCase(cfg *config.Config, parserRepo parser.Repository) *ParserUseCase {
	return &ParserUseCase{
		cfg: cfg,
		parserRepo: parserRepo,
	}
}


func (u *ParserUseCase) Get(ctx context.Context) ([]map[string]map[string]float64, error) {
	err := u.jsonParser(ctx)
	if err != nil {
		return nil, err
	}

	data, err := u.parserRepo.Get(ctx)
	if err != nil {
		return nil, err
	}

	mapData := convertData(data)
	return mapData, nil
}

func (u *ParserUseCase) jsonParser(ctx context.Context) error {
	allData := getJsonData()
	for _, data := range allData {
		err := u.parserRepo.Create(ctx, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func getJsonData() []models.Data {
	url := "https://api.blockchain.com/v3/exchange/tickers"

	var netClient = http.Client{
		Timeout: time.Second * 10,
	}

	res, err := netClient.Get(url)
	if err != nil {
		logger.Fatal(err)
	}

	defer res.Body.Close()

	var jsonParseData []models.Data
	err = json.NewDecoder(res.Body).Decode(&jsonParseData)
	if err != nil {
		logger.Fatal(err)
	}

	return jsonParseData
}

func convertData(jsonParseData []models.Data) []map[string]map[string]float64 {
	var allData []map[string]map[string]float64
	for i := 0; i < len(jsonParseData); i++ {
		var oneData map[string]map[string]float64 = make(map[string]map[string]float64)
		// oneData = make(map[string]map[string]float32, 1)
		oneData[jsonParseData[i].Symbol] = make(map[string]float64)
		oneData[jsonParseData[i].Symbol]["price"] = jsonParseData[i].Price
		oneData[jsonParseData[i].Symbol]["volume"] = jsonParseData[i].Volume
		oneData[jsonParseData[i].Symbol]["last_trade"] = jsonParseData[i].LastTrade

		allData = append(allData, oneData)
	}
	return allData
}