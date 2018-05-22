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

	router.OPTIONS("/word", func(c *gin.Context) {
		crossDomain(c)
	})
	router.OPTIONS("/addbook", func(c *gin.Context) {
		crossDomain(c)
	})
	router.POST("/word", ParseWord)
	router.POST("/addbook", AddBook)

	router.GET("/latestlists", LatestLists)
	router.GET("/listdetail", ListDetail)
}
