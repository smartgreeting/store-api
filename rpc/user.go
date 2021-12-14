/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:17:43
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-14 22:01:39
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
}
type userRpc struct {
}

var rpc userclient.User

func init() {
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "user.rpc"},
	})
	rpc = userclient.NewUser(client)
}

func NewUserRpc() UserRpcInterface {
	return &userRpc{}
}

func (u *userRpc) GetSms(ctx context.Context, in *user.GetSmsReq) (*user.UserReply, error) {
	res, err := rpc.GetSms(ctx, &userclient.GetSmsReq{
		Phone: in.Phone,
	})
	if err != nil {
		return nil, err
	}
	return &user.UserReply{
		SmsCode: res.SmsCode,
	}, nil
}
func (r *userRpc) Register(ctx context.Context, in *user.RegisterReq) (*user.UserReply, error) {

	_, err := rpc.Register(ctx, &userclient.RegisterReq{
		Phone:    in.Phone,
		Password: in.Password,
		SmsCode:  in.SmsCode,
	})
	if err != nil {
		return nil, err
	}

	return &user.UserReply{}, nil
}
