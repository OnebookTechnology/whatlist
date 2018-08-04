package server

import ()

//func GetLatestBiggie(ctx *gin.Context) {
//	crossDomain(ctx)
//	pageNumStr := ctx.Query("page_num")
//	pageNum, err := strconv.Atoi(pageNumStr)
//	if err != nil {
//		sendFailedResponse(ctx, Err, "parse page_num error:", err, "page_num:", pageNumStr)
//		return
//	}
//	pageCountStr := ctx.Query("page_count")
//
//}
//
//type CollectReq struct {
//	UserId   string `json:"user_id" form:"user_id"`
//	BiggieId int    `json:"biggie_id" form:"biggie_id"`
//	Page
//}
//
//func CollectBiggie(ctx *gin.Context) {
//	crossDomain(ctx)
//	var req CollectReq
//	if err := ctx.BindJSON(&req); err == nil {
//		c := &models.BiggieCollect{
//			UserId:   req.UserId,
//			BiggieId: req.BiggieId,
//		}
//		err := server.DB.AddCollectBiggie(c)
//		if err != nil {
//			sendFailedResponse(ctx, Err, "AddCollectBiggie err:", err)
//			return
//		}
//
//		sendSuccessResponse(ctx, nil)
//		return
//
//	} else {
//		sendFailedResponse(ctx, Err, "BindJSON err:", err)
//		return
//	}
//}
