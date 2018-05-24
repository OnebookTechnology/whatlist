package server

import (
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"strconv"
)

func AddLikeNum(ctx *gin.Context) {
	crossDomain(ctx)
	idStr := ctx.Query("discover_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendJsonResponse(ctx, Err, "invalid id: %s", idStr)
		return
	}
	err = server.DB.AddLikeNum(id)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when AddLikeNum. id: %s", idStr)
		return
	}
	sendJsonResponse(ctx, OK, "%s", "ok")
}

func GetDiscoverDetail(ctx *gin.Context) {
	crossDomain(ctx)
	idStr := ctx.Query("discover_id")
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
	res, err := jsoniter.MarshalToString(detail)
	if err != nil {
		sendJsonResponse(ctx, Err, "MarshalToString err. id: %s", idStr)
		return
	}
	sendJsonResponse(ctx, OK, "%s", res)
}
