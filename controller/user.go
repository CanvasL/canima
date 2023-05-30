package controller

import (
	"github.com/gin-gonic/gin"
	"canvine/logic"
)

func SignupHandler(c *gin.Context) {
	logic.Signup()
}

func LoginHandler(c *gin.Context) {
	logic.Login()
}

func LogoutHandler(c *gin.Context) {

}