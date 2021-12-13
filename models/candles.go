package models

import "time"

type Unit string
type DateTrunc struct {
	Date    string `json:"date"`
	Unit    string `json:"unit"`
	BinSize int    `json:"binSize"`
}

const (
	OneMinute      Unit = "1m"
	ThreeMinutes   Unit = "3m"
	FifeMinutes    Unit = "5m"
	FifteenMunutes Unit = "15m"
	ThirtyMunutes  Unit = "30m"

	OneHour     Unit = "1h"
	TwoHours    Unit = "2h"
	FourHours   Unit = "4h"
	TwelveHours Unit = "12h"

	OneDay    Unit = "1d"
	ThreeDays Unit = "3d"

	OneWeek Unit = "1w"
)

type Candle struct {
	Id struct {
		Symbol string    `json:"symbol" bson:"symbol"`
		Time   time.Time `json:"time" bson:"time"`
	} `json:"id" bson:"_id"`
	Hight float64 `json:"high" bson:"high"`
	Low   float64 `json:"low" bson:"low"`
	Open  float64 `json:"open" bson:"open"`
	Close float64 `json:"close" bson:"close"`
}

var DateTruncs = map[Unit]DateTrunc{
	OneMinute: {
		Date:    "$time",
		Unit:    "minute",
		BinSize: 1,
	},
	ThreeMinutes: {
		Date:    "$time",
		Unit:    "minute",
		BinSize: 3,
	},
	FifeMinutes: {
		Date:    "$time",
		Unit:    "minute",
		BinSize: 5,
	},
	FifteenMunutes: {
		Date:    "$time",
		Unit:    "minute",
		BinSize: 15,
	},
	ThirtyMunutes: {
		Date:    "$time",
		Unit:    "minute",
		BinSize: 30,
	},
	OneHour: {
		Date:    "$time",
		Unit:    "hour",
		BinSize: 1,
	},
	TwoHours: {
		Date:    "$time",
		Unit:    "hour",
		BinSize: 2,
	},
	FourHours: {
		Date:    "$time",
		Unit:    "hour",
		BinSize: 4,
	},
	TwelveHours: {
		Date:    "$time",
		Unit:    "hour",
		BinSize: 12,
	},
	OneDay: {
		Date:    "$time",
		Unit:    "day",
		BinSize: 1,
	},
	ThreeDays: {
		Date:    "$time",
		Unit:    "day",
		BinSize: 3,
	},
	OneWeek: {
		Date:    "$time",
		Unit:    "week",
		BinSize: 1,
	},
}
