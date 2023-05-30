package controller

import (
	"canvine/logic"
	"canvine/model"
	"canvine/dao/mysql"
	"fmt"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func SignupHandler(c *gin.Context) {
	var paramsSignup model.ParamsSignup
	if err := c.ShouldBindJSON(&paramsSignup); err != nil {
		validateErr, isValidateErr := err.(validator.ValidationErrors)
		fmt.Println("validateErr:", validateErr)
		fmt.Println("isValidateErr:", isValidateErr)
		if isValidateErr {
			ResponseErrorWithMsg(c, CodeInvalidParam, validateErr)
		} else {
			ResponseError(c, CodeInvalidParam)
		}
		return
	}

	fmt.Println("=====")

	if err := logic.Signup(&paramsSignup); err != nil {
		fmt.Println("err happed in signup:", err)
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		} else {
			ResponseError(c, CodeServerBusy)
		}
		return
	}

	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	logic.Login()
}

func LogoutHandler(c *gin.Context) {

}
