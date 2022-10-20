package handler

import (
	"my-gram/config"
	"my-gram/model"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserCreate(ctx *gin.Context) {
	var (
		user model.User
	)

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Invalid user",
			"error":   err,
		})
		return
	}

	userCreated := config.Db.Debug().Create(&user)

	if userCreated.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed create user",
			"error":   userCreated.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success create user",
		"data":    user,
	})
}

func UserRegister(ctx *gin.Context) {
	var (
		user model.User
	)

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Invalid user",
			"error":   err,
		})
		return
	}

	userCreated := config.Db.Debug().Create(&user)

	if userCreated.Error != nil {
		userErr := userCreated.Error.Error()

		if strings.Contains(userErr, "duplicate key value violates unique constraint \"idx_email\"") {
			userErr = "Email already used"
		}

		if strings.Contains(userErr, "duplicate key value violates unique constraint \"idx_username\"") {
			userErr = "Username already used"
		}

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed create user",
			"error":   userErr,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success create user",
		"data":    user,
	})
}

func UserFindAll(ctx *gin.Context) {
	var (
		users []model.User
	)

	userAll := config.Db.Debug()
	userAll = userAll.Select("ID", "Username", "Email", "DOB", "Age")
	userAll = userAll.Preload("SocialMedias")
	userAll = userAll.Preload("Comments")
	userAll = userAll.Preload("Photos")
	userAll = userAll.Find(&users)

	if userAll.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find all user",
			"error":   userAll.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success find all user",
		"data":    users,
	})
}

func UserFindOne(ctx *gin.Context) {
	var (
		user model.User
	)

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	id := ctx.Param("id")

	if id != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    97,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	userById := config.Db.Debug()
	userById = userById.Select("ID", "Username", "Email", "DOB", "Age")
	userById = userById.Where("id=?", id)
	userById = userById.Preload("SocialMedias")
	userById = userById.Preload("Comments")
	userById = userById.Preload("Photos")
	userById = userById.First(&user)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find user",
			"error":   userById.Error.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success find user",
		"data":    user,
	})
}

func UserUpdate(ctx *gin.Context) {
	var (
		user model.User
	)

	id := ctx.Param("id")
	userById := config.Db.Debug().Where("id=?", id).First(&user)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find user",
			"error":   userById.Error.Error(),
		})
		return
	}

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Invalid user",
			"error":   err,
		})
		return
	}

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	if user.ID != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    97,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	config.Db.Debug().Save(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success update user",
		"data":    user,
	})
}

func UserDelete(ctx *gin.Context) {
	var (
		user model.User
	)

	id := ctx.Param("id")
	userById := config.Db.Debug().Where("id=?", id).First(&user)

	if userById.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    96,
			"type":    "BAD_REQUEST",
			"message": "Failed find user",
			"error":   userById.Error.Error(),
		})
		return
	}

	userAuth := ctx.MustGet("userData")
	userAuthId := userAuth.(jwt.MapClaims)["id"]

	if user.ID != userAuthId {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    97,
			"type":    "FORBIDDEN",
			"message": "User forbidden",
		})
		return
	}

	config.Db.Debug().Delete(&user)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    00,
		"type":    "SUCCESS",
		"message": "Success delete user",
		"data":    user,
	})
}
