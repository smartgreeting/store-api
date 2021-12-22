/*
 * @Author: lihuan
 * @Date: 2021-12-19 18:12:36
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-22 22:36:25
 * @Email: 17719495105@163.com
 */
package service

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/store-api/models"
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
		utils.ErrorResponse(ctx, err)
		return
	}
	utils.SuccessResponse(ctx, gin.H{
		"list": res.List,
	})
}

func (p *ProductService) GetProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Query("id"))
	res, err := rpc.NewProductRpc().GetProduct(context.TODO(), &product.GetProductReq{
		Id: int64(id),
	})
	if err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}

	utils.SuccessResponse(ctx, &models.Product{
		ID:        res.Id,
		DetailId:  res.DetailId,
		CommentId: res.CommentId,
		Url:       res.Url,
		Des:       res.Des,
		Name:      res.Name,
		ShortName: res.ShortName,
		Type:      res.Type,
		Price:     res.Price,
		Sale:      res.Sale,
		Inventory: res.Inventory,
		Score:     res.Score,
		Discount:  res.Discount,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	})
}

func (p *ProductService) GetProductList(ctx *gin.Context) {
	res, err := rpc.NewProductRpc().GetProductList(context.TODO(), &product.GetProductListReq{})
	if err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}
	utils.SuccessResponse(ctx, gin.H{
		"list": res.List,
	})
}
