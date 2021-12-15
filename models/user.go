/*
 * @Author: lihuan
 * @Date: 2021-12-14 22:04:53
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-15 20:46:19
 * @Email: 17719495105@163.com
 */
package models

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
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Avatar    string `json:"avatar"`
	Gender    int32  `json:"gender"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Address   string `json:"address"`
	Hobbies   string `json:"hobbies"`
	Deleted   int    `json:"deleted"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
}
