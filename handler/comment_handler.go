package handler

import (
	"fmt"
	"my-gram/config"
	"my-gram/model"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CommentCreate(ctx *gin.Context) {
	var (
		comment model.Comment
	)

	err := ctx.ShouldBindJSON(&comment)

	fmt.Println(err)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Invalid comment",
			"error":   err,
		})
		return
	}

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	if comment.UserID != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    9,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	commentCreated := config.Db.Debug().Create(&comment)

	if commentCreated.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed create comment",
			"error":   commentCreated.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success create comment",
		"data":    comment,
	})
}

func CommentFindAll(ctx *gin.Context) {
	var (
		comments []model.Comment
	)

	photoIdQuery := ctx.Query("photo_id")
	userIdQuery := ctx.Query("user_id")

	userAll := config.Db.Debug()

	if photoIdQuery != "" {
		userAll = userAll.Where("photo_id=?", photoIdQuery)
	}

	if userIdQuery != "" {
		userAll = userAll.Where("user_id=?", userIdQuery)
	}

	userAll = userAll.Find(&comments)

	if userAll.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find all comment",
			"error":   userAll.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success find all comment",
		"data":    comments,
	})
}

func CommentFindOne(ctx *gin.Context) {
	var (
		comment model.Comment
	)

	id := ctx.Param("id")
	userById := config.Db.Debug().Where("id=?", id).First(&comment)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find comment",
			"error":   userById.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success find comment",
		"data":    comment,
	})
}

func CommentUpdate(ctx *gin.Context) {
	var (
		comment model.Comment
	)

	id := ctx.Param("id")
	userById := config.Db.Debug().Where("id=?", id).First(&comment)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find comment",
			"error":   userById.Error.Error(),
		})
		return
	}

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	if comment.UserID != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    9,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	err := ctx.ShouldBindJSON(&comment)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Invalid comment",
			"error":   err,
		})
		return
	}

	config.Db.Debug().Save(&comment)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success update comment",
		"data":    comment,
	})
}

func CommentDelete(ctx *gin.Context) {
	var (
		comment model.Comment
	)

	id := ctx.Param("id")
	userById := config.Db.Debug().Where("id=?", id).First(&comment)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find comment",
			"error":   userById.Error.Error(),
		})
		return
	}

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	if comment.UserID != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    9,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	config.Db.Debug().Delete(&comment)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success delete comment",
		"data":    comment,
	})
}
