package server

import (
	"github.com/OnebookTechnology/WhatList/server/models"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"strconv"
	"sync"
)

var UserMap sync.Map //map[uint64]*models.User

type RecommendResponse struct {
	returnList []*models.Book
}

const ReturnCount = 10

func recommend(ctx *gin.Context) {
	crossDomain(ctx)
	uidStr := ctx.Query("user_id")
	pageNumStr := ctx.Query("page_num")
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		sendJsonResponse(ctx, Err, "recommend uid ParseUint err: %s", err.Error())
		return
	}
	pageNum, err := strconv.ParseUint(pageNumStr, 10, 64)
	if err != nil {
		sendJsonResponse(ctx, Err, "recommend pageNum ParseUint err: %s", err.Error())
		return
	}
	u, ok := UserMap.Load(uid)
	if !ok {
		//TODO：查SQL，获取user
	}
	user := u.(*models.User)
	// update recommend maps
	doRecommend(user)

	sl, ok := UserSuitMap.Load(uid)
	if !ok {
		sendJsonResponse(ctx, Err, "recommend UserSuitMap is empty. uid: %s", uidStr)
		return
	}

	usl30, ok := UserUnSuit30Map.Load(uid)
	if !ok {
		sendJsonResponse(ctx, Err, "recommend UserUnSuit30Map is empty. uid: %s", uidStr)
		return
	}

	usl10, ok := UserUnSuit10Map.Load(uid)
	if !ok {
		sendJsonResponse(ctx, Err, "recommend UserUnSuit10Map is empty. uid: %s", uidStr)
		return
	}

	suit := sl.([]*models.Book)
	unsuit30 := usl30.([]*models.Book)
	unsuit10 := usl10.([]*models.Book)
	var returnCount = ReturnCount
	res := new(RecommendResponse)
	suitLen := len(suit[(pageNum-1)*6:])
	if suitLen < 6 {
		res.returnList = append(res.returnList, suit[(pageNum-1)*6:]...)
	} else {
		res.returnList = append(res.returnList, suit[(pageNum-1)*6:(pageNum-1)*6+6]...)
	}
	returnCount -= suitLen

	unsuit30Len := len(unsuit30[(pageNum-1)*3:])
	if unsuit30Len < 3 {
		res.returnList = append(res.returnList, unsuit30[(pageNum-1)*3:]...)
	} else {
		res.returnList = append(res.returnList, unsuit30[(pageNum-1)*3:(pageNum-1)*3+3]...)
	}

	returnCount -= unsuit30Len

	res.returnList = append(res.returnList, unsuit10[(pageNum-1):(pageNum-1)+1]...)
	returnCount -= 1
	//凑够10个
	if returnCount > 0 {
		res.returnList = append(res.returnList, unsuit10[pageNum*5:pageNum*5+uint64(returnCount)]...)
	}

	response, err := jsoniter.MarshalToString(res)
	if err != nil {
		sendJsonResponse(ctx, Err, "recommend MarshalToString err: %s", err.Error())
		return
	}
	sendJsonResponse(ctx, OK, "%s", response)
}
