package server

import (
	"database/sql"
	"fmt"
	"github.com/OnebookTechnology/whatlist/server/models"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"strconv"
)

type RecommendResponse struct {
	ReturnList []*models.Book `json:"return_list"`
}

const ReturnCount = 10

func recommend(ctx *gin.Context) {
	crossDomain(ctx)
	user_id := ctx.Query("user_id")
	if user_id == "" {
		sendJsonResponse(ctx, Err, "recommend need user_id")
		return
	}
	pageNumStr := ctx.Query("page_num")
	if pageNumStr == "" {
		sendJsonResponse(ctx, Err, "recommend need page_num")
		return
	}
	pageNum, err := strconv.ParseUint(pageNumStr, 10, 64)
	if err != nil {
		sendJsonResponse(ctx, Err, "recommend pageNum ParseUint err: %s", err.Error())
		return
	}
	var user *models.User
	u, ok := UserMap.Load(user_id)
	if !ok {
		//TODO：查SQL，获取user
		user, err = server.DB.FindUser(user_id)
		if err != nil {
			if err == sql.ErrNoRows {
				sendJsonResponse(ctx, NoResultErr, "db error when FindUser. err: %s", err.Error())
				return
			}
			sendJsonResponse(ctx, Err, "db error when FindUser. err: %s", err.Error())
			return
		}
		user.NeedUpdateRecommend = true
	} else {
		user = u.(*models.User)
	}
	// update recommend maps
	doRecommend(user)

	var suit, unsuit10, unsuit30 *ListResult
	sl, ok := UserSuitMap.Load(user_id)
	if !ok {
		sendJsonResponse(ctx, Err, "recommend UserSuitMap is empty. uid: %s", user_id)
		return
	}

	usl30, ok := UserUnSuit30Map.Load(user_id)
	//if !ok {
	//	//sendJsonResponse(ctx, Err, "recommend UserUnSuit30Map is empty. uid: %s", user_id)
	//	//return
	//}

	usl10, ok := UserUnSuit10Map.Load(user_id)
	if !ok {
		sendJsonResponse(ctx, Err, "recommend UserUnSuit10Map is empty. uid: %s", user_id)
		return
	}

	suit = sl.(*ListResult)
	unsuit30 = usl30.(*ListResult)
	unsuit10 = usl10.(*ListResult)

	var returnCount = ReturnCount
	res := new(RecommendResponse)
	suitLen := len(suit.List[(pageNum-1)*6:])
	if suitLen < 6 {
		res.ReturnList = append(res.ReturnList, suit.List[(pageNum-1)*6:]...)
	} else {
		res.ReturnList = append(res.ReturnList, suit.List[(pageNum-1)*6:(pageNum-1)*6+6]...)
	}
	returnCount -= suitLen
	fmt.Println("resultsuit:")
	for i := range res.ReturnList {
		fmt.Println(res.ReturnList[i].BookName)
	}
	var unsuit30Len int
	if unsuit30.List != nil {
		unsuit30Len := len(unsuit30.List[(pageNum-1)*3:])
		if unsuit30Len < 3 {
			res.ReturnList = append(res.ReturnList, unsuit30.List[(pageNum-1)*3:]...)
		} else {
			res.ReturnList = append(res.ReturnList, unsuit30.List[(pageNum-1)*3:(pageNum-1)*3+3]...)
		}
	}
	fmt.Println("result30:")
	for i := range res.ReturnList {
		fmt.Println(res.ReturnList[i].BookName)
	}
	returnCount -= unsuit30Len

	res.ReturnList = append(res.ReturnList, unsuit10.List[(pageNum-1)*pageNum:]...)
	fmt.Println("result10:")
	for i := range res.ReturnList {
		fmt.Println(res.ReturnList[i].BookName)
	}
	//凑够10个
	res.ReturnList = res.ReturnList[:10]
	fmt.Println("result:")
	for i := range res.ReturnList {
		fmt.Println(res.ReturnList[i].BookName)
	}
	response, err := jsoniter.MarshalToString(res)
	if err != nil {
		sendJsonResponse(ctx, Err, "recommend MarshalToString err: %s", err.Error())
		return
	}
	sendJsonResponse(ctx, OK, "%s", response)
}
