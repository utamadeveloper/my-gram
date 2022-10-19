package middleware

import (
	"my-gram/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, err := helpers.VerifyToken(ctx)

		if err != nil {
			ctx.Writer.Header().Set("Content-Type", "application/json")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    97,
				"type":    "UNAUTHENTICATED",
				"message": "Failed verify user",
				"error":   err,
			})
			return

		} else {
			ctx.Set("userData", userData)
			ctx.Next()
		}
	}
}
