/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:17:50
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-15 23:06:24
 * @Email: 17719495105@163.com
 */
package service

import (
	"context"
	"regexp"
	"strconv"

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
	phone := ctx.Query("phone")
	reg := `^1([38][0-9]|14[579]|5[^4]|16[6]|7[1-35-8]|9[189])\d{8}$`
	rgx := regexp.MustCompile(reg)
	if !rgx.MatchString(phone) {
		utils.ErrorReponse(ctx, utils.ErrorPhoneNotExit)
		return
	}

	res, err := rpc.NewUserRpc().GetSms(context.TODO(), &user.GetSmsReq{
		Phone: phone,
	})
	if err != nil {
		utils.ErrorReponse(ctx, err)
	}
	utils.SuccessResponse(ctx, gin.H{
		"smsCode": res.SmsCode,
	})
}

func (u *UserService) Register(ctx *gin.Context) {

	var req models.Register

	err := ctx.ShouldBindWith(&req, binding.JSON)
	// 参数校验
	if err != nil {

		utils.ErrorReponse(ctx, err)
		return
	}

	// 注册
	res, err := rpc.NewUserRpc().Register(context.TODO(), &user.RegisterReq{
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

func (u *UserService) Login(ctx *gin.Context) {
	var req models.Login
	err := ctx.ShouldBindWith(&req, binding.JSON)
	if err != nil {
		utils.ErrorReponse(ctx, err)
		return
	}
	// 登陆
	res, err := rpc.NewUserRpc().Login(context.TODO(), &user.LoginReq{
		Phone:    req.Phone,
		Password: utils.EncodeMd5(req.Password, []byte(utils.Cfg.Md5.Secret)),
	})
	if err != nil {
		utils.ErrorReponse(ctx, err)
		return
	}
	token, err := utils.GenerateToken(res.Id, req.Phone, utils.EncodeMd5(req.Password, []byte(utils.Cfg.Md5.Secret)), []byte(utils.Cfg.Token.Secret), utils.Cfg.Token.ExpireTime)
	if err != nil {
		utils.ErrorReponse(ctx, err)
		return
	}
	utils.SuccessResponse(ctx, gin.H{
		"id":    res.Id,
		"token": token,
	})
}

func (u *UserService) GetUserInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	res, err := rpc.NewUserRpc().GetUserInfo(context.TODO(), &user.GetUserInfoReq{
		Id: uint64(id),
	})
	if err != nil {
		utils.ErrorReponse(ctx, err)
		return
	}
	utils.SuccessResponse(ctx, res)
}
