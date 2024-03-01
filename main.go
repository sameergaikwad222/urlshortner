package main

import (
	"context"
	"fmt"
	"log"
	"urlShortner/app/config"
	"urlShortner/app/database"
	"urlShortner/app/inject"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	r:= gin.Default()
	ctx := context.Background()
	defer ctx.Done()
	config.InitConfig()
	mongoClient,err := database.ConnectMongoDB(ctx)
	if err != nil {
		panic("error while connecting database")
	}
	db := mongoClient.Database(viper.GetString("DBname"))
	DI := inject.InitDI(db)
	DI.Inject(ctx,r)
	err = r.Run(fmt.Sprintf(":%s",viper.GetString("port")))
	if err != nil {
		log.Fatal("Error while starting server")
	}
}