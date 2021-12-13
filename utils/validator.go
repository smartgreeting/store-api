/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:15:52
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-13 21:38:41
 * @Email: 17719495105@163.com
 */
package utils

import (
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin/binding"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func Validator() {

	uni := ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册翻译器
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
		//注册自定义函数
		// _ = v.RegisterValidation()

		//注册一个函数，获取struct tag里自定义的label作为字段名
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := fld.Tag.Get("label")
			return name
		})

	}
}

func ErrorTranslate(e validator.ValidationErrors) string {
	var errs = e.Translate(trans)

	m := make(map[string]string)

	for key, val := range errs {
		m["msg"] = val
		m["field"] = key
	}

	return m["msg"]
}
