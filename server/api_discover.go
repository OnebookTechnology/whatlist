package server

import (
	"database/sql"
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
	err = server.DB.GetDiscoverDetailIsThumb(id, uid)
	if err != nil && err == sql.ErrNoRows {
		err = server.DB.ThumbsUpDiscover(id, uid)
		if err != nil {
			sendJsonResponse(ctx, Err, "db error when ThumbsUpDiscover. err: %s", err.Error())
			return
		}
		err = server.DB.AddLikeNum(id)
		if err != nil {
			sendJsonResponse(ctx, Err, "db error when AddLikeNum. err: %s", err.Error())
			return
		}
	} else if err != nil {
		sendJsonResponse(ctx, Err, "db error when AddLikeNum. err: %s", err.Error())
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
		logger.Error("db error when AddReadNum. err:", err, "id:", id)
	}
	detail, err := server.DB.GetDiscoverDetail(id)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when GetDiscoverDetail. id: %s", idStr)
		return
	}
	err = server.DB.GetDiscoverDetailIsThumb(id, uid)
	if err != nil && err != sql.ErrNoRows {
		sendJsonResponse(ctx, Err, "db error when GetDiscoverDetailIsThumb. id: %s", idStr)
		return
	}
	if err == sql.ErrNoRows {
		detail.IsThumb = false
	} else {
		detail.IsThumb = true
	}

	res, err := jsoniter.MarshalToString(detail)
	if err != nil {
		sendJsonResponse(ctx, Err, "MarshalToString err. id: %s", idStr)
		return
	}
	sendJsonResponse(ctx, OK, "%s", res)
}

func GetDiscoverList(ctx *gin.Context) {
	crossDomain(ctx)
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
	uid := ctx.Query("user_id")
	list, err := server.DB.GetDiscoverList(uid, pageNum, pageCount)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when GetDiscoverList: %s", err.Error())
		return
	}
	res, err := jsoniter.MarshalToString(list)
	if err != nil {
		sendJsonResponse(ctx, Err, "MarshalToString err when GetDiscoverList. err %s", err.Error())
		return
	}
	sendJsonResponse(ctx, OK, "%s", res)
}
