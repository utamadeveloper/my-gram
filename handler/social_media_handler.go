package handler

import (
	"my-gram/config"
	"my-gram/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func SocialMediaCreate(ctx *gin.Context) {
	var (
		socialMedia model.SocialMedia
	)

	err := ctx.ShouldBindJSON(&socialMedia)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Invalid social media",
			"error":   err,
		})
		return
	}

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	if socialMedia.UserID != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    9,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	socialMediaCreated := config.Db.Debug().Create(&socialMedia)

	if socialMediaCreated.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed create social media",
			"error":   socialMediaCreated.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success create social media",
		"data":    socialMedia,
	})
}

func SocialMediaFindAll(ctx *gin.Context) {
	var (
		socialMedias []model.SocialMedia
	)

	userAll := config.Db.Debug().Find(&socialMedias)

	if userAll.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find all social media",
			"error":   userAll.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success find all social media",
		"data":    socialMedias,
	})
}

func SocialMediaFindOne(ctx *gin.Context) {
	var (
		socialMedia model.SocialMedia
	)

	id := ctx.Param("id")
	userById := config.Db.Debug().Where("id=?", id).First(&socialMedia)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find social media",
			"error":   userById.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success find social media",
		"data":    socialMedia,
	})
}

func SocialMediaUpdate(ctx *gin.Context) {
	var (
		socialMedia model.SocialMedia
	)

	id := ctx.Param("id")
	userById := config.Db.Debug().Where("id=?", id).First(&socialMedia)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find social media",
			"error":   userById.Error.Error(),
		})
		return
	}

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	if socialMedia.UserID != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    9,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	err := ctx.ShouldBindJSON(&socialMedia)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Invalid social media",
			"error":   err,
		})
		return
	}

	config.Db.Debug().Save(&socialMedia)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success update social media",
		"data":    socialMedia,
	})
}

func SocialMediaDelete(ctx *gin.Context) {
	var (
		socialMedia model.SocialMedia
	)

	id := ctx.Param("id")
	userById := config.Db.Debug().Where("id=?", id).First(&socialMedia)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find social media",
			"error":   userById.Error.Error(),
		})
		return
	}

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	if socialMedia.UserID != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    9,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	config.Db.Debug().Delete(&socialMedia)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success delete social media",
		"data":    socialMedia,
	})
}
