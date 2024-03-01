package controllers

import (
	"context"
	"net/http"
	"time"
	"urlShortner/app/models"
	"urlShortner/app/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
)

var urlCache *cache.Cache

func InitURLCache(){
	urlCache = cache.New(60*time.Minute,120*time.Minute)
}

func GetUrlcache(shortpath string)(interface{},bool){
	return urlCache.Get(shortpath)
}
func SetUrlcache(url interface{},shortpath string){
	urlCache.Set(shortpath,url, 15 * time.Minute)
}



type URLController struct {
	urlrepo *repositories.URLRepo
	ctx		context.Context
}

func NewURLController(ctx context.Context, urlrepo *repositories.URLRepo) *URLController{
	return &URLController{
		urlrepo: urlrepo,
		ctx: ctx,
	}
}


func (uc *URLController) RedirectDestination(ctx *gin.Context){
	shortpath := ctx.Param("shortpath")
	if shortpath == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{"status":"Failed","message":"Invalid Short Path"})				
	}

	//Check Cache get/set first
	cachedurl,ok := GetUrlcache(shortpath);if ok {
		URL, ok := cachedurl.(models.URLS); if ok {
			ctx.Redirect(http.StatusMovedPermanently,URL.Destination)
			return
		}
	}

	url,err := uc.urlrepo.GetSingleURL(ctx,shortpath)
	if err!= nil {
		ctx.JSON(http.StatusBadGateway,gin.H{"status":"failed","message":"Error while fetching url"})
	}

	//Cache URL
	SetUrlcache(url,shortpath)
	ctx.Redirect(http.StatusMovedPermanently,url.Destination)	
}
func (uc *URLController) GetDestination(ctx *gin.Context){
	shortpath := ctx.Param("shortpath")
	if shortpath == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{"status":"Failed","message":"Invalid Short Path"})				
	}

	//Check Cache get/set first
	cachedurl,ok := GetUrlcache(shortpath);if ok {
		URL, ok := cachedurl.(models.URLS); if ok {
			ctx.JSON(http.StatusOK,gin.H{"status":"success","data":URL})	
			return
		}
	}

	url,err := uc.urlrepo.GetSingleURL(ctx,shortpath)
	if err!= nil {
		ctx.JSON(http.StatusBadGateway,gin.H{"status":"failed","message":"Error while fetching url"})
	}

	//Cache URL
	SetUrlcache(url,shortpath)
	ctx.JSON(http.StatusOK,gin.H{"status":"success","data":url})	
}

func (uc *URLController) UpdateDestination(ctx *gin.Context){
	shortpath := ctx.Param("shortpath")
	if shortpath == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{"status":"Failed","message":"Invalid Short Path"})				
	}
	var urlUpdate  interface{}
	err := ctx.ShouldBindJSON(&urlUpdate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"status":"Failed","message":"Invalid Update Payload"})	
	}

	updatedURL,err := uc.urlrepo.UpdateSingleURL(ctx,shortpath, urlUpdate)
	if err!= nil {
		ctx.JSON(http.StatusBadGateway,gin.H{"status":"failed","message":"Error while fetching url"})
	}
	ctx.JSON(http.StatusOK,gin.H{"status":"success","data":updatedURL})	
}

func (uc *URLController) DeleteDestination(ctx *gin.Context){
	shortpath := ctx.Param("shortpath")
	if shortpath == "" {
		ctx.JSON(http.StatusBadRequest,gin.H{"status":"Failed","message":"Invalid Short Path"})				
	}

	result,err := uc.urlrepo.DeleteSingleURL(ctx,shortpath)
	if err!= nil {
		ctx.JSON(http.StatusBadGateway,gin.H{"status":"failed","message":"Error while fetching url"})
	}
	ctx.JSON(http.StatusOK,gin.H{"status":"success","data":result})	
}

func (uc *URLController) InsertManyDestination(ctx *gin.Context){
	var urlsStrings []string
	err := ctx.ShouldBindJSON(&urlsStrings)
	if err!= nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"status":"failed","message":"Invalid payloads"})
	}
	var urls []interface{}
	for _,urlstring := range urlsStrings{

		shortstr := uuid.NewString()[0:8]
		var urlInstance = models.URLS{
			ShortPath: string(shortstr),
			Destination: urlstring,
		}
		urls = append(urls, urlInstance)
	}
	result,err := uc.urlrepo.InsertMultipleURLs(ctx,urls)
	if err!= nil {
		ctx.JSON(http.StatusBadGateway,gin.H{"status":"failed","message":"Error while Inserting urls"})
	}
	ctx.JSON(http.StatusOK,gin.H{"status":"success","data":result})	
}