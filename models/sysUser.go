package models

import (
	"love-recycling-go/utils"
	"strconv"
	"time"
)

type SysUser struct {
	UserId     uint      `json:"user_id" gorm:"primary_key;auto_increment"`
	UserName   string    `json:"user_name" binding:"required"`
	NickName   string    `json:"nick_name"`
	Phone      int       `json:"phone"`
	Avatar     string    `json:"avatar"`
	Password   string    `json:"password" binding:"required"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"update_time" gorm:"autoUpdateTime"`
	Remark     string    `json:"remark"`
}

func (u SysUser) GetMessages() utils.ValidatorMessages {
	return utils.ValidatorMessages{
		"UserName.required": "用户名称不能为空",
		"Phone.required":    "手机号码不能为空",
		"Password.required": "用户密码不能为空",
	}
}

func (u SysUser) GetUid() string {
	return strconv.Itoa(int(u.UserId))
}
