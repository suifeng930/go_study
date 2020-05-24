package model

import "errors"

var (
	ERROR_USER__NOTEXIST = errors.New("用户不存在...")
	ERROR_USER__EXISTS   = errors.New("用户已经存在...")
	ERROR_USER__Pwd      = errors.New("密码不正确...")
)
