package database

import (
	"context"
	"fmt"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(ctx context.Context) (*mongo.Client,error){
	dbURL := viper.GetString("DBurl")
	fmt.Println(dbURL)
	opts := options.Client().ApplyURI(dbURL)
	mongoClient,err := mongo.Connect(ctx,opts)
	if err != nil {
		fmt.Println("error while connecting database",err)
		return nil,err
	}
	return mongoClient,nil
}