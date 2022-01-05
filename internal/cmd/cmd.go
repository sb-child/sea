package cmd

import (
	"context"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func routes(ctx context.Context, parser *gcmd.Parser) (err error) {
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
	// admin
	v1AdminWater := func(group *ghttp.RouterGroup) {
		// todo
		group.POST("add", adminApi.AdminWater.AddWater)
		group.POST("delete", adminApi.AdminWater.DeleteWater)
		group.POST("query", adminApi.AdminWater.QueryWater)
		group.POST("edit", adminApi.AdminWater.EditWater)
	}
	v1Admin := func(group *ghttp.RouterGroup) {
		// todo
		group.Group("/auth", v1AdminWater)
		group.Group("/water", v1AdminWater)
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
		group.Group("/admin", v1Admin)
	})
	s.Run()

	return nil
}

var (
	Main = gcmd.Command{
		Name:  "sea",
		Usage: "sea",
		Brief: "start sea server",
		Func:  routes,
	}
)
