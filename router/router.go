package router

import (
	"my-gram/handler"
	"my-gram/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.JSONResponseMiddleware())
	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			// authRoute := v1.Group("/auth")
			// {
			// }

			userRoute := v1.Group("/users")
			{
				userRoute.POST("/login", handler.AuthLogin)
				userRoute.POST("/register", handler.UserRegister)

				userRoute.Use(middleware.Authentication())
				userRoute.Use(middleware.Authorization())
				userRoute.POST("/", handler.UserCreate)
				userRoute.GET("/", handler.UserFindAll)
				userRoute.GET("/:id", handler.UserFindOne)
				userRoute.PUT("/:id", handler.UserUpdate)
				userRoute.DELETE("/:id", handler.UserDelete)
			}

			socialMediaRoute := v1.Group("/socialmedias")
			{
				socialMediaRoute.Use(middleware.Authentication())
				socialMediaRoute.Use(middleware.Authorization())
				socialMediaRoute.POST("/", handler.SocialMediaCreate)
				socialMediaRoute.GET("/", handler.SocialMediaFindAll)
				socialMediaRoute.GET("/:id", handler.SocialMediaFindOne)
				socialMediaRoute.PUT("/:id", handler.SocialMediaUpdate)
				socialMediaRoute.DELETE("/:id", handler.SocialMediaDelete)
			}

			photoRoute := v1.Group("/photos")
			{
				photoRoute.Use(middleware.Authentication())
				photoRoute.Use(middleware.Authorization())
				photoRoute.POST("/", handler.PhotoCreate)
				photoRoute.GET("/", handler.PhotoFindAll)
				photoRoute.GET("/:id", handler.PhotoFindOne)
				photoRoute.PUT("/:id", handler.PhotoUpdate)
				photoRoute.DELETE("/:id", handler.PhotoDelete)
			}

			commentRoute := v1.Group("/comments")
			{
				commentRoute.Use(middleware.Authentication())
				commentRoute.Use(middleware.Authorization())
				commentRoute.POST("/", handler.CommentCreate)
				commentRoute.GET("/", handler.CommentFindAll)
				commentRoute.GET("/:id", handler.CommentFindOne)
				commentRoute.PUT("/:id", handler.CommentUpdate)
				commentRoute.DELETE("/:id", handler.CommentDelete)
			}

		}
	}

	return r
}
