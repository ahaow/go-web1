package router

import (
	"github.com/gin-gonic/gin"
	"go-web1/controllers"
	"go-web1/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRates)
	api.Use(middlewares.AuthMiddleware())
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
		api.POST("/articles", controllers.CreateArticle)
		api.GET("/articles", controllers.GetArticle)
		api.GET("/articles/:id", controllers.GetArticlesById)

		api.POST("articles/:id/like", controllers.LikeArticle)
		api.GET("articles/:id/like", controllers.GetArticleLikes)
	}

	return r
}
