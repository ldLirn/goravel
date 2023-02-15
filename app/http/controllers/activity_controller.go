/**
 * @Author: lirn
 * @Description:
 * @File: activity_controller
 * @Date: 2022/12/27 15:16
 */
package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/http/common"
	"goravel/app/http/requests"
	"goravel/app/models"
	"strconv"
	"time"
)

type ActivityController struct {
	*common.Activity
}

func NewActivityController() *ActivityController {
	return &ActivityController{
		//Inject services
	}
}

func (a *ActivityController) GetActivity(ctx http.Context) {
	id := ctx.Request().Input("id")
	orm_id, err1 := strconv.ParseUint(id, 10, 32)
	if err1 != nil {
		common.JsonReturnError(common.ParseUintErrCode, err1.Error(), ctx)
		return
	}
	var activity models.Activity
	data, err := activity.Get(orm_id)
	if err != nil {
		common.JsonReturnError(common.FindOneErrCode, err.Error(), ctx)
		return
	} else {
		common.JsonReturnSuccess(data, "", ctx)
	}
}

func (a *ActivityController) GetActivityList(ctx http.Context) {
	page := ctx.Request().Query("page", "1")
	limit := ctx.Request().Query("limit", "10")
	var activity models.Activity
	data, _ := activity.GetList(page, limit)
	count := activity.GetCount()
	for i, s := range data {
		s_t := common.StringToUnix(s.Startime.String())
		e_t := common.StringToUnix(s.Endtime.String())
		d_t := time.Now().Unix()
		if d_t >= s_t && d_t <= e_t {
			data[i].Status = 1 //进行中
		} else if d_t < s_t {
			data[i].Status = 0 //未开始
		} else if d_t > e_t {
			data[i].Status = 2 //已结束
		}
	}
	ctx.Response().Success().Json(http.Json{
		"code":  20000,
		"msg":   "success",
		"info":  "",
		"data":  data,
		"count": count,
	})
}

func (a *ActivityController) AddActivity(ctx http.Context) {
	var activity models.Activity
	var activityPost requests.ActivityAddPostRequest

	errors, err3 := ctx.Request().ValidateRequest(&activityPost)

	if err3 != nil {
		common.JsonReturnError(common.RequestBindErrCode, err3.Error(), ctx)
		return
	}

	if errors != nil {
		common.JsonReturnErrorJson(common.RequestBindErrCode, errors.One(), ctx)
		return
	}

	err := activity.AddActivity(activityPost)
	if err != true {
		common.JsonReturnError(common.AddErrCode, common.AddErrMsg, ctx)
		return
	} else {
		common.JsonReturnSuccess("", common.AddSuccessMsg, ctx)
	}
}

func (a *ActivityController) UpdateActivity(ctx http.Context) {
	var activity models.Activity
	var acStruct common.ActivityUpdate
	err1 := ctx.Request().Bind(&acStruct)
	if err1 != nil {
		common.JsonReturnError(common.RequestBindErrCode, err1.Error(), ctx)
		return
	}
	err := activity.UpdateActivity(acStruct)
	if err != true {
		common.JsonReturnError(common.UpdateErrCode, common.UpdateErrMsg, ctx)
		return
	} else {
		common.JsonReturnSuccess("", common.UpdateSuccessMsg, ctx)
	}
}

func (a *ActivityController) DeleteActivity(ctx http.Context) {
	var activity models.Activity
	id := ctx.Request().Query("id", "0")

	orm_id, err1 := strconv.ParseUint(id, 10, 32)
	if orm_id == 0 {
		common.JsonReturnError(common.IdMissCode, common.IdMissMsg, ctx)
		return
	}
	if err1 != nil {
		common.JsonReturnError(common.ParseUintErrCode, err1.Error(), ctx)
		return
	}
	err := activity.DelActivity(orm_id)
	if err == false {
		common.JsonReturnError(common.DeleteErrCode, common.DeleteErrMsg, ctx)
		return
	} else {
		common.JsonReturnSuccess("", common.DeleteSuccessMsg, ctx)
	}
}
