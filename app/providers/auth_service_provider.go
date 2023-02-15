package providers

import (
	"context"
	"github.com/goravel/framework/contracts/auth/access"
	"github.com/goravel/framework/facades"
)

type AuthServiceProvider struct {
}

func (receiver *AuthServiceProvider) Register() {

}

func (receiver *AuthServiceProvider) Boot() {
	facades.Gate.Define("check", func(ctx context.Context, arguments map[string]any) *access.Response {
		if arguments["userId"] == uint(1) { //超级管理员,拥有所有权限
			return access.NewAllowResponse()
		} else {
			return access.NewDenyResponse("没有权限")
		}

	})
}
