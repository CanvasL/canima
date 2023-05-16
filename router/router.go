package router

import (
	"github.com/gin-gonic/gin"
	"canima/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1.POST("/user", controller.SignupHandler);

	v1.POST("/user/:username", controller.LoginHandler);

	return r;
}