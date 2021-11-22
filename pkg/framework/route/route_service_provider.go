package route

import (
	"github.com/goravel/framework/support/facades"
)

type RouteServiceProvider struct {
}

func (router *RouteServiceProvider) Boot() {
	gin := Gin{}
	facades.Route = gin.Init()
}

func (router *RouteServiceProvider) Register() {

}
