package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func LoadRouter(r *gin.Engine) {
	router := r.Group("/whatlist")
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
	// 大咖书单详细信息路由
	router.GET("/listbigmandetail", ListBigManDetail)
	// 大咖推荐路由
	router.GET("/bigmanrecommend", BigManRecommend)
	// 每日书单路由
	router.GET("/everydayrecommend", EveryDayRecommend)
	// 轮播图
	router.GET("/carousel", Carousel)
	//分类书单
	router.GET("/categorylist", CategoryBooks)
	// 发现
	router.GET("/discover/list", GetDiscoverList)
	router.GET("/discover/get", GetDiscoverDetail)
	router.GET("/discover/like/add", AddLikeNum)
	router.GET("/discover/like/sub", SubLikeNum)

	//最近浏览
	router.GET("/record/list", GetBrowseListRecord)
	router.GET("/record/book", GetBrowseBookRecord)

	//支付
	router.POST("/pay/bigman", Pay)
	router.POST("/pay/callback/:token", PayCallback)

	biggieRouter := router.Group("/biggie")
	{
		biggieRouter.GET("/getlatest", GetLatestBiggie)
		biggieRouter.GET("/get", GetBiggie)
		biggieRouter.GET("/list/get", GetBiggieList)
		biggieRouter.GET("/recommend", GetRecommendBiggie)
		biggieRouter.GET("/listbooks", GetBiggieListBooks)
		biggieRouter.GET("/latestlist", GetLatestBiggieList)
		biggieRouter.POST("/collect/add", CollectBiggie)
		biggieRouter.GET("/collect/get", GetCollectBiggie)
		biggieRouter.DELETE("/collect/delete", RemoveBiggie)
	}
}
