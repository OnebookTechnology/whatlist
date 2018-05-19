package server

import (
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
)

type TagResponse struct {
	Hobby  []string `json:"hobby"`            //喜好id
	Field1 string   `json:"age,omitempty"`    //年龄id
	Field2 string   `json:"sex,omitempty"`    //性别
	Field3 string   `json:"marry,omitempty"`  //婚姻状况id
	Field4 string   `json:"edu,omitempty"`    //教育程度
	Field5 string   `json:"income,omitempty"` //收入id
	Field6 string   `json:"job,omitempty"`    //工作行业id
	Field7 string   `json:"ratio,omitempty"`  //身高体重比例
}

func GetAllTags(ctx *gin.Context) {
	crossDomain(ctx)
	userId := ctx.Query("user_id")
	if userId == "" {
		sendJsonResponse(ctx, Err, "GetAllTags needs user_id. ")
		return
	}
	user, err := server.DB.FindUser(userId)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when FindUser. err: %s", err.Error())
		return
	}
	var res = new(TagResponse)
	for i := range user.Hobby {
		res.Hobby = append(res.Hobby, CategoryMap[user.Hobby[i]])
	}
	res.Field1 = AgeMap[user.Field1]
	res.Field2 = SexMap[user.Field2]
	res.Field3 = MarriageMap[user.Field3]
	res.Field4 = EduMap[user.Field4]
	res.Field5 = IncomeMap[user.Field5]
	res.Field6 = WorkMap[user.Field6]
	//res.Field7 = WeightMap[user.Field7]

	resp, _ := jsoniter.MarshalToString(res)
	sendJsonResponse(ctx, OK, "%s", resp)
}
