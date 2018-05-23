package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LoadRouter(router *gin.Engine) {
	router.GET("/whoami", func(context *gin.Context) {
		context.String(http.StatusOK, "I am %s. Time,%s", server.ServerName, time.Now().String())
	})

	router.GET("/recommend", recommend)
	router.GET("/sign", Sign)
	router.GET("/update", UpdateUserData)
	router.GET("/gettags", GetAllTags)
	router.GET("/gettagsnum", GetAllTagsNumber)
	router.GET("/bookdetail", GetBookDetail)

	router.GET("/latestlists", LatestLists)
	router.GET("/HeatestLists", HeatLists)
	router.GET("/recommendlists", RecommendLists)
	router.GET("/listdetail", ListDetail)
}
