package main

import (
	"context"
	"cryptocurrency-portfolio/handler"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

const (
	dbName                       = "crypto-db"
	cryptoCurrencyCollectionName = "crypto_currencies"
	apiPort                      = ":8080"
)

func main() {
	//"mongodb://localhost:27017"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err, "can not connect mongoDB")
	}

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			log.Fatal(err, "can not disconnect mongoDB")
		}
	}()

	if err = mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err, "can not ping mongoDB")
	}

	fmt.Println("connected to the mongo db")

	app := iris.New()
	app.UseRouter(recover.New())
	api := mvc.New(app.Party("/"))

	// I have 1 collection and handled in main
	coll := mongoClient.Database(dbName).Collection(cryptoCurrencyCollectionName)
	api.Handle(&handler.Handler{
		MongoCollection: coll,
	})

	if err := app.Listen(apiPort); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}
