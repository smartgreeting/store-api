/*
 * @Author: lihuan
 * @Date: 2021-12-13 20:17:33
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-15 20:43:36
 * @Email: 17719495105@163.com
 */
package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/store-api/utils"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var code int

		code = utils.Success

		Authorization := ctx.GetHeader("Authorization")

		token := strings.Split(Authorization, " ")

		if Authorization == "" {
			code = utils.InvalidToken
		} else {
			claims, err := utils.ParseToken(token[1], []byte(utils.Cfg.Token.Secret))
			if err != nil {

				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = utils.TokenTimeout
				default:
					code = utils.TokenParse
				}
			} else {
				id := claims.ID
				ctx.Set("userId", id)
			}
		}

		if code != utils.Success {
			utils.ErrorReponse(ctx, code)

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
