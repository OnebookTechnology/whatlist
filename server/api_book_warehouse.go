package server

import (
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"strconv"
)

func LatestLists (c *gin.Context){
	lists, err := server.DB.GetLatestSixList()
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

func ListDetail (c *gin.Context){
	str_listID := c.Query("listID")
	listID, err := strconv.ParseUint(str_listID, 10, 64)
	if err != nil {
		sendJsonResponse(c, Err, "Can not convert listID to uint64. error:%s, listID:%s",
			err.Error(), str_listID)
		return
	}
	list, err := server.DB.GetList(listID)
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