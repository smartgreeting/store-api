/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:14:22
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-13 22:20:16
 * @Email: 17719495105@163.com
 */
package main

import (
	"fmt"

	"github.com/smartgreeting/store-api/handler"
	"github.com/smartgreeting/store-api/utils"
)

func main() {

	cfg, err := utils.GetConf("conf/conf.yaml")
	if err != nil {
		panic(err)
	}
	r := handler.SetupRouter()
	r.Run(fmt.Sprintf("%s:%d", cfg.Application.Address, cfg.Application.Port))
}
