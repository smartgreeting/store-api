/*
 * @Author: lihuan
 * @Date: 2021-12-19 18:12:36
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-23 21:58:43
 * @Email: 17719495105@163.com
 */
package service

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	utils.SuccessResponse(ctx, models.ProductMapProduct(res))
}

func (p *ProductService) GetProductList(ctx *gin.Context) {
	res, err := rpc.NewProductRpc().GetProductList(context.TODO(), &product.GetProductListReq{})
	if err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}
	list := make([]*models.Product, 0)
	for _, v := range res.List {
		list = append(list, models.ProductMapProduct(v))
	}
	utils.SuccessResponse(ctx, gin.H{
		"list": list,
	})
}

func (p *ProductService) InrementProduct(ctx *gin.Context) {

	var req models.Product

	err := ctx.ShouldBindWith(&req, binding.JSON)

	if err != nil {
		utils.ErrorResponse(ctx, utils.ParamsParseError)
		return
	}
	_, err = rpc.NewProductRpc().InrementProduct(context.TODO(), &product.ProductReq{
		Id:        req.ID,
		Url:       req.Url,
		Des:       req.Des,
		Name:      req.Name,
		ShortName: req.ShortName,
		Type:      req.Type,
		Price:     req.Price,
		Inventory: req.Inventory,
		Discount:  req.Discount,
	})
	if err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}
	utils.SuccessResponse(ctx, nil)
}

func (p *ProductService) UpdateProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var req models.Product

	err := ctx.ShouldBindWith(&req, binding.JSON)

	if err != nil {
		utils.ErrorResponse(ctx, utils.ParamsParseError)
		return
	}
	_, err = rpc.NewProductRpc().UpdateProduct(context.TODO(), &product.ProductReq{
		Id:        int64(id),
		DetailId:  req.DetailId,
		CommentId: req.CommentId,
		Url:       req.Url,
		Des:       req.Des,
		Name:      req.Name,
		ShortName: req.ShortName,
		Type:      req.Type,
		Price:     req.Price,
		Sale:      req.Sale,
		Inventory: req.Inventory,
		Score:     req.Score,
		Discount:  req.Discount,
	})
	if err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}
	utils.SuccessResponse(ctx, nil)
}

func (p *ProductService) DeleteProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	_, err := rpc.NewProductRpc().DeleteProduct(context.TODO(), &product.DeleteProductReq{
		Id: int64(id),
	})
	if err != nil {
		utils.ErrorResponse(ctx, err)
		return
	}
	utils.SuccessResponse(ctx, nil)
}
