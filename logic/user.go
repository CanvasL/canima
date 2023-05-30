package logic

import (
	"fmt"
	"canvine/model"
	"canvine/dao/mysql"
	"canvine/utils/snowflake"
)

func Signup(p *model.ParamsSignup) (err error) {
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	userID := snowflake.GenerateID()
	fmt.Println("id:", userID)

	user := &model.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	err = mysql.InsertUser(user)
	return err
}

func Login() {
	fmt.Println("logic: Login")
}