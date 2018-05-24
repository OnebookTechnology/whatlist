package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LoadRouter(router *gin.Engine) {
	router.GET("/whoami", func(context *gin.Context) {
		context.String(http.StatusOK, "I am %s. Time,%s", server.ServerName, time.Now().String())
	})

	router.GET("/recommend", recommend)
	router.GET("/sign", Sign)
	router.GET("/update", UpdateUserData)
	router.GET("/gettags", GetAllTags)
	router.GET("/gettagsnum", GetAllTagsNumber)
	router.GET("/bookdetail", GetBookDetail)

	// 获得所有出版社
	router.GET("/presses", GetPresses)
	// 获得出版社推荐书单
	router.GET("/pressrecommendlists", GetPressRecommendLists)

	// 添加喜欢图书
	router.GET("/addinterestedbook", AddInterestedBook)
	// 删除喜欢图书
	router.GET("/deleteinterestedbook", DeleteInterestedBook)
	// 列出喜欢图书
	router.GET("/interestedbooks", InterestedBooks)
	// 最新书单路由
	router.GET("/latestlists", LatestLists)
	// 最热书单路由
	router.GET("/heatlists", HeatLists)
	// 推荐书单路由
	router.GET("/recommendlists", RecommendLists)
	// 书单详细信息路由
	router.GET("/listdetail", ListDetail)
	// 大咖推荐路由
	router.GET("/bigmanrecommend", BigManRecommend)
	// 每日书单路由
	router.GET("/everydayrecommend", EveryDayRecommend)
	// 轮播图
	router.GET("/carousel", Carousel)
}
