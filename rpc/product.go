/*
 * @Author: lihuan
 * @Date: 2021-12-19 17:45:36
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-19 18:11:56
 * @Email: 17719495105@163.com
 */
package rpc

import (
	"context"

	"github.com/smartgreeting/store-rpc/product/product"
	"github.com/smartgreeting/store-rpc/product/productclient"
	"github.com/tal-tech/go-zero/core/discov"
	"github.com/tal-tech/go-zero/zrpc"
)

type ProductRpcInterface interface {
	GetBanner(ctx context.Context, in *product.GetBannerReq) (*product.BannerReply, error)
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
	return &product.BannerReply{
		List: res.List,
	}, nil
}
