package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"goravel/app/http/common"
	"goravel/app/models"
)

type UserController struct {
	*common.User
}

type ReUser struct {
	common.ReturnUser
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) {
	data := ctx.Value("userData")
	ctx.Response().Success().Json(http.Json{
		"code": 20000,
		"user": data,
	})
}

func (r *UserController) RefreshToken(ctx http.Context) {

}

func (u *UserController) Login(ctx http.Context) {
	err := ctx.Request().Bind(&u)
	if err != nil {
		common.JsonReturnError(common.RequestBindErrCode, err.Error(), ctx)
		return
	}
	var user models.User
	user.Name = u.Username
	isTure, msg := user.ContrastPassword(u.Password)
	if isTure.Name == "" || isTure.Password == "" {
		common.JsonReturnError(common.UserNameOrPassErrCode, msg, ctx)
		return
	}
	user.ID = isTure.Id
	token, err1 := facades.Auth.Login(ctx, &user)
	if err1 != nil {
		common.JsonReturnError(common.AuthorizationCreateErrCode, err1.Error(), ctx)
		return
	}
	ctx.Response().Success().Json(http.Json{
		"code":  20000,
		"token": token,
	})
}

func (u *UserController) Logout(ctx http.Context) {
	err := facades.Auth.Logout(ctx)
	if err != nil {
		common.JsonReturnError(common.AuthorizationCreateErrCode, err.Error(), ctx)
		return
	}
	common.JsonReturnSuccess("", "success", ctx)
}
