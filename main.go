package main

import (
	// Go imports
	"context"
	"fmt"
	"log"
	"os"
	"time"

	// External imports
	"github.com/asaskevich/govalidator"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	// Internal imports
	_ "cryptocurrency-portfolio/docs"
	"cryptocurrency-portfolio/handler"
)

const (
	dbName                       = "crypto-db"
	cryptoCurrencyCollectionName = "crypto_currencies"
	apiPort                      = ":8080"
)

// @title CRYPTOCURRENCY PORTFOLIO
// @version 1.0
// @description This is a sample CRUD operations on currency system.
func main() {

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

	app := iris.Default()
	app.UseRouter(recover.New())

	config := &swagger.Config{
		URL:         "http://localhost:8080/swagger/doc.json",
		DeepLinking: true,
	}
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))

	api := mvc.New(app.Party("/"))

	// I have 1 collection and handle it on main
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
