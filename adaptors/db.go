package adaptors

import (
	"context"
	"sync"
	"time"

	"github.com/torvald2/solanamon/config"
	"github.com/torvald2/solanamon/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DbClient struct {
	client *mongo.Client
}

func (c *DbClient) GetAllSymbols() (symbols []models.Symbol, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	options := options.Find()
	collection := c.client.Database("solana").Collection("symbols")

	cur, err := collection.Find(ctx, filter, options)
	defer cur.Close(ctx)
	if err != nil {
		return
	}
	for cur.Next(ctx) {
		var symbol models.Symbol
		err = cur.Decode(&symbol)
		if err != nil {
			return
		}
		symbols = append(symbols, symbol)

	}

	return

}

func (c *DbClient) GetCandles(symbol string, pair string, unit models.DateTrunc) (candles []models.Candle, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := c.client.Database("solana").Collection("ticker")
	objID, _ := primitive.ObjectIDFromHex(symbol)

	match := bson.D{{
		"$match", bson.D{{"symbol", objID}, {"pair", pair}},
	}}

	group := bson.D{
		{"$group", bson.D{
			{"_id", bson.D{
				{"symbol", "$symbol"},
				{"time", bson.D{
					{"$dateTrunc", bson.D{
						{"date", "$time"},
						{"unit", unit.Unit},
						{"binSize", unit.BinSize},
					}},
				},
				},
			},
			},

			{"high", bson.D{{"$max", "$price"}}},
			{"low", bson.D{{"$min", "$price"}}},
			{"open", bson.D{{"$first", "$price"}}},
			{"close", bson.D{{"$last", "$price"}}},
		},
		},
	}

	sort := bson.D{
		{"$sort", bson.D{
			{"_id.time", 1},
		}},
	}

	pipe := mongo.Pipeline{match, group, sort}
	cur, err := collection.Aggregate(ctx, pipe)
	if err != nil {
		return
	}
	err = cur.All(ctx, &candles)

	return
}

var db DbClient
var once sync.Once

func initDb() {
	conf := config.GetConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.DBConnectionString))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	db = DbClient{client: client}
}

func GetDb() *DbClient {
	once.Do(initDb)
	return &db
}
