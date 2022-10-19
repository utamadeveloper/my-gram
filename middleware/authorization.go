package middleware

import (
	"my-gram/config"
	"my-gram/helpers"
	"my-gram/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			user model.User
		)

		claim, _ := helpers.VerifyToken(ctx)

		userId := claim.(jwt.MapClaims)["id"]
		userById := config.Db.Debug().Where("id=?", userId).First(&user)

		if userById.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    96,
				"type":    "BAD_REQUEST",
				"message": "Failed find user",
				"error":   userById.Error.Error(),
			})
			return
		}

		ctx.Next()
	}
}
