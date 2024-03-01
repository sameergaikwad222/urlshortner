package repositories

import (
	"context"
	"errors"
	"urlShortner/app/models"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLRepo struct {
	Collection *mongo.Collection
	Database *mongo.Database
}


func NewURLRepo(db *mongo.Database) *URLRepo{
	collection := db.Collection(viper.GetString("DBcollection"))
	return &URLRepo{
		Database: db,
		Collection: collection,
	}
}


func (u *URLRepo) GetSingleURL(ctx context.Context, shortpath string) (*models.URLS,error){
	//Return Single URL object
	var url *models.URLS
	err := u.Collection.FindOne(ctx,bson.M{"shortpath":shortpath}).Decode(&url)
	if err != nil{
		return nil,err
	}
	return url,err
}

func (u *URLRepo) UpdateSingleURL(ctx context.Context,shortpath string, urlUpdate interface{}) (*mongo.UpdateResult,error){
	//Update Single URL object
	result,err := u.Collection.UpdateOne(ctx,bson.M{"shortpath":shortpath},urlUpdate)
	if err != nil {
		return nil,err
	}
	return result,nil
}

func (u *URLRepo) DeleteSingleURL(ctx context.Context,shortpath string)(*mongo.DeleteResult,error){
	//Delete Single URL object
	result,err := u.Collection.DeleteOne(ctx,bson.M{"shortpath":shortpath})
	if err != nil {
		return nil,err
	}
	return result,nil
}

func (u *URLRepo) InsertMultipleURLs(ctx context.Context, urls []interface{})(*mongo.InsertManyResult, error){
	//Insert Multiple URL short paths
	if len(urls) == 0 {
		return nil,errors.New("no documents for insertion found")
	}
	result,err := u.Collection.InsertMany(ctx,urls)
	if err != nil {
		return nil,err
	}
	return result,nil
}