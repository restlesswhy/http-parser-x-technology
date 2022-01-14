package usecase

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
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

type single struct {
}

var singleInstance *single

func (u *ParserUseCase) StartParse() {
	if singleInstance == nil {
		u.CallAt(u.jsonParser)
		singleInstance = &single{}
    }
}

func (u *ParserUseCase) Get(ctx context.Context) ([]map[string]map[string]float64, error) {
	data, err := u.parserRepo.Get(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Get.Get")
	}

	mapData := convertData(data)
	return mapData, nil
}

//jsonParser запускает парсинг и каждый полученный элемент добавляет в базу
func (u *ParserUseCase) jsonParser() error {
	allData := getJsonData()
	for _, data := range allData {
		err := u.parserRepo.Create(data)
		if err != nil {
			return errors.Wrap(err, "jsonParser.Create")
		}
	}

	return nil
}

//getJsonData парсит json в слайс струкутур Data.
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

//convertData конвертирует структуру в необходимую для json мапу.
func convertData(jsonParseData []models.Data) []map[string]map[string]float64 {
	var allData []map[string]map[string]float64
	for i := 0; i < len(jsonParseData); i++ {
		var oneData map[string]map[string]float64 = make(map[string]map[string]float64)
		oneData[jsonParseData[i].Symbol] = make(map[string]float64)
		oneData[jsonParseData[i].Symbol]["price"] = jsonParseData[i].Price
		oneData[jsonParseData[i].Symbol]["volume"] = jsonParseData[i].Volume
		oneData[jsonParseData[i].Symbol]["last_trade"] = jsonParseData[i].LastTrade

		allData = append(allData, oneData)
	}
	return allData
}

//CallAt запускает отдельную горутину которая в бесконечном цикле через определенное время парсит json и обновляет базу
func (u *ParserUseCase) CallAt(f func() error ) error {
	go func() {
		for {
			f()
			// Следующий запуск через n времени
			logger.Info("DB is update!")
			time.Sleep(u.cfg.Parser.IterationTime)
		}
	}()

	return nil
}