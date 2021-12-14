/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:17:50
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-14 22:13:54
 * @Email: 17719495105@163.com
 */
package service

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/smartgreeting/store-api/models"
	"github.com/smartgreeting/store-api/rpc"
	"github.com/smartgreeting/store-api/utils"
	"github.com/smartgreeting/store-rpc/user/user"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}
func (u *UserService) GetSms(ctx *gin.Context) {
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

func (s *UserService) Register(ctx *gin.Context) {

	var req models.User

	err := ctx.ShouldBindWith(&req, binding.JSON)
	// 参数校验
	if err != nil {

		utils.ErrorReponse(ctx, err)
		return
	}

	// 注册
	res, err := rpc.NewUserRpc().Register(ctx, &user.RegisterReq{
		Phone:    req.Phone,
		Password: utils.EncodeMd5(req.Password, []byte(utils.Cfg.Md5.Secret)),
		SmsCode:  req.SmsCode,
	})
	// 注册失败
	if err != nil {
		utils.ErrorReponse(ctx, err)
		return
	}
	utils.SuccessResponse(ctx, gin.H{
		"id": res.Id,
	})
}
