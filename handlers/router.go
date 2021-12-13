package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/torvald2/solanamon/adaptors"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	store := adaptors.GetDb()

	candlesHandler := GetCandleHandler(store)
	symbolsHandler := GetSymbolsHandler(store)

	api := r.PathPrefix("/api/").Subrouter()

	api.Handle("/symbols", symbolsHandler).Methods("GET")
	api.Handle("/candles", candlesHandler).Methods("GET")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	return r

}
