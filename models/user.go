/*
 * @Author: lihuan
 * @Date: 2021-12-14 22:04:53
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-22 21:55:40
 * @Email: 17719495105@163.com
 */
package models

import (
	"github.com/smartgreeting/store-rpc/user/user"
)

type Register struct {
	Phone    string `json:"phone" binding:"required,len=11" label:"手机号"`
	Password string `json:"password" binding:"required,max=60,min=6" label:"密码"`
	SmsCode  string `json:"smsCode" binding:"required,max=6,min=4" label:"验证码"`
}
type Login struct {
	Phone    string `json:"phone" binding:"required,len=11" label:"手机号"`
	Password string `json:"password" binding:"required,max=60,min=6" label:"密码"`
}

type UserInfo struct {
	ID        int64  `json:"id"`
	Password  string `json:"-"`
	Username  string `json:"username"`
	Avatar    string `json:"avatar"`
	Gender    int32  `json:"gender"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Hobbies   string `json:"hobbies"`
	CreatedAt int32  `json:"createdAt"`
	UpdatedAt int32  `json:"updatedAt"`
}

func UserMapUserInfo(u *user.UserReply) *UserInfo {
	return &UserInfo{
		ID:        u.Id,
		Username:  u.Username,
		Avatar:    u.Avatar,
		Gender:    u.Gender,
		Phone:     u.Phone,
		Email:     u.Email,
		Address:   u.Address,
		Hobbies:   u.Hobbies,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
