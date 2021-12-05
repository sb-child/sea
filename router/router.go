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
	// user
	v1UserAuth := func(group *ghttp.RouterGroup) {
		// /v1/user/auth/login
		group.POST("login", api.Hello.Index)
	}
	v1User := func(group *ghttp.RouterGroup) {
		group.Group("/auth", v1UserAuth)
	}
	// water
	v1WaterJoin := func(group *ghttp.RouterGroup) {
		group.Middleware(
			waterApi.WaterInviteApiMiddleware,
		)
		// /v1/water/auth/join/1
		group.POST("1", waterApi.WaterInvite.Step1)
		// /v1/water/auth/join/2
		group.POST("2", waterApi.WaterInvite.Step2)
	}
	v1WaterAuth := func(group *ghttp.RouterGroup) {
		group.Group("/join", v1WaterJoin)
		// /v1/water/auth/login
		group.POST("login", api.Water.VerifyID)
	}
	v1WaterSync := func(group *ghttp.RouterGroup) {
		group.POST("event", api.Water.VerifyID)
		group.POST("verify", api.Water.VerifyID)
	}
	v1Water := func(group *ghttp.RouterGroup) {
		group.Group("/auth", v1WaterAuth)
		group.Group("/sync", v1WaterSync)
	}
	// main router
	root := s.Group(rootURL.String())
	root.Middleware(
		api.HeaderMiddleware,
	)
	root.GET("/version", api.GetVersion)
	root.Group("/v1", func(group *ghttp.RouterGroup) {
		group.Group("/water", v1Water)
		group.Group("/user", v1User)
	})
}
