/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:17:50
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-16 22:05:31
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

// 获取短信验证码
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

// 注册
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

// 登陆
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

// 获取用户信息
func (u *UserService) GetUserInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	// 防止A绕过token校验获取B用户信息
	if ok := models.ValidatorUserIdFromToken(id, ctx); !ok {
		return
	}
	res, err := rpc.NewUserRpc().GetUserInfo(context.TODO(), &user.GetUserInfoReq{
		Id: uint64(id),
	})
	if err != nil {
		utils.ErrorReponse(ctx, err)
		return
	}
	utils.SuccessResponse(ctx, models.UserMapUserInfo(res))
}

// 更新用户信息
func (u *UserService) UpdateUserInfo(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if ok := models.ValidatorUserIdFromToken(id, ctx); !ok {
		return
	}
	var req models.UserInfo

	err := ctx.BindJSON(&req)
	if err != nil {
		utils.ErrorReponse(ctx, utils.ParamsParseError)
		return
	}
	_, err = rpc.NewUserRpc().UpdateUserInfo(context.TODO(), &user.UpdateUserInfoReq{
		Id:       int64(id),
		Username: req.Username,
		Password: utils.EncodeMd5(req.Password, []byte(utils.Cfg.Md5.Secret)),
		Avatar:   req.Avatar,
		Gender:   req.Gender,
		Phone:    req.Phone,
		Email:    req.Email,
		Address:  req.Address,
		Hobbies:  req.Hobbies,
	})
	if err != nil {
		utils.ErrorReponse(ctx, err)
		return
	}
	utils.SuccessResponse(ctx, nil)

}
