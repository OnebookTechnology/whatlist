package server

import (
	"database/sql"
	"fmt"
	"github.com/OnebookTechnology/whatlist/server/models"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"strconv"
	"strings"
	"time"
)

type VerifyVCodeResponse struct {
	UserInfo  *models.User `json:"user_info"`
	IsNewUser bool         `json:"is_new_user"`
	UserRank  int64        `json:"user_rank"`
}

func Sign(ctx *gin.Context) {
	crossDomain(ctx)
	wxCode := ctx.Query("wx_code")
	wxInfo, err := GetWeChatInfo(wxCode)
	if err != nil {
		sendJsonResponse(ctx, Err, "login error when GetWeChatInfo code: %s err: %s", wxCode, err.Error())
		return
	}
	var user *models.User
	var isNewUser = false
	user, err = server.DB.FindUser(wxInfo.OpenId)
	if err != nil {
		//if NO user found, register user
		if err == sql.ErrNoRows {
			user = &models.User{
				UserId: wxInfo.OpenId,
			}
			isNewUser = true
			user, err = registerUser(user)
			if err != nil {
				sendJsonResponse(ctx, Err, "db error when RegisterUser. err: %s", err.Error())
				return
			}
		} else {
			sendJsonResponse(ctx, Err, "db error when FindUser. err: %s", err.Error())
			return
		}
	}
	if user.Hobby == nil {
		isNewUser = true
	}
	res := &VerifyVCodeResponse{
		UserInfo:  user,
		IsNewUser: isNewUser,
		UserRank:  user.RegisterRank,
	}
	resStr, _ := jsoniter.MarshalToString(res)
	sendJsonResponse(ctx, OK, resStr)
	return
}

func registerUser(user *models.User) (*models.User, error) {
	now := time.Now().Format("2006-01-02 15:04:05")
	logger.Info("New User! userId:", user.UserId, "time:", now)

	lastId, err := server.DB.RegisterUser(user)
	if err != nil {
		return nil, err
	}

	user.RegisterRank = lastId

	return user, nil
}

type UpdateRequest struct {
	Code        string        `json:"code"`
	PhoneNumber uint64        `json:"phone_number"`
	CreateTime  time.Duration `json:"create_time"`
	UserInfo    string        `json:"user_info"`
}

//Field1              int     `json:"field_1,omitempty"` //年龄id
//Field2              int     `json:"field_2,omitempty"` //性别
//Field3              int     `json:"field_3,omitempty"` //婚姻状况id
//Field4              int     `json:"field_4,omitempty"` //教育程度
//Field5              int     `json:"field_5,omitempty"` //收入id
//Field6              int     `json:"field_6,omitempty"` //工作行业id
//Field7              float64 `json:"field_7,omitempty"` //身高体重比例
func UpdateUserData(ctx *gin.Context) {
	crossDomain(ctx)
	userId := ctx.Query("user_id")
	hobbies := ctx.Query("hobbies")
	f1 := ctx.Query("field1")
	f2 := ctx.Query("field2")
	f3 := ctx.Query("field3")
	f4 := ctx.Query("field4")
	f5 := ctx.Query("field5")
	f6 := ctx.Query("field6")
	f7 := ctx.Query("field7")
	field1, _ := strconv.Atoi(f1)
	field2, _ := strconv.Atoi(f2)
	field3, _ := strconv.Atoi(f3)
	field4, _ := strconv.Atoi(f4)
	field5, _ := strconv.Atoi(f5)
	field6, _ := strconv.Atoi(f6)
	field7, _ := strconv.ParseFloat(fmt.Sprintf("%s", f7), 64)
	user := &models.User{
		UserId: userId,
		Field1: field1,
		Field2: field2,
		Field3: field3,
		Field4: field4,
		Field5: field5,
		Field6: field6,
		Field7: field7,
	}
	hobbyArray := strings.Split(hobbies, ",")
	if hobbyArray[0] == "" {
		sendJsonResponse(ctx, Err, "hobbies is empty")
		return
	}
	for _, h := range hobbyArray {
		hi, _ := strconv.Atoi(h)
		user.Hobby = append(user.Hobby, hi)
	}

	err := server.DB.UpdateUser(userId, user)
	if err != nil {
		sendJsonResponse(ctx, Err, "db error when RegisterUser. err: %s", err.Error())
		return
	}

	sendJsonResponse(ctx, OK, "ok")
	return
}
