package server

import (
	"database/sql"
	"fmt"
	"github.com/OnebookTechnology/whatlist/server/models"
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
	userId := ctx.Query("user_id")
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
	b.IsCollect = false

	_, err = server.DB.FindBiggieIsCollected(userId, biggieId)
	if err != nil {
		if err != sql.ErrNoRows {
			sendFailedResponse(ctx, Err, "db error when FindBiggieIsCollected. err:", err)
			return
		}
	} else {
		b.IsCollect = true
	}
	res := &ResData{
		Biggie: b,
	}
	sendSuccessResponse(ctx, res)
	return
}

func GetBiggieList(ctx *gin.Context) {
	crossDomain(ctx)
	userId := ctx.Query("user_id")
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

	bs, err := server.DB.FindListsByBiggie(userId, biggieId, pageNum, pageCount)
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
	userId := ctx.Query("user_id")
	idStr := ctx.Query("list_id")
	listId, err := strconv.Atoi(idStr)
	if err != nil {
		sendFailedResponse(ctx, Err, "parse list_id error:", err, "list_id:", idStr)
		return
	}

	isPayed := false
	//查询是否有购买过
	_, err = server.DB.FindListPurchaseRecord(userId, listId)
	if err != nil {
		//没买过，不返会书单目录
		if err != sql.ErrNoRows {
			sendFailedResponse(ctx, Err, "db error when FindListPurchaseRecord. error: ", err.Error())
			return

		}
	} else {
		isPayed = true
	}

	b, err := server.DB.FindBiggieListById(listId)
	if err != nil {
		sendFailedResponse(ctx, Err, "db error when FindBiggieListById. error: ", err.Error())
		return
	}

	bs, err := server.DB.FindBiggieListBooks(listId)
	if err != nil {
		sendFailedResponse(ctx, Err, "db error when FindBiggieListBooks. error: ", err.Error())
		return
	}

	err = server.DB.AddClickCount(listId)
	if err != nil {
		logger.Error("db error when AddClickCount. error: ", err.Error())
		return
	}

	res := &ResData{
		BiggieList:  b,
		BiggieBooks: bs,
		IsPayed:     isPayed,
	}
	sendSuccessResponse(ctx, res)
	return
}

func GetLatestBiggieList(ctx *gin.Context) {
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

	bs, err := server.DB.FindLatestBiggieList(pageNum, pageCount)
	if err != nil {
		sendFailedResponse(ctx, Err, "db error when FindRecommendBiggies. err:", err)
		return
	}
	res := &ResData{
		BiggieLists: bs,
	}
	sendSuccessResponse(ctx, res)
	return
}

type CollectReq struct {
	UserId   string `json:"user_id" form:"user_id"`
	BiggieId int    `json:"biggie_id" form:"biggie_id"`
	Page
}

func CollectBiggie(ctx *gin.Context) {
	crossDomain(ctx)
	var req CollectReq
	if err := ctx.BindJSON(&req); err == nil {
		c := &models.BiggieCollect{
			UserId:   req.UserId,
			BiggieId: req.BiggieId,
		}
		err := server.DB.AddCollectBiggie(c)
		if err != nil {
			sendFailedResponse(ctx, Err, "AddCollectBiggie err:", err)
			return
		}

		sendSuccessResponse(ctx, nil)
		return

	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}

func RemoveBiggie(ctx *gin.Context) {
	crossDomain(ctx)
	var req CollectReq
	if err := ctx.BindJSON(&req); err == nil {
		err := server.DB.DeleteCollectBiggie(req.UserId, req.BiggieId)
		if err != nil {
			sendFailedResponse(ctx, Err, "DeleteCollectBiggie err:", err)
			return
		}
		sendSuccessResponse(ctx, nil)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}

func GetCollectBiggie(ctx *gin.Context) {
	crossDomain(ctx)
	var req CollectReq
	if err := ctx.ShouldBindQuery(&req); err == nil {
		fmt.Println(req.UserId, req.PageNum, req.PageCount)
		bs, err := server.DB.FindCollectBiggies(req.UserId, req.PageNum, req.PageCount)
		if err != nil {
			if err == sql.ErrNoRows {
				sendFailedResponse(ctx, NoResultErr, "FindCollectBiggies err:", err)
				return
			}
			sendFailedResponse(ctx, Err, "FindCollectBiggies err:", err)
			return
		}
		res := &ResData{
			Biggies: bs,
		}
		sendSuccessResponse(ctx, res)
		return
	} else {
		sendFailedResponse(ctx, Err, "BindJSON err:", err)
		return
	}
}
