package adaptors

import (
	"testing"

	"github.com/torvald2/solanamon/models"
)

func TestStore_GetAllSymbols(t *testing.T) {

	store := GetDb()
	data, err := store.GetAllSymbols()
	if err != nil {
		t.Errorf("Db error occured: %v", err)
	}
	if len(data) == 0 {
		t.Errorf("Data is empty")
	}

}

func TestStore_GetCandles(t *testing.T) {

	store := GetDb()
	unit := models.DateTruncs["1m"]

	data, err := store.GetCandles("bchbear", "usd", unit)
	if err != nil {
		t.Errorf("Db error occured: %v", err)
	}
	if len(data) == 0 {
		t.Errorf("Data is empty")
	}

}
