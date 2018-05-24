package server

import (
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"strconv"
)

func AddLikeNum(ctx *gin.Context) {
	crossDomain(ctx)
	idStr := ctx.Query("discover_id")
	uid := ctx.Query("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "invalid id: %s", idStr)
		return
	}
	err = server.DB.ThumbsUpDiscover(id, uid)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when ThumbsUpDiscover. did: %s", idStr)
		return
	}
	err = server.DB.AddLikeNum(id)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when AddLikeNum. did: %s", idStr)
		return
	}
	sendJsonResponse(ctx, OK, "%s", "ok")
}

func SubLikeNum(ctx *gin.Context) {
	crossDomain(ctx)
	idStr := ctx.Query("discover_id")
	uid := ctx.Query("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "invalid id: %s", idStr)
		return
	}
	err = server.DB.CancelThumbsUpDiscover(id, uid)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when CancelThumbsUpDiscover. did: %s", idStr)
		return
	}
	err = server.DB.SubLikeNum(id)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when SubLikeNum. did: %s", idStr)
		return
	}
	sendJsonResponse(ctx, OK, "%s", "ok")
}

func GetDiscoverDetail(ctx *gin.Context) {
	crossDomain(ctx)
	idStr := ctx.Query("discover_id")
	uid := ctx.Query("user_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "invalid id: %s", idStr)
		return
	}
	err = server.DB.AddReadNum(id)
	if err != nil {
		logger.Error("db error when AddReadNum. id:", id)
		return
	}
	detail, err := server.DB.GetDiscoverDetail(id)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when GetDiscoverDetail. id: %s", idStr)
		return
	}
	err = server.DB.GetDiscoverDetailIsThumb(id, uid)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when GetDiscoverDetailIsThumb. id: %s", idStr)
		return
	}
	res, err := jsoniter.MarshalToString(detail)
	if err != nil {
		sendJsonResponse(ctx, Err, "MarshalToString err. id: %s", idStr)
		return
	}
	sendJsonResponse(ctx, OK, "%s", res)
}
