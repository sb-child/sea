package router

import (
	"context"
	"sea/app/api"
	waterApi "sea/app/api/water"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func init() {
	s := g.Server()
	rootURL, err := g.Config().Get(context.Background(), "water.root")
	if err != nil {
		rootURL = gvar.New("/")
	}
	root := s.Group(rootURL.String())
	// user
	v1User := func(group *ghttp.RouterGroup) {
		group.Middleware(
			api.CookieMiddleware,
		)
		group.ALL("login", api.Hello)
	}
	// water
	v1WaterJoin := func(group *ghttp.RouterGroup) {
		group.Middleware(
			waterApi.WaterInviteApiMiddleware,
			api.CookieMiddleware,
		)
		group.POST("1", waterApi.WaterInvite.VerifyID)
		group.POST("2", waterApi.WaterInvite.VerifyID)
	}
	v1WaterAuth := func(group *ghttp.RouterGroup) {
		group.Middleware(
			api.CookieMiddleware,
		)
		group.Group("join", v1WaterJoin)
		group.POST("login", api.Water.VerifyID)
	}
	v1WaterSync := func(group *ghttp.RouterGroup) {
		group.Middleware(
			api.CookieMiddleware,
		)
		group.POST("event", api.Water.VerifyID)
		group.POST("verify", api.Water.VerifyID)
	}
	v1Water := func(group *ghttp.RouterGroup) {
		group.Middleware(
			api.CookieMiddleware,
		)
		group.Group("auth/", v1WaterAuth)
		group.Group("sync/", v1WaterSync)
	}
	// main router
	root.GET("version", api.GetVersion)
	root.Group("v1/", func(group *ghttp.RouterGroup) {
		group.Group("water/", v1Water)
		group.Group("user/", v1User)
	})
}
