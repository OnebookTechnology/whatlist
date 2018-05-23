package server

import (
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"strconv"
)

// 获取最新书单API
func LatestLists(c *gin.Context) {
	lists, err := server.DB.GetLatestSixLists()
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
	lists, err := server.DB.GetRecommendSixLists()
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
	list, err := server.DB.GetListDetail(listID)
	if err != nil {
		sendJsonResponse(c, Err, "GetList error in GetListDetail api. error: %s", err.Error())
		return
	}
	rs, err := jsoniter.MarshalToString(list)
	if err != nil {
		sendJsonResponse(c, Err, "MarshToString error in GetListDetail api. error: %s", err.Error())
		return
	}
	sendJsonResponse(c, OK, "%s", rs)
	return
}
