package server

import (
	"database/sql"
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
	return
}

// 添加喜爱图书 CX
func AddInterestedBook(ctx *gin.Context) {
	crossDomain(ctx)
	isbnStr := ctx.Query("isbn")
	userID := ctx.Query("userID")
	if isbnStr == "" || userID == "" {
		sendJsonResponse(ctx, Err, "%s", "Empty params userID or isbn")
	}
	isbn, err := strconv.ParseUint(isbnStr, 10, 64)
	if err != nil {
		sendJsonResponse(ctx, Err, "Can not convert isbn to uint in AddInterestedBook api. "+
			"Error: %s, isbn: %s", err.Error(), isbnStr)
		return
	}
	err = server.DB.AddInterestedBook(userID, isbn)
	if err != nil {
		sendJsonResponse(ctx, Err, "AddInterestedBook error in AddInterestedBook api. Error: %s",
			err.Error())
		return
	}
	sendJsonResponse(ctx, OK, "%s", "add interested book success")
	return
}

// 删除喜爱图书
func DeleteInterestedBook(ctx *gin.Context) {
	crossDomain(ctx)
	isbnStr := ctx.Query("isbn")
	userID := ctx.Query("userID")
	if isbnStr == "" || userID == "" {
		sendJsonResponse(ctx, Err, "%s", "Empty params userID or isbn")
	}
	isbn, err := strconv.ParseUint(isbnStr, 10, 64)
	if err != nil {
		sendJsonResponse(ctx, Err, "Can not convert isbn to uint in AddInterestedBook api. "+
			"Error: %s, isbn: %s", err.Error(), isbnStr)
		return
	}
	err = server.DB.DeleteInterestedBook(userID, isbn)
	if err != nil {
		sendJsonResponse(ctx, Err, "AddInterestedBook error in AddInterestedBook api. Error: %s",
			err.Error())
		return
	}
	sendJsonResponse(ctx, OK, "%s", "add interested book success")
	return
}

// 列出喜爱图书
func InterestedBooks(ctx *gin.Context) {
	crossDomain(ctx)
	userID := ctx.Query("userID")
	if userID == "" {
		sendJsonResponse(ctx, Err, "%s", "Empty params userID.")
		return
	}
	books, err := server.DB.GetInterestedBooksByUserID(userID)
	if err == sql.ErrNoRows {
		sendJsonResponse(ctx, NoResultErr, "%s", "No interested books.")
		return
	}
	if err != nil {
		sendJsonResponse(ctx, Err, "GetInterestedBooksByUserID error in InterestedBooks api."+
			"Error: %s", err.Error())
		return
	}
	res, err := jsoniter.MarshalToString(books)
	if err != nil {
		sendJsonResponse(ctx, Err, "MarshalToString error in InterestedBooks api.Error :%s",
			err.Error())
		return
	}
	sendJsonResponse(ctx, OK, "%s", res)
	return
}
