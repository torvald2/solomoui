package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/torvald2/solanamon/services"
)

func GetSymbolsHandler(ps services.SymbolsGetter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data, err := ps.GetAllSymbols()
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(map[string]error{"err": err})
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(data)

	})
}
