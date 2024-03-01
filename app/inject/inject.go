package inject

import (
	"context"
	"urlShortner/app/controllers"
	"urlShortner/app/repositories"
	"urlShortner/app/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type DI struct {
	DB *mongo.Database
}

func InitDI(db *mongo.Database) (*DI) {
	return &DI{
		DB: db,
	}
}

func (di *DI) Inject(ctx context.Context,router *gin.Engine){
	controllers.InitURLCache()
	urlRepo := repositories.NewURLRepo(di.DB)
	urlController := controllers.NewURLController(ctx,urlRepo)
	routes.RegisterURLRoutes(router,urlController)
}