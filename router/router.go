package router

import (
	"sea/app/api"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	root := s.Group("_sea-ety729053r/")
	// user
	v1User := func(group *ghttp.RouterGroup) {
		group.ALL("login", api.Hello)
	}
	// water
	v1WaterAuthAdd := func(group *ghttp.RouterGroup) {
		group.POST("1", api.Water.VerifyID)
		group.POST("2", api.Water.VerifyID)
	}
	v1WaterAuth := func(group *ghttp.RouterGroup) {
		group.Group("add", v1WaterAuthAdd)
		group.POST("login", api.Water.VerifyID)
	}
	v1WaterSync := func(group *ghttp.RouterGroup) {
		group.POST("event", api.Water.VerifyID)
		group.POST("verify", api.Water.VerifyID)
	}
	v1Water := func(group *ghttp.RouterGroup) {
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
