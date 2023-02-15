package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"
	"goravel/app/http/controllers"
	middleware2 "goravel/app/http/middleware"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(200, http.Json{
			"Hello": "Goravel",
		})
	})

	user := controllers.NewUserController()
	activity := controllers.NewActivityController()
	facades.Route.Post("/users/login", user.Login)
	facades.Route.Middleware(middleware2.Jwt()).Prefix("admin").Group(func(routes route.Route) {
		//用户路由
		routes.Post("/users", user.Show)
		routes.Post("/users/logout", user.Logout)
		//活动配置路由
		routes.Post("/activity/{id}", activity.GetActivity)
		routes.Post("/activity/list", activity.GetActivityList)
		routes.Post("/activity/add", activity.AddActivity)
		routes.Put("/activity/update", activity.UpdateActivity)
		routes.Delete("/activity/delete", activity.DeleteActivity)
	})

}
