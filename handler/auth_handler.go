package handler

import (
	"my-gram/config"
	"my-gram/helpers"
	"my-gram/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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

	if !errHashChecked {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Password is wrong",
		})
		return
	}

	accessToken := helpers.GenerateAccessToken(user.ID)
	refreshToken := helpers.GenerateRefreshToken(user.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success auth login",
		"data": gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}

func AuthRefresh(ctx *gin.Context) {
	var (
		user model.User
	)

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]
	userById := config.Db.Debug().Where("id=?", userAuthId).First(&user)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find user",
			"error":   userById.Error.Error(),
		})
		return
	}

	accessToken := helpers.GenerateAccessToken(user.ID)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success auth login",
		"data": gin.H{
			"access_token": accessToken,
		},
	})
}
