/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:17:07
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-22 22:38:15
 * @Email: 17719495105@163.com
 */
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/store-api/middleware"
	"github.com/smartgreeting/store-api/service"
	"github.com/smartgreeting/store-api/utils"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 注册参数校验器
	utils.Validator()

	userService := service.NewUserService()
	productService := service.NewProductService()

	v1Group := r.Group("v1")

	{
		v1Group.GET("user/getSms", userService.GetSms)
		v1Group.POST("user/register", userService.Register)
		v1Group.POST("user/login", userService.Login)
		// 注册JWT
		v1Group.Use(middleware.JWT())
		v1Group.GET("user/userInfo", userService.GetUserInfo)
		v1Group.PUT("user/userInfo/:id", userService.UpdateUserInfo)
		// 商品
		v1Group.GET("product/banner", productService.GetBanner)
		v1Group.GET("product/getProduct", productService.GetProduct)
		v1Group.GET("product/getProductList", productService.GetProductList)
	}

	return r
}
