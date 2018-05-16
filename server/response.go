package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func crossDomain(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "Authorization, Origin, No-Cache, X-Requested-With, Content-Range, X_FILENAME, If-Modified-Since, Pragma, Last-Modified, Cache-Control, Expires, Content-Type, X-E4M-With")
}

type JsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func sendJsonResponse(ctx *gin.Context, code int, format string, values ...interface{}) {
	retMsg := fmt.Sprintf(format, values...)
	logger.Info("[", ctx.Request.URL, "]", "code:", code, "response:", retMsg, "Remote:", ctx.Request.RemoteAddr)
	ctx.JSON(http.StatusOK, JsonResponse{
		Code:    code,
		Message: retMsg,
	})
}
