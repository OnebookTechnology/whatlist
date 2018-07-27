package server

import (
	"fmt"
	"github.com/OnebookTechnology/whatlist/server/models"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"net/http"
)

type JsonResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Page struct {
	PageNum   int `json:"page_num,omitempty" form:"page_num"`
	PageCount int `json:"page_count,omitempty" form:"page_count"`
}

func sendJsonResponse(ctx *gin.Context, code int, format string, values ...interface{}) {
	retMsg := fmt.Sprintf(format, values...)
	logger.Info("[", ctx.Request.URL, "]", "code:", code, "response:", retMsg, "Remote:", ctx.Request.RemoteAddr)
	ctx.JSON(http.StatusOK, JsonResponse{
		Code:    code,
		Message: retMsg,
	})
}

func Options(ctx *gin.Context) {
	crossDomain(ctx)
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Uri     string      `json:"uri"`
	Data    interface{} `json:"data,omitempty"`
}

//注册返回数据结构
type ResData struct {
	Biggies     []*models.Biggie      `json:"biggies,omitempty"`
	Biggie      *models.Biggie        `json:"biggie,omitempty"`
	BiggieLists []*models.BiggieList  `json:"biggie_lists,omitempty"`
	BiggieList  *models.BiggieList    `json:"biggie_list,omitempty"`
	BiggieBooks []*models.BiggieBooks `json:"biggie_books,omitempty"`
	BiggieBook  *models.BiggieBooks   `json:"biggie_book,omitempty"`
	TotalCount  int                   `json:"total_count,omitempty"`
	IsPayed     bool                  `json:"is_payed,omitempty"`
	Name        string                `json:"name,omitempty"`
}

func sendFailedResponse(ctx *gin.Context, code int, v ...interface{}) {
	msg := resFormat(v...)
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Uri:     ctx.Request.RequestURI,
		Message: msg,
	})
	logger.Error("[", ctx.Request.RequestURI, "]", "ErrCode:", code, "response:", msg)

}

func sendSuccessResponse(ctx *gin.Context, data *ResData) {
	ctx.JSON(http.StatusOK, Response{
		Code:    OK,
		Uri:     ctx.Request.RequestURI,
		Message: "ok",
		Data:    data,
	})
	s, _ := jsoniter.MarshalToString(data)
	logger.Info("[", ctx.Request.RequestURI, "]", "response:", s)
}

func sendSuccessResponseWithMessage(ctx *gin.Context, msg string, data *ResData) {
	ctx.JSON(http.StatusOK, Response{
		Code:    OK,
		Uri:     ctx.Request.RequestURI,
		Message: msg,
		Data:    data,
	})
	s, _ := jsoniter.MarshalToString(data)
	logger.Info("[", ctx.Request.RequestURI, "]", "response:", s)
}

func resFormat(v ...interface{}) string {
	formatStr := ""
	for i := 0; i < len(v); i++ {
		formatStr += "%v "
	}
	formatStr += "\n"
	return fmt.Sprintf(formatStr, v...)
}

func crossDomain(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "SESSION,Authorization, Origin, No-Cache, X_Requested_With, X-Requested-With, Content-Range, X_FILENAME, If-Modified-Since, Pragma, Last-Modified, Cache-Control, Expires, Content-Type, X-E4M-With")
	ctx.Header("Custom-Header", "SESSION")
	ctx.Header("Access-Control-Expose-Headers", "SESSION")
}
