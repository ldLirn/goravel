/**
 * @Author: lirn
 * @Description:
 * @File: common_return
 * @Date: 2022/12/28 15:09
 */
package common

const (
	//Authorization
	AuthorizationMissCode      = 20001
	AuthorizationMissMsg       = "Authorization缺失"
	AuthorizationParseCode     = 20002
	AuthorizationCreateErrCode = 20003
	AuthorizationExpiredCode   = 20004
	AuthorizationExpiredMsg    = "Token已经过期"
	//user
	UserNameOrPassErrCode = 20011
	//Request Bind
	RequestBindErrCode = 20021
	//curd
	AddErrCode       = 20031
	AddErrMsg        = "添加失败"
	AddSuccessMsg    = "添加成功"
	UpdateErrCode    = 20032
	UpdateErrMsg     = "修改失败"
	UpdateSuccessMsg = "修改成功"
	DeleteErrCode    = 20033
	DeleteErrMsg     = "删除失败"
	DeleteSuccessMsg = "删除成功"
	FindOneErrCode   = 20034
	//form
	IdMissCode = 20041
	IdMissMsg  = "id不能为空"
	//Parse
	ParseUintErrCode = 20051

	//权限相关
	AccessDeniedCode   = 20061
	AccessDeniedErrMsg = "没有权限进行此操作"
)
