package server

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetLatestBiggie(ctx *gin.Context) {
	crossDomain(ctx)
	pageNumStr := ctx.Query("page_num")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		sendFailedResponse(ctx, Err, "parse page_num error:", err, "page_num:", pageNumStr)
		return
	}
	pageCountStr := ctx.Query("page_count")
	pageCount, err := strconv.Atoi(pageCountStr)
	if err != nil {
		sendFailedResponse(ctx, Err, "parse page_count error:", err, "page_count:", pageCountStr)
		return
	}

	bs, err := server.DB.FindLatestBiggie(pageNum, pageCount)
	if err != nil {
		sendFailedResponse(ctx, Err, "db error when FindLatestBiggie. err:", err)
		return
	}
	res := &ResData{
		Biggies: bs,
	}
	sendSuccessResponse(ctx, res)
	return
}

func GetBiggie(ctx *gin.Context) {
	crossDomain(ctx)
	idStr := ctx.Query("biggie_id")
	biggieId, err := strconv.Atoi(idStr)
	if err != nil {
		sendFailedResponse(ctx, Err, "parse biggie_id error:", err, "biggie_id:", idStr)
		return
	}

	b, err := server.DB.FindBiggieById(biggieId)
	if err != nil {
		sendFailedResponse(ctx, Err, "db error when FindBiggieById. err:", err)
		return
	}
	res := &ResData{
		Biggie: b,
	}
	sendSuccessResponse(ctx, res)
	return
}

func GetBiggieList(ctx *gin.Context) {
	crossDomain(ctx)
	idStr := ctx.Query("biggie_id")
	biggieId, err := strconv.Atoi(idStr)
	if err != nil {
		sendFailedResponse(ctx, Err, "parse biggie_id error:", err, "biggie_id:", idStr)
		return
	}
	pageNumStr := ctx.Query("page_num")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		sendFailedResponse(ctx, Err, "parse page_num error:", err, "page_num:", pageNumStr)
		return
	}
	pageCountStr := ctx.Query("page_count")
	pageCount, err := strconv.Atoi(pageCountStr)
	if err != nil {
		sendFailedResponse(ctx, Err, "parse page_count error:", err, "page_count:", pageCountStr)
		return
	}

	bs, err := server.DB.FindListsByBiggie(biggieId, pageNum, pageCount)
	if err != nil {
		sendFailedResponse(ctx, Err, "db error when FindLatestBiggie. err:", err)
		return
	}
	res := &ResData{
		BiggieLists: bs,
	}
	sendSuccessResponse(ctx, res)
	return
}

func GetRecommendBiggie(ctx *gin.Context) {
	crossDomain(ctx)
	pageNumStr := ctx.Query("page_num")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		sendFailedResponse(ctx, Err, "parse page_num error:", err, "page_num:", pageNumStr)
		return
	}
	pageCountStr := ctx.Query("page_count")
	pageCount, err := strconv.Atoi(pageCountStr)
	if err != nil {
		sendFailedResponse(ctx, Err, "parse page_count error:", err, "page_count:", pageCountStr)
		return
	}
	bs, err := server.DB.FindRecommendBiggies(pageNum, pageCount)
	if err != nil {
		sendFailedResponse(ctx, Err, "db error when FindRecommendBiggies. err:", err)
		return
	}
	res := &ResData{
		Biggies: bs,
	}

	sendSuccessResponse(ctx, res)
	return
}

func GetBiggieListBooks(ctx *gin.Context) {
	crossDomain(ctx)
	idStr := ctx.Query("list_id")
	listId, err := strconv.Atoi(idStr)
	if err != nil {
		sendFailedResponse(ctx, Err, "parse list_id error:", err, "list_id:", idStr)
		return
	}
	bs, err := server.DB.FindBiggieListBooks(listId)
	res := &ResData{
		BiggieBooks: bs,
	}
	sendSuccessResponse(ctx, res)
	return
}

func CollectBiggie(ctx *gin.Context) {
	crossDomain(ctx)
}
