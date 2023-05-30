package mysql

import (
	"canvine/utils"
	"canvine/model"
	"fmt"
)

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username=?`
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	} else {
		return nil
	}
}

func InsertUser(user *model.User) (err error) {
	user.Password = utils.EncryptPassword(user.Password)
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	fmt.Println("users:", user)
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}