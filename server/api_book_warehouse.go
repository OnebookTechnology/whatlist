package server

import (
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
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