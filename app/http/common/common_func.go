/**
 * @Author: lirn
 * @Description:
 * @File: common_func
 * @Date: 2022/12/27 15:12
 */
package common

import (
	"github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"reflect"
	"strconv"
	"time"
)

//通用分页查询
func Paginator(page string, limit string) func(methods orm.Query) orm.Query {
	return func(query orm.Query) orm.Query {
		page, _ := strconv.Atoi(page)
		limit, _ := strconv.Atoi(limit)
		offset := (page - 1) * limit
		return query.Offset(offset).Limit(limit)
	}
}

//返回成功信息
func JsonReturnSuccess(data interface{}, info string, ctx http.Context) {
	ctx.Response().Success().Json(http.Json{
		"code": 20000,
		"msg":  "success",
		"info": info,
		"data": data,
	})
}

func JsonReturnError(code int, info interface{}, ctx http.Context) {
	ctx.Request().AbortWithStatusJson(200, http.Json{
		"code": code,
		"msg":  "error",
		"info": info,
		"data": "",
	})
}

func JsonReturnErrorJson(code int, info interface{}, ctx http.Context) {
	ctx.Response().Success().Json(http.Json{
		"code": code,
		"msg":  "error",
		"info": info,
		"data": "",
	})
}

func JsonReturnErrorMap(code int, info map[string]map[string]string, ctx http.Context) {
	ctx.Request().AbortWithStatusJson(code, http.Json{
		"code": code,
		"msg":  "error",
		"info": info,
		"data": "",
	})
}

func AuthCheck(checkName string, arguments map[string]any, ctx http.Context) bool {
	response := facades.Gate.Inspect(checkName, arguments)
	if response.Allowed() == false {
		return false
	}
	return true
}

//转换时间戳
func StringToUnix(t string) int64 {
	fT, err := time.Parse("2006-01-02 15:04:05", t)
	if err == nil {
		return fT.Unix()
	} else {
		return 0
	}
}

//将结构体source复制给dst，只复制相同名称和相同类型的
//CopyStruct(a,b)  a可以传值，引用，b只能引用，且
func CopyStruct(src, dst interface{}) interface{} {
	st := reflect.TypeOf(src)
	sv := reflect.ValueOf(src)
	dt := reflect.TypeOf(dst)
	dv := reflect.ValueOf(dst)
	if st.Kind() == reflect.Ptr { //处理指针
		st = st.Elem()
		sv = sv.Elem()
	}
	if dt.Kind() == reflect.Ptr { //处理指针
		dt = dt.Elem()
	}
	if st.Kind() != reflect.Struct || dt.Kind() != reflect.Struct { //如果不是struct类型，直接返回dst
		return dst
	}

	dv = reflect.ValueOf(dv.Interface())
	// 遍历TypeOf 类型
	for i := 0; i < dt.NumField(); i++ { //通过索引来取得它的所有字段，这里通过t.NumField来获取它多拥有的字段数量，同时来决定循环的次数
		f := dt.Field(i) //通过这个i作为它的索引，从0开始来取得它的字段
		dVal := dv.Elem().Field(i)
		sVal := sv.FieldByName(f.Name)
		//fmt.Println(dVal.CanSet())
		//src数据有效，且dst字段能赋值,类型一致
		if sVal.IsValid() && dVal.CanSet() && f.Type.Kind() == sVal.Type().Kind() {
			dVal.Set(sVal)
		}
	}
	return dst
}
