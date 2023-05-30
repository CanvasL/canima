package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("User already exists")
	ErrorUserNotExist    = errors.New("User not exsits")
	ErrorInvalidPassword = errors.New("Username or password error")
	ErrorInvalidID       = errors.New("Invalid ID")
)
