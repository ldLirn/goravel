/**
 * @Author: lirn
 * @Description:
 * @File: ActivityAddPostRequest
 * @Date: 2022/12/29 11:10
 */
package requests

import (
	"errors"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/validation"
	"goravel/app/http/common"
)

type ActivityAddPostRequest struct {
	Title              string `form:"Title" json:"Title"`
	Name               string `form:"Name" json:"Name"`
	Startime           string ` json:"startime" form:"startime"`
	Endtime            string ` json:"endtime" form:"endtime"`
	DayStartime        string ` json:"daystartime" form:"daystartime"`
	DayEndtime         string `json:"dayendtime" form:"dayendtime"`
	MaxIp              uint   `form:"MaxIp" json:"MaxIp"`
	MaxOpenid          uint   `form:"maxOpenid" json:"maxOpenid"`
	MaxVote            uint   `form:"maxVote" json:"maxVote"`
	Vote               uint   `form:"Vote" json:"Vote"`
	Table              string `form:"Table" gorm:"_table" json:"Table"`
	Option_Table       string `form:"option_table" json:"option_table"`
	Option_table_field string `form:"option_table_field" json:"option_table_field"`
}

//授权验证
func (r *ActivityAddPostRequest) Authorize(ctx http.Context) error {
	userId := ctx.Value("userId")
	//权限验证(实列代码,拦截方法定义在app-providers-auth_service_provider)
	if common.AuthCheck("check", map[string]any{"userId": userId}, ctx) == false {
		return errors.New(common.AccessDeniedErrMsg)
	}
	return nil
}

//验证规则
func (r *ActivityAddPostRequest) Rules() map[string]string {
	return map[string]string{
		"Title":              "required|max_len:255",
		"Name":               "required|string",
		"startime":           "required|date",
		"endtime":            "required|date",
		"MaxIp":              "required|uint",
		"MaxOpenid":          "required|uint",
		"MaxVote":            "required|uint",
		"Vote":               "required|uint",
		"Table":              "required|alpha_dash",
		"option_table":       "required|alpha_dash",
		"option_table_field": "required|alpha_dash",
	}
}

//自定义错误消息
func (r *ActivityAddPostRequest) Messages() map[string]string {
	return map[string]string{
		"Title.required":                "活动配置title不能为空",
		"Title.max_len":                 "活动配置title最大长度为255",
		"Name.required":                 "活动名称必须填写",
		"startime.required":             "活动开始时间必须填写",
		"startime.date":                 "活动开始时间应该为日期类型",
		"endtime.required":              "活动结束时间必须填写",
		"endtime.date":                  "活动开始时间应该为日期类型",
		"MaxIp.required":                "每个IP每天最多投票数必须填写",
		"MaxIp.uint":                    "每个IP每天最多投票数必须为正整数",
		"MaxOpenid.required":            "每个微信每天最多投票数必须填写",
		"MaxOpenid.uint":                "每个微信每天最多投票数必须为正整数",
		"MaxVote.required":              "每天为同一选手最多投票数必须填写",
		"MaxVote.uint":                  "每天为同一选手最多投票数必须为正整数",
		"Vote.required":                 "一次投票加的票数必须填写",
		"Vote.uint":                     "一次投票加的票数必须为正整数",
		"Table.required":                "列表表名必须填写",
		"Table.alpha_dash":              "列表表名仅包含字母、数字、破折号（ - ）以及下划线（ _ ）",
		"option_table.required":         "pc端投票记录表名必须填写",
		"option_table.alpha_dash":       "pc端投票记录表名仅包含字母、数字、破折号（ - ）以及下划线（ _ ）",
		"option_table_field.required":   "pc端票数字段必须填写",
		"option_table_field.alpha_dash": "pc端票数字段仅包含字母、数字、破折号（ - ）以及下划线（ _ ）",
	}
}

//自定义验证属性
func (r *ActivityAddPostRequest) Attributes() map[string]string {
	return map[string]string{}
}

//准备验证输入，在应用验证规则之前修改或清理请求中的任何数据
func (r *ActivityAddPostRequest) PrepareForValidation(data validation.Data) {

	if MaxIp, exist := data.Get("MaxIp"); exist {
		_ = data.Set("MaxIp", uint(MaxIp.(float64)))
	}
	if MaxOpenid, exist := data.Get("MaxOpenid"); exist {
		_ = data.Set("MaxOpenid", uint(MaxOpenid.(float64)))
	}
	if MaxVote, exist := data.Get("MaxVote"); exist {
		_ = data.Set("MaxVote", uint(MaxVote.(float64)))
	}
	if Vote, exist := data.Get("Vote"); exist {
		_ = data.Set("Vote", uint(Vote.(float64)))
	}

}
