/**
 * @Author: lirn
 * @Description:
 * @File: struct
 * @Date: 2022/12/15 16:19
 */
package common

import (
	"goravel/app/utils"
	"time"
)

//登录
type User struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

//返回用户信息
type ReturnUser struct {
	Id        uint
	Name      string
	Password  string
	UpdatedAt time.Time
	CreatedAt time.Time
	DeletedAt time.Time
}

//活动配置信息  (数据库绑定，新增使用)
type Activity struct {
	Id                 uint            `gorm:"primaryKey"`
	Title              string          `form:"title" binding:"required"`
	Name               string          `form:"name" binding:"required"`
	Startime           utils.LocalTime ` json:"startime" form:"startime" binding:"required"`
	Endtime            utils.LocalTime ` json:"endtime" form:"endtime" binding:"required"`
	DayStartime        string          ` json:"daystartime" form:"daystartime" binding:"required"`
	DayEndtime         string          `json:"dayendtime" form:"dayendtime" binding:"required"`
	MaxIp              uint            `form:"max_ip" binding:"required"`
	MaxOpenid          uint            `form:"max_openid" binding:"required"`
	MaxVote            uint            `form:"max_vote" binding:"required"`
	Vote               uint            `form:"vote" binding:"required"`
	Table              string          `form:"table" gorm:"_table" binding:"required"`
	OptionTable        string          `form:"option_table"`
	Option_table_field string          `form:"option_table_field"`
	Status             int
}

//活动配置信息 （修改操作使用）
type ActivityUpdate struct {
	Id                 uint            `form:"Id" json:"Id" binding:"required"`
	Title              string          `form:"Title" json:"Title"`
	Name               string          `form:"Name" json:"Name"`
	Startime           utils.LocalTime ` json:"startime" form:"startime"`
	Endtime            utils.LocalTime ` json:"endtime" form:"endtime"`
	DayStartime        string          ` json:"daystartime" form:"daystartime"`
	DayEndtime         string          `json:"dayendtime" form:"dayendtime"`
	MaxIp              uint            `form:"MaxIp" json:"MaxIp"`
	MaxOpenid          uint            `form:"MaxOpenid" json:"MaxOpenid"`
	MaxVote            uint            `form:"MaxVote" json:"MaxVote"`
	Vote               uint            `form:"Vote" json:"Vote"`
	Table              string          `form:"Table" json:"Table"`
	OptionTable        string          `form:"OptionTable" json:"OptionTable"`
	Option_table_field string          `form:"Option_table_field" json:"Option_table_field"`
}
