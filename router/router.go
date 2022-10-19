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
			authRoute := v1.Group("/auth")
			{
				authRoute.POST("/login", handler.AuthLogin)
			}

			userRoute := v1.Group("/users")
			{
				userRoute.POST("/register", handler.UserRegister)

				userRoute.Use(middleware.Authentication())
				userRoute.Use(middleware.Authorization())
				userRoute.POST("/", handler.UserCreate)
				userRoute.GET("/", handler.UserFindAll)
				userRoute.GET("/:id", handler.UserFindOne)
				userRoute.PUT("/:id", handler.UserUpdate)
				userRoute.DELETE("/:id", handler.UserDelete)
			}

		}
	}

	return r
}
