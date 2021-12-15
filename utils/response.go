/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:15:52
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-15 19:46:57
 * @Email: 17719495105@163.com
 */
package utils

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

const (
	Success           = 2000
	ParamsParseError  = 2001
	RpcError          = 2002
	InvalidToken      = 2003
	TokenTimeout      = 2004
	TokenParse        = 2005
	ErrorPhoneNotExit = 2006
)

var resMap = map[int]string{
	Success:           "成功",
	ParamsParseError:  "参数格式错误",
	RpcError:          "rpc端返回错误",
	InvalidToken:      "无效的token",
	TokenTimeout:      "token过期",
	TokenParse:        "token解析失败",
	ErrorPhoneNotExit: "手机号码不正确",
}

func ErrorReponse(ctx *gin.Context, r interface{}) {

	switch v := r.(type) {

	case int:
		getJson(ctx, v, GetMsg(v))

	case validator.ValidationErrors:

		getJson(ctx, ParamsParseError, ErrorTranslate(r.(validator.ValidationErrors)))

	case string:

		getJson(ctx, ParamsParseError, r.(string))

	case error:
		str := v.Error()
		if strings.Contains(str, "code = Unknown desc =") {
			string_slice := strings.Split(str, "code = Unknown desc = ")
			str = string_slice[1]
		}
		getJson(ctx, RpcError, str)

	default:
		getJson(ctx, 5000, "服务异常")
	}
}

func getJson(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}

func SuccessResponse(ctx *gin.Context, data interface{}) {

	ctx.JSON(http.StatusOK, gin.H{
		"code": Success,
		"msg":  GetMsg(Success),
		"data": data,
	})

}

func GetMsg(code int) string {

	str, ok := resMap[code]

	if ok {
		return str
	}
	return "未知错误"
}
