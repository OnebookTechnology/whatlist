package server

import (
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"strconv"
)

func GetBrowseListRecord(ctx *gin.Context) {
	crossDomain(ctx)
	userId := ctx.Query("user_id")
	pageNumStr := ctx.Query("page_num")
	pageCountStr := ctx.Query("page_count")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "invalid page_num: %s", pageNumStr)
		return
	}
	pageCount, err := strconv.Atoi(pageCountStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "invalid page_count: %s", pageCountStr)
		return
	}
	records, err := server.DB.FindListRecordByUserId(userId, pageNum, pageCount)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when FindListRecordByUserId. err: %s", err.Error())
		return
	}
	res, err := jsoniter.MarshalToString(records)
	if err != nil {
		sendJsonResponse(ctx, Err, "MarshalToString err when GetBrowseListRecord. err %s", err.Error())
		return
	}
	sendJsonResponse(ctx, OK, "%s", res)
}

func GetBrowseBookRecord(ctx *gin.Context) {
	crossDomain(ctx)
	userId := ctx.Query("user_id")
	pageNumStr := ctx.Query("page_num")
	pageCountStr := ctx.Query("page_count")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "invalid page_num: %s", pageNumStr)
		return
	}
	pageCount, err := strconv.Atoi(pageCountStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "invalid page_count: %s", pageCountStr)
		return
	}
	records, err := server.DB.FindBookRecordByUserId(userId, pageNum, pageCount)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when FindBookRecordByUserId. err: %s", err.Error())
		return
	}
	res, err := jsoniter.MarshalToString(records)
	if err != nil {
		sendJsonResponse(ctx, Err, "MarshalToString err when GetBrowseBookLRecord. err %s", err.Error())
		return
	}
	sendJsonResponse(ctx, OK, "%s", res)
}
