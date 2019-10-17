package app

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/viknesh-nm/sellerapp/backend"
	"github.com/viknesh-nm/sellerapp/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Init initializes the application by setting up various configurations like
database connections,
*/
func Init() (*echo.Echo, error) {

	db, err := gorm.Open("mysql", fmt.Sprintf(
		"%s:%s%s",
		conf.Config.MySQL.User,
		conf.Config.MySQL.Password,
		conf.Config.MySQL.DSN,
	))
	if err != nil {
		return nil, err
	}

	mongoDB, err := connectMongo(conf.Config.Mongo.DSN)
	if err != nil {
		fmt.Println(err)
	}

	err = mongoDB.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("Mongo Connected")

	// Inits Routes
	m := Routes()

	infra := &backend.Infra{
		MongoDB: mongoDB,
		MySQLDB: db,
	}

	backend.Init(infra)

	return m, nil
}

// connectMongo
func connectMongo(dsn string) (*mongo.Client, error) {

	// establish the connection to the Mongo server
	clientOptions := options.Client().ApplyURI("mongodb://" + dsn)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err == nil {
		return client, err
	}

	return nil, err
}
