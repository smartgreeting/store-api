/*
 * @Author: lihuan
 * @Date: 2021-12-19 18:12:36
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-19 18:18:34
 * @Email: 17719495105@163.com
 */
package service

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/store-api/rpc"
	"github.com/smartgreeting/store-api/utils"
	"github.com/smartgreeting/store-rpc/product/product"
)

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (p *ProductService) GetBanner(ctx *gin.Context) {
	res, err := rpc.NewProductRpc().GetBanner(context.TODO(), &product.GetBannerReq{})
	if err != nil {
		utils.ErrorReponse(ctx, err)
		return
	}
	utils.SuccessResponse(ctx, gin.H{
		"list": res.List,
	})
}
