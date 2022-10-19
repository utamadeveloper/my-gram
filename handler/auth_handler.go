package handler

import (
	"my-gram/config"
	"my-gram/helpers"
	"my-gram/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthLoginReq struct {
	Email    string
	Password string
}

func AuthLogin(ctx *gin.Context) {
	var (
		payload AuthLoginReq
		user    model.User
	)

	err := ctx.ShouldBindJSON(&payload)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Invalid payload",
			"error":   err,
		})
		return
	}

	userByEmail := config.Db.Debug().Where("username=? OR email=?", payload.Email, payload.Email).First(&user)

	if userByEmail.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find user",
			"error":   userByEmail.Error.Error(),
		})
		return
	}

	errHashChecked := helpers.CheckPasswordHash([]byte(user.Password), []byte(payload.Password))

	if errHashChecked != true {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Password is wrong",
		})
		return
	}

	jwt := helpers.GenerateToken(user.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success auth login",
		"data": gin.H{
			"access_token": jwt,
		},
	})
}
