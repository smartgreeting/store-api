package middleware

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/smartgreeting/store-api/utils"
)

func ValidatorUserIdFromToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := strconv.Atoi(ctx.Query("id"))
		userId, _ := ctx.Get("userId")
		if id != int(userId.(int64)) {
			utils.ErrorResponse(ctx, utils.InvalidToken)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
