package router

import (
	"github.com/gin-gonic/gin"
	"go-web1/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	auth := r.Group("/api/auth")
	{
		auth.GET("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}
	return r
}
