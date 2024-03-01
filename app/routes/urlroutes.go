package routes

import (
	"net/http"
	"time"
	"urlShortner/app/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterURLRoutes(route *gin.Engine,uc *controllers.URLController){
	route.GET("",func(ctx *gin.Context){
		ctx.JSON(http.StatusOK,gin.H{"status":"success","data":time.Now()})
	})
	route.GET("/:shortpath",uc.RedirectDestination)
	route.GET("/get/:shortpath",uc.GetDestination)
	route.PATCH("/:shortpath",uc.UpdateDestination)
	route.DELETE("/:shortpath",uc.DeleteDestination)
	route.POST("",uc.InsertManyDestination)
}