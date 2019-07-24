package router

import (
	"net/http"

	"go_demo/handler/check"
	"go_demo/handler/storage"
	"go_demo/handler/user"
	"go_demo/router/middleware"
	"github.com/gin-gonic/gin"
)

func Load(ctx *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	ctx.Use(gin.Recovery())
	ctx.Use(middleware.NoCache)
	ctx.Use(middleware.Options)
	ctx.Use(mw...)

	ctx.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	ctx.POST("/login", user.Login)

	userRoute := ctx.Group("/v1/user")
	userRoute.Use(middleware.AuthMiddleware())
	{
		userRoute.POST("", user.Create)
		userRoute.DELETE("/:id", user.Delete)
		userRoute.PUT("/:id", user.Update)
		userRoute.GET("", user.List)
		userRoute.GET("/:username", user.Get)
	}

	checkRoute := ctx.Group("/check")
	{
		checkRoute.GET("/health", check.HealthCheck)
		checkRoute.GET("/disk", check.DiskCheck)
		checkRoute.GET("/cpu", check.CPUCheck)
		checkRoute.GET("/ram", check.RAMCheck)
	}

	storageRoute := ctx.Group("/v1/file")
	{
		storageRoute.GET("/download", storage.Download)
	}

	return ctx
}
