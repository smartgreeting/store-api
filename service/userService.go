/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:17:50
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-13 21:42:05
 * @Email: 17719495105@163.com
 */
package service

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/store-api/rpc"
	"github.com/smartgreeting/store-api/utils"
	"github.com/smartgreeting/store-rpc/user/user"
)

type userService struct{}

func NewUserService() *userService {
	return &userService{}
}
func (u *userService) GetSms(ctx *gin.Context) {
	res, err := rpc.NewUserRpc().GetSms(context.TODO(), &user.GetSmsReq{
		Phone: "17719495105",
	})
	if err != nil {
		utils.ErrorReponse(ctx, err)
	}
	utils.SuccessResponse(ctx, gin.H{
		"smsCode": res.SmsCode,
	})
}
