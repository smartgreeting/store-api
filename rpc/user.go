/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:17:43
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-22 21:52:39
 * @Email: 17719495105@163.com
 */
package rpc

import (
	"context"

	"github.com/smartgreeting/store-rpc/user/user"
	"github.com/smartgreeting/store-rpc/user/userclient"
	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc"
)

type UserRpcInterface interface {
	GetSms(ctx context.Context, in *user.GetSmsReq) (*user.UserReply, error)
	Register(ctx context.Context, in *user.RegisterReq) (*user.UserReply, error)
	Login(ctx context.Context, in *user.LoginReq) (*user.UserReply, error)
	GetUserInfo(ctx context.Context, in *user.GetUserInfoReq) (*user.UserReply, error)
	UpdateUserInfo(ctx context.Context, in *user.UpdateUserInfoReq) (*user.UserReply, error)
}
type userRpc struct {
}

var userC userclient.User

func init() {
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "user.rpc"},
	})
	userC = userclient.NewUser(client)
}

func NewUserRpc() UserRpcInterface {
	return &userRpc{}
}

// 获取验证码
func (u *userRpc) GetSms(ctx context.Context, in *user.GetSmsReq) (*user.UserReply, error) {
	res, err := userC.GetSms(ctx, &userclient.GetSmsReq{
		Phone: in.Phone,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 注册
func (u *userRpc) Register(ctx context.Context, in *user.RegisterReq) (*user.UserReply, error) {

	res, err := userC.Register(ctx, &userclient.RegisterReq{
		Phone:    in.Phone,
		Password: in.Password,
		SmsCode:  in.SmsCode,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

// 登陆
func (u *userRpc) Login(ctx context.Context, in *user.LoginReq) (*user.UserReply, error) {
	res, err := userC.Login(ctx, &userclient.LoginReq{
		Phone:    in.Phone,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

//获取用户信息
func (u *userRpc) GetUserInfo(ctx context.Context, in *user.GetUserInfoReq) (*user.UserReply, error) {

	res, err := userC.GetUserInfo(ctx, &userclient.GetUserInfoReq{
		Id: in.Id,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 更新用户信息

func (u *userRpc) UpdateUserInfo(ctx context.Context, in *user.UpdateUserInfoReq) (*user.UserReply, error) {
	_, err := userC.UpdateUserInfo(ctx, &user.UpdateUserInfoReq{
		Id:       in.Id,
		Username: in.Username,
		Password: in.Password,
		Avatar:   in.Avatar,
		Gender:   in.Gender,
		Phone:    in.Phone,
		Email:    in.Email,
		Address:  in.Address,
		Hobbies:  in.Hobbies,
	})
	if err != nil {
		return nil, err
	}
	return &user.UserReply{}, nil
}
