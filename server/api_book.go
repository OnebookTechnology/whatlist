package server

import (
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"strconv"
)

func AddBook(c *gin.Context) {
	crossDomain(c)
}

func GetBookDetail(ctx *gin.Context) {
	crossDomain(ctx)
	isbnStr := ctx.Query("isbn")
	isbn, err := strconv.ParseUint(isbnStr, 10, 64)
	if err != nil {
		sendJsonResponse(ctx, Err, "GetBookDetail isbn is invalid. isbn:%s ", isbnStr)
		return
	}

	book, err := server.DB.FindBook(isbn)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when FindBook. err: %s", err.Error())
		return
	}

	resp, _ := jsoniter.MarshalToString(book)
	sendJsonResponse(ctx, OK, "%s", resp)
}
