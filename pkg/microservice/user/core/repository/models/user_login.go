package models

import (
	"github.com/koderover/zadig/pkg/microservice/user/core"
)

type UserLogin struct {
	core.Model
	Uid           string `json:"uid"`
	Password      string `json:"password"`
	LastLoginTime int64  `json:"last_login_time"`
}

// TableName sets the insert table name for this struct type
func (UserLogin) TableName() string {
	return "user_login"
}
