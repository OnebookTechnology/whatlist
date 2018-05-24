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

	// 添加喜欢图书
	router.GET("/addInterestedBook", AddInterestedBook)
	// 删除喜欢图书
	router.GET("/deleteInterestedBook", DeleteInterestedBook)
	// 列出喜欢图书
	router.GET("/InterestedBooks", InterestedBooks)
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
