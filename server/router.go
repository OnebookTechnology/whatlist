package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LoadRouter(router *gin.Engine) {
	router.GET("/whoami", func(context *gin.Context) {
		context.String(http.StatusOK, "I am %s", server.ServerName, time.Now().String())
	})

	router.GET("/recommend", recommend)
	router.GET("/sign", Sign)
	router.GET("/update", UpdateUserData)

	router.OPTIONS("/word", func(c *gin.Context) {
		crossDomain(c)
	})
	router.OPTIONS("/addbook", func(c *gin.Context) {
		crossDomain(c)
	})
	router.POST("/word", ParseWord)
	router.POST("/addbook", AddBook)
}
