package server

import (
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"strconv"
	"database/sql"
)

func GetPresses(c *gin.Context) {
	crossDomain(c)
	presses, err := server.DB.GetPresses()
	if err != nil {
		sendJsonResponse(c, Err, "GetPresses error in GetPresses api. Error: %s", err.Error())
		return
	}
	rs, err := jsoniter.MarshalToString(presses)
	if err != nil {
		sendJsonResponse(c, Err, "MarshalToString error in GetPresses api. Error: %s", err.Error())
		return
	}
	sendJsonResponse(c, OK, "%s", rs)
	return
}

func GetPressRecommendLists(c *gin.Context) {
	crossDomain(c)
	pressIDStr := c.Query("press_id")
	if pressIDStr == "" {
		sendJsonResponse(c, Err, "%s", "Empty params press_id")
		return
	}
	pressID, err := strconv.ParseUint(pressIDStr, 10, 64)
	if err != nil {
		sendJsonResponse(c, Err, "Can not convert pressID to uint. Error: %s, press_id: %s",
			err.Error(), pressIDStr)
		return
	}
	lists, err := server.DB.GetPressRecommendLists(pressID)
	if err == sql.ErrNoRows {
		sendJsonResponse(c, EmptyErr, "No recommend lists in %s press.", pressID)
		return
	}
	if err != nil {
		sendJsonResponse(c, Err, "GetPressRecommendLists error in GetPressRecommendLists. Error: %s",
			err.Error())
		return
	}
	rs, err := jsoniter.MarshalToString(lists)
	if err != nil {
		sendJsonResponse(c, Err, "MarshalToString error in GetPressRecommendLists. Error: %s",
			err.Error())
		return
	}
	sendJsonResponse(c, OK, "%s", rs)
	return
}