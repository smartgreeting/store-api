/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:15:52
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-13 21:38:24
 * @Email: 17719495105@163.com
 */
package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Conf struct {
	Application Application
	Token       Token
	Md5         Md5
}

type Token struct {
	Secret     string
	ExpireTime int `yaml:"expire_time"`
}
type Md5 struct {
	Secret string
}

type Application struct {
	Address string
	Port    int
	Mode    string
}

var Cfg *Conf = nil

func GetConf(path string) (*Conf, error) {
	// 加载文件
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(yamlFile, &Cfg)

	return Cfg, err
}
