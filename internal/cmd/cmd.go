package cmd

import (
	"context"
	"sea/internal/handler"
	serviceWater "sea/internal/service/water"

	_ "github.com/lib/pq"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gbuild"
	"github.com/gogf/gf/v2/os/gcmd"
)

func start(ctx context.Context, parser *gcmd.Parser) (err error) {
	// display welcome info
	welcome(ctx)
	// check the key pair exists
	g.Log().Info(ctx, "Checking the key pair ...")
	checkKeyPair(ctx)
	// start the web server
	g.Log().Info(ctx, "Starting http server ...")
	return routes(ctx, parser)
}

func welcome(ctx context.Context) {
	g.Log().Info(ctx, "Starting the chat server - Sea ...")
	buildInfo := gbuild.Info()
	if buildInfo["gf"] == "" {
		g.Log().Warning(ctx, "(debug version)")
	} else {
		g.Log().Infof(ctx, "built at %s in %s with gf%s", buildInfo["time"], buildInfo["go"], buildInfo["gf"])
		g.Log().Infof(ctx, "commit: %s", buildInfo["git"])
	}
}

func checkKeyPair(ctx context.Context) {
	k, err := serviceWater.WaterKey.GetSelfKey(context.Background())
	if err != nil {
		g.Log().Warningf(ctx, "Failed to get the key pair: %s, will generate...", err.Error())
		ks, err := serviceWater.GenerateKey()
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
		kid, _ := serviceWater.GetKeyID(&ks.PublicKey)
		g.Log().Infof(ctx, "Generated the key pair: %s", kid)
		_, err = serviceWater.WaterKey.AddSelfKey(context.Background(), ks)
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
		g.Log().Infof(ctx, "Added the key pair to database.")
	} else {
		g.Log().Infof(ctx, "Got the key pair: %s", k.GetKeyID())
	}
	k, _ = serviceWater.WaterKey.GetSelfKey(context.Background()) // refresh the key
	if p, err := k.GetPublicKey(); err == nil {
		g.Log().Info(ctx, "Public key: \n", p)
	} else {
		g.Log().Fatal(ctx, err)
	}
}

func routes(ctx context.Context, parser *gcmd.Parser) (err error) {
	s := g.Server()
	rootURL, err := g.Config().Get(context.Background(), "water.root")
	if err != nil {
		rootURL = gvar.New("/")
	}
	// user
	// v1UserAuth := func(group *ghttp.RouterGroup) {
	// 	// /v1/user/auth/login
	// 	group.POST("login", api.Hello.Index)
	// }
	// v1User := func(group *ghttp.RouterGroup) {
	// 	group.Group("/auth", v1UserAuth)
	// }
	// water
	// v1WaterJoin := func(group *ghttp.RouterGroup) {
	// 	group.Middleware(
	// 		waterApi.WaterInviteApiMiddleware,
	// 	)
	// 	// /v1/water/auth/join/1
	// 	group.POST("1", waterApi.WaterInvite.Step1)
	// 	// /v1/water/auth/join/2
	// 	group.POST("2", waterApi.WaterInvite.Step2)
	// }
	// v1WaterAuth := func(group *ghttp.RouterGroup) {
	// 	group.Group("/join", v1WaterJoin)
	// 	// /v1/water/auth/login
	// 	group.POST("login", api.Water.VerifyID)
	// }
	// v1WaterSync := func(group *ghttp.RouterGroup) {
	// 	group.POST("event", api.Water.VerifyID)
	// 	group.POST("verify", api.Water.VerifyID)
	// }
	v1Water := func(group *ghttp.RouterGroup) {
		// group.Group("/auth", v1WaterAuth)
		// group.Group("/sync", v1WaterSync)
		group.Bind(handler.Water)
	}
	// admin
	// v1AdminWater := func(group *ghttp.RouterGroup) {
	// 	// todo
	// 	group.POST("add", adminApi.AdminWater.AddWater)
	// 	group.POST("delete", adminApi.AdminWater.DeleteWater)
	// 	group.POST("query", adminApi.AdminWater.QueryWater)
	// 	group.POST("edit", adminApi.AdminWater.EditWater)
	// }
	v1Admin := func(group *ghttp.RouterGroup) {
		// todo
		// group.Group("/auth", v1AdminWater)
		// group.Group("/water", v1AdminWater)
		group.Bind(handler.Admin)
	}
	// main router
	root := s.Group(rootURL.String())
	root.Middleware(
		handler.HeaderMiddleware,
	)
	// root.GET("/version", api.GetVersion)
	root.Group("/v1", func(group *ghttp.RouterGroup) {
		group.Group("/water", v1Water)
		// group.Group("/user", v1User)
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
		Func:  start,
	}
)
