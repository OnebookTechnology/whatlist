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
	userID := ctx.Query("user_id")
	if userID == "" {
		sendJsonResponse(ctx, Err, "%s", "Empty params user_id")
		return
	}
	if err != nil {
		sendJsonResponse(ctx, Err, "GetBookDetail isbn is invalid. isbn:%s ", isbnStr)
		return
	}

	book, err := server.DB.FindBook(isbn)
	if err != nil {
		sendJsonResponse(ctx, Err, "DB error when FindBook. err: %s", err.Error())
		return
	}

	flag, err := server.DB.IsBookInterested(book.ISBN, userID)
	if err == sql.ErrNoRows {
		flag = 0
	}
	if err != nil && err != sql.ErrNoRows {
		sendJsonResponse(ctx, Err, "DB error when IsBookInterested. Error: %s", err.Error())
		return
	}
	err = server.DB.AddBookRecord(userID, isbn)
	if err != nil {
		logger.Error("db error when AddBookRecord. err", err, "userId:", userID, "ISBN:", isbn)
	}

	// 如果flag>0,表示用户喜欢了这本书
	if flag > 0 {
		book.IsInterested = true
	}

	resp, _ := jsoniter.MarshalToString(book)
	sendJsonResponse(ctx, OK, "%s", resp)
	return
}

// 添加喜爱图书 CX
func AddInterestedBook(ctx *gin.Context) {
	crossDomain(ctx)
	isbnStr := ctx.Query("isbn")
	userID := ctx.Query("user_id")
	if isbnStr == "" || userID == "" {
		sendJsonResponse(ctx, Err, "%s", "Empty params user_id or isbn")
	}
	isbn, err := strconv.ParseUint(isbnStr, 10, 64)
	if err != nil {
		sendJsonResponse(ctx, Err, "Can not convert isbn to uint in AddInterestedBook api. "+
			"Error: %s, isbn: %s", err.Error(), isbnStr)
		return
	}

	flag, err := server.DB.IsBookInterested(isbn, userID)
	if err == sql.ErrNoRows {
		flag = 0
	}
	if err != nil && err != sql.ErrNoRows {
		sendJsonResponse(ctx, Err, "IsBookInterested error in AddInterestedBook api. "+
			"Error: %s", err.Error())
		return
	}

	// 如果flag>0，表示已经添加过喜欢了
	if flag > 0 {
		sendJsonResponse(ctx, Err, "Already added this book")
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
	userID := ctx.Query("user_id")
	if isbnStr == "" || userID == "" {
		sendJsonResponse(ctx, Err, "%s", "Empty params user_id or isbn")
		return
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
	sendJsonResponse(ctx, OK, "%s", "delete interested book success")
	return
}

// 列出喜爱图书
func CategoryBooks(ctx *gin.Context) {
	crossDomain(ctx)
	categoryIdStr := ctx.Query("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "%s", "invalid params category_id.")
		return
	}
	pageNumStr := ctx.Query("page_num")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "%s", "invalid params page_num.")
		return
	}
	pageCountStr := ctx.Query("page_count")
	pageCount, err := strconv.Atoi(pageCountStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "%s", "invalid params page_count.")
		return
	}

	books, err := server.DB.FindBookByCateGory(categoryId, pageNum, pageCount)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when FindBookByCateGory."+
			"Error: %s", err.Error())
		return
	}
	res, err := jsoniter.MarshalToString(books)
	if err != nil {
		sendJsonResponse(ctx, Err, "MarshalToString error.Error :%s",
			err.Error())
		return
	}
	sendJsonResponse(ctx, OK, "%s", res)
	return
}

// 列出喜爱图书
func InterestedBooks(ctx *gin.Context) {
	crossDomain(ctx)
	userID := ctx.Query("user_id")
	if userID == "" {
		sendJsonResponse(ctx, Err, "%s", "Empty params user_id.")
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
