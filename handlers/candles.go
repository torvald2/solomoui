package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/torvald2/solanamon/models"
	"github.com/torvald2/solanamon/services"
)

func GetCandleHandler(ps services.CandleGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		symbol := r.URL.Query().Get("symbol")
		unitRaw := r.URL.Query().Get("unit")
		pair := r.URL.Query().Get("pair")
		unit, ok := models.DateTruncs[models.Unit(unitRaw)]
		if !ok || len(symbol) == 0 {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]string{"err": "Bad Request"})
			return
		}
		data, err := ps.GetCandles(symbol, pair, unit)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(map[string]error{"err": err})
		}

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(data)

	})
}
