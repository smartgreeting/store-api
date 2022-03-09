/*
 * @Author: lihuan
 * @Date: 2021-12-19 17:45:36
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-23 20:29:08
 * @Email: 17719495105@163.com
 */
package rpc

import (
	"context"

	"github.com/smartgreeting/store-rpc/product/product"
	"github.com/smartgreeting/store-rpc/product/productclient"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/zrpc"
)

type ProductRpcInterface interface {
	GetBanner(ctx context.Context, in *product.GetBannerReq) (*product.BannerReply, error)
	GetProduct(ctx context.Context, in *product.GetProductReq) (*product.ProductReply, error)
	GetProductList(ctx context.Context, in *product.GetProductListReq) (*product.ProductListReply, error)
	InrementProduct(ctx context.Context, in *product.ProductReq) (*product.IncrementProductReply, error)
	UpdateProduct(ctx context.Context, in *product.ProductReq) (*product.UpdateProductReply, error)
	DeleteProduct(ctx context.Context, in *product.DeleteProductReq) (*product.DeleteProductReply, error)
}

type ProductRpc struct{}

var productC productclient.Product

func init() {
	client := zrpc.MustNewClient(zrpc.RpcClientConf{
		Etcd: discov.EtcdConf{
			Hosts: []string{"127.0.0.1:2379"},
			Key:   "product.rpc",
		},
	})
	productC = productclient.NewProduct(client)
}

func NewProductRpc() ProductRpcInterface {
	return &ProductRpc{}
}

// 获取轮播图
func (p *ProductRpc) GetBanner(ctx context.Context, in *product.GetBannerReq) (*product.BannerReply, error) {
	res, err := productC.GetBanner(ctx, &product.GetBannerReq{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 获取商品
func (p *ProductRpc) GetProduct(ctx context.Context, in *product.GetProductReq) (*product.ProductReply, error) {
	res, err := productC.GetProduct(ctx, &product.GetProductReq{
		Id: in.Id,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 获取商品列表
func (p *ProductRpc) GetProductList(ctx context.Context, in *product.GetProductListReq) (*product.ProductListReply, error) {
	res, err := productC.GetProductList(ctx, &product.GetProductListReq{})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 新增商品
func (p *ProductRpc) InrementProduct(ctx context.Context, in *product.ProductReq) (*product.IncrementProductReply, error) {
	res, err := productC.IncrementProduct(ctx, &product.ProductReq{
		Id:        in.Id,
		Url:       in.Url,
		Des:       in.Des,
		Name:      in.Name,
		ShortName: in.ShortName,
		Type:      in.Type,
		Price:     in.Price,
		Inventory: in.Inventory,
		Discount:  in.Discount,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 更新产品
func (p *ProductRpc) UpdateProduct(ctx context.Context, in *product.ProductReq) (*product.UpdateProductReply, error) {
	res, err := productC.UpdateProduct(ctx, &product.ProductReq{
		Id:        in.Id,
		CommentId: in.CommentId,
		DetailId:  in.DetailId,
		Url:       in.Url,
		Des:       in.Des,
		Name:      in.Name,
		ShortName: in.ShortName,
		Type:      in.Type,
		Price:     in.Price,
		Sale:      in.Sale,
		Inventory: in.Inventory,
		Score:     in.Score,
		Discount:  in.Discount,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// 删除产品

func (p *ProductRpc) DeleteProduct(ctx context.Context, in *product.DeleteProductReq) (*product.DeleteProductReply, error) {
	res, err := productC.DeleteProduct(ctx, &product.DeleteProductReq{
		Id: in.Id,
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
