/*
 * @Author: lihuan
 * @Date: 2021-12-22 22:13:11
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-22 22:16:22
 * @Email: 17719495105@163.com
 */
package models

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
