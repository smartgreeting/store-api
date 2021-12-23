/*
 * @Author: lihuan
 * @Date: 2021-12-22 22:13:11
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-23 21:47:42
 * @Email: 17719495105@163.com
 */
package models

import "github.com/smartgreeting/store-rpc/product/product"

type Product struct {
	ID        int64  `json:"id"`
	DetailId  int64  `json:"detailId"`
	CommentId int64  `json:"commentId"`
	Url       string `json:"url"`
	Des       string `json:"des"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	Type      int32  `json:"type"`
	Price     string `json:"price"`
	Sale      int64  `json:"sale"`
	Inventory int64  `json:"inventory"`
	Score     string `json:"score"`
	Discount  string `json:"discount"`
	CreatedAt int32  `json:"createdAt"`
	UpdatedAt int32  `json:"updatedAt"`
}

func ProductMapProduct(res *product.ProductReply) *Product {
	return &Product{
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
	}
}
