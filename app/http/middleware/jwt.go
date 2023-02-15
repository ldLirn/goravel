/**
 * @Author: lirn
 * @Description:
 * @File: Jwt
 * @Date: 2022/12/26 17:31
 */
package middleware

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/http/common"
	"goravel/app/models"
)

//判断jwt状态并解析用户信息
func Jwt() http.Middleware {
	return func(ctx http.Context) {
		Authorization := ctx.Request().Header("X-Token", "")

		if Authorization == "" {
			common.JsonReturnError(common.AuthorizationMissCode, common.AuthorizationMissMsg, ctx)
			return
		}

		var user models.User
		err := facades.Auth.Parse(ctx, Authorization)
		err1 := facades.Auth.User(ctx, &user)

		if err != nil || err1 != nil {
			common.JsonReturnError(common.AuthorizationParseCode, "错误信息1："+err1.Error()+"  错误信息2："+err.Error(), ctx)
			return
		}
		ctx.WithValue("userData", user)
		ctx.WithValue("userId", user.ID)

		ctx.Request().Next()
	}
}
