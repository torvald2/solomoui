package models

type Symbol struct {
	Name   string   `json:"label" bson:"name"`
	Symbol string   `json:"symbol" bson:"symbol"`
	Image  string   `json:"image" bson:"image"`
	Id     string   `json:"id" bson:"_id"`
	Pairs  []string `json:"pairs" bson:"pairs"`
}
