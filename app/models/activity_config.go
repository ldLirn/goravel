/**
 * @Author: lirn
 * @Description:
 * @File: activity_config
 * @Date: 2022/12/27 14:46
 */
package models

import (
	"github.com/goravel/framework/facades"
	"goravel/app/http/common"
	"goravel/app/http/requests"
)

type Activity struct {
	common.Activity
}

//获取单条数据
func (a *Activity) Get(id uint64) (Activity, error) {
	var ac Activity
	err := facades.Orm.Query().Table("activity_config").Where("id = ?", id).First(&ac)
	if err != nil {
		return Activity{}, err
	}
	return ac, err
}

//分页获取列表
func (a *Activity) GetList(page string, limit string) ([]Activity, error) {
	var ac []Activity
	err := facades.Orm.Query().Table("activity_config").Scopes(common.Paginator(page, limit)).Find(&ac)

	if err != nil {
		return ac, err
	}
	//log.Println(ac[0].EndTime.Date())
	return ac, err
}

//获取总数
func (a *Activity) GetCount() (c int64) {
	var count int64
	err := facades.Orm.Query().Table("activity_config").Count(&count)
	if err != nil {
		return 0
	}
	return count
}

//增加活动配置
func (a *Activity) AddActivity(activity requests.ActivityAddPostRequest) bool {
	err := facades.Orm.Query().Table("activity_config").Create(&activity)
	if err != nil {
		return false
	}
	return true
}

//修改活动配置
func (a *Activity) UpdateActivity(activity common.ActivityUpdate) bool {
	err := facades.Orm.Query().Table("activity_config").Save(&activity)
	if err != nil {
		return false
	}
	return true
}

//删除活动配置
func (a *Activity) DelActivity(id uint64) bool {
	err := facades.Orm.Query().Table("activity_config").Where("id = ?", id).Delete(id)
	if err != nil {
		return false
	}
	return true
}
