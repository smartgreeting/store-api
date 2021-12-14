/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:17:07
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-14 22:12:42
 * @Email: 17719495105@163.com
 */
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/store-api/service"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userService := service.NewUserService()

	v1Group := r.Group("v1")
	v1Group.GET("user/getSms", userService.GetSms)
	v1Group.POST("user/register", userService.Register)
	return r
}
