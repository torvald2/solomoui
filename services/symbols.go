package services

import "github.com/torvald2/solanamon/models"

type SymbolsGetter interface {
	GetAllSymbols() (symbols []models.Symbol, err error)
}

func GetSymbols(s SymbolsGetter) ([]models.Symbol, error) {
	return s.GetAllSymbols()
}
