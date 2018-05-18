package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRouter(router *gin.Engine) {
	router.GET("/whoami", func(context *gin.Context) {
		context.String(http.StatusOK, "I am %s", server.ServerName)
	})

	router.GET("/recommend", recommend)
	router.GET("/sign", Sign)
	router.GET("/update", UpdateUserData)

	router.POST("/word", ParseWord)
	router.POST("/addbook",AddBook)
}
