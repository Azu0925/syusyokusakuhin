package model

import (
	"time"
)

type User struct {
	UserId   int64      `xorm:"pk autoincr int(64)" form:"id" json:"id"`
	Profile  string     `xorm:"varchar(200)" json:"profile" form:"profile"`
	Name     string     `xorm:"varchar(10)" json:"name" form:"name"`
	Birthday *time.Time `xorm:"date json:"birthday" form:"birthday"`
}
