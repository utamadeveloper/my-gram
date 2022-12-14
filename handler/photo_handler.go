package handler

import (
	"my-gram/config"
	"my-gram/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func PhotoCreate(ctx *gin.Context) {
	var (
		photo model.Photo
	)

	err := ctx.ShouldBindJSON(&photo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Invalid photo",
			"error":   err,
		})
		return
	}

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	if photo.UserID != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    97,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	photoCreated := config.Db.Debug().Create(&photo)

	if photoCreated.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed create photo",
			"error":   photoCreated.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success create photo",
		"data":    photo,
	})
}

func PhotoFindAll(ctx *gin.Context) {
	var (
		photos []model.Photo
	)

	userIdQuery := ctx.Query("user_id")

	photoAll := config.Db.Debug()
	photoAll = photoAll.Preload("Comments")

	if userIdQuery != "" {
		photoAll = photoAll.Where("user_id=?", userIdQuery)
	}

	photoAll = photoAll.Find(&photos)

	if photoAll.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find all photo",
			"error":   photoAll.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success find all photo",
		"data":    photos,
	})
}

func PhotoFindOne(ctx *gin.Context) {
	var (
		photo model.Photo
	)

	id := ctx.Param("id")
	userById := config.Db.Debug().Where("id=?", id).Preload("Comments").First(&photo)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find photo",
			"error":   userById.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success find photo",
		"data":    photo,
	})
}

func PhotoUpdate(ctx *gin.Context) {
	var (
		photo model.Photo
	)

	id := ctx.Param("id")
	userById := config.Db.Debug().Where("id=?", id).First(&photo)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find photo",
			"error":   userById.Error.Error(),
		})
		return
	}

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	if photo.UserID != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    97,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	err := ctx.ShouldBindJSON(&photo)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Invalid photo",
			"error":   err,
		})
		return
	}

	config.Db.Debug().Save(&photo)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success update photo",
		"data":    photo,
	})
}

func PhotoDelete(ctx *gin.Context) {
	var (
		photo model.Photo
	)

	id := ctx.Param("id")
	userById := config.Db.Debug().Where("id=?", id).First(&photo)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find photo",
			"error":   userById.Error.Error(),
		})
		return
	}

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	if photo.UserID != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    97,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	config.Db.Debug().Delete(&photo)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success delete photo",
		"data":    photo,
	})
}
