package services

import "github.com/torvald2/solanamon/models"

type CandleGetter interface {
	GetCandles(symbol string, pair string, unit models.DateTrunc) (candles []models.Candle, err error)
}

func GetCandles(symbol string, pair string, unit models.DateTrunc, s CandleGetter) ([]models.Candle, error) {

	return s.GetCandles(symbol, pair, unit)
}
