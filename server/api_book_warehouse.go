package server

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"strconv"
)

// 获取最新书单API
func LatestLists(c *gin.Context) {
	index_str := c.Query("index")
	index, err := strconv.ParseUint(index_str, 10, 64)
	lists, err := server.DB.GetLatestSixLists(index)
	if err != nil {
		sendJsonResponse(c, Err, "GetLatestSixList error in LatestLists api. error: %s", err.Error())
		return
	}
	rs, err := jsoniter.MarshalToString(lists)
	if err != nil {
		sendJsonResponse(c, Err, "MarshToString error in LatestLists api. error: %s", err.Error())
		return
	}
	sendJsonResponse(c, OK, "%s", rs)
	return
}

// 获取最热书单API
func HeatLists(c *gin.Context) {
	lists, err := server.DB.GetHeatSixLists()
	if err != nil {
		sendJsonResponse(c, Err, "GetHeatSixLists error in HeatLists api. error: %s", err.Error())
		return
	}
	rs, err := jsoniter.MarshalToString(lists)
	if err != nil {
		sendJsonResponse(c, Err, "MarshToString error in HeatLists api. error: %s", err.Error())
		return
	}
	sendJsonResponse(c, OK, "%s", rs)
	return
}

// 获取推荐书单API
func RecommendLists(c *gin.Context) {
	index_str := c.Query("index")
	index, err := strconv.ParseUint(index_str, 10, 64)
	lists, err := server.DB.GetRecommendSixLists(index)
	if err != nil {
		sendJsonResponse(c, Err, "GetRecommendSixList error in RecommendLists api. error: %s", err.Error())
		return
	}
	rs, err := jsoniter.MarshalToString(lists)
	if err != nil {
		sendJsonResponse(c, Err, "MarshToString error in RecommendLists api. error: %s", err.Error())
		return
	}
	sendJsonResponse(c, OK, "%s", rs)
	return
}

// 获取指定书单详细信息API
func ListDetail(c *gin.Context) {
	str_listID := c.Query("listID")
	listID, err := strconv.ParseUint(str_listID, 10, 64)
	if err != nil {
		sendJsonResponse(c, Err, "Can not convert listID to uint64. error:%s, listID:%s",
			err.Error(), str_listID)
		return
	}
	userId := c.Query("user_id")
	list, err := server.DB.GetListDetail(listID)
	if err != nil {
		sendJsonResponse(c, Err, "GetList error in GetListDetail api. error: %s", err.Error())
		return
	}
	err = server.DB.AddListClickCount(listID)
	if err != nil {
		logger.Error("click count add failed in listID:", listID)
	}

	if userId != "" {
		server.DB.AddListRecord(userId, listID)
		if err != nil {
			logger.Error("AddListRecord failed in listID:", listID, "userId:", userId)
		}
	}

	rs, err := jsoniter.MarshalToString(list)
	if err != nil {
		sendJsonResponse(c, Err, "MarshToString error in GetListDetail api. error: %s", err.Error())
		return
	}
	sendJsonResponse(c, OK, "%s", rs)
	return
}

// 获取大咖书单详细信息API
func ListBigManDetail(c *gin.Context) {
	str_listID := c.Query("listID")
	listID, err := strconv.ParseUint(str_listID, 10, 64)
	if err != nil {
		sendJsonResponse(c, Err, "Can not convert listID to uint64. error:%s, listID:%s",
			err.Error(), str_listID)
		return
	}
	userId := c.Query("user_id")

	list, err := server.DB.GetListDetail(listID)
	if err != nil {
		sendJsonResponse(c, Err, "GetList error in GetListDetail api. error: %s", err.Error())
		return
	}
	list.IsPayed = true

	//查询是否有购买过
	_, err = server.DB.FindListPurchaseRecord(userId, int(listID))
	if err != nil {
		//没买过，不返会书单目录
		if err == sql.ErrNoRows {
			list.IsPayed = false
		} else {
			sendJsonResponse(c, Err, "FindListPurchaseRecord error in ListBigManDetail api. error: %s", err.Error())
			return
		}
	}

	err = server.DB.AddListClickCount(listID)
	if err != nil {
		logger.Error("click count add failed in listID:", listID)
	}

	if userId != "" {
		server.DB.AddListRecord(userId, listID)
		if err != nil {
			logger.Error("AddListRecord failed in listID:", listID, "userId:", userId)
		}
	}

	rs, err := jsoniter.MarshalToString(list)
	if err != nil {
		sendJsonResponse(c, Err, "MarshToString error in GetListDetail api. error: %s", err.Error())
		return
	}
	sendJsonResponse(c, OK, "%s", rs)
	return
}

// 获取大咖推荐
func BigManRecommend(c *gin.Context) {
	lists, err := server.DB.GetBigManRecommendLists()
	if err != nil {
		sendJsonResponse(c, Err, "GetBigManRecommendLists error in BigManRecommend api. error: %s", err.Error())
		return
	}
	rs, err := jsoniter.MarshalToString(lists)
	if err != nil {
		sendJsonResponse(c, Err, "MarshToString error in BigManRecommend api. error: %s", err.Error())
		return
	}
	sendJsonResponse(c, OK, "%s", rs)
	return
}

// 获取每日推荐
func EveryDayRecommend(c *gin.Context) {
	index_str := c.Query("index")
	index, err := strconv.ParseUint(index_str, 10, 64)
	if err != nil {
		sendJsonResponse(c, Err, "Can not convert index to uint64. error:%s, index:%s",
			err.Error(), index_str)
		return
	}
	everyDayRecommend, err := server.DB.GetEveryDayRecommendList(index)
	if err == sql.ErrNoRows {
		// -3
		sendJsonResponse(c, NoResultErr, "%s", "已浏览到最后一个每日推荐")
		return
	}
	if err != nil {
		sendJsonResponse(c, Err, "GetEveryDayRecommendList error in everyDayRecommend api. error: %s",
			err.Error())
		return
	}
	rs, err := jsoniter.MarshalToString(everyDayRecommend)
	if err != nil {
		sendJsonResponse(c, Err, "MarshToString error in EveryDayRecommend api. error: %s", err.Error())
		return
	}
	sendJsonResponse(c, OK, "%s", rs)
	return
}

// 获取轮播图
func Carousel(c *gin.Context) {
	carousels, err := server.DB.GetCarousel()
	if err != nil {
		sendJsonResponse(c, Err, "GetCarousel error in Carousel api. error: %s", err.Error())
		return
	}
	rs, err := jsoniter.MarshalToString(carousels)
	if err != nil {
		sendJsonResponse(c, Err, "MarshToString error in Carousel api. error: %s", err.Error())
		return
	}
	sendJsonResponse(c, OK, "%s", rs)
	return
}
