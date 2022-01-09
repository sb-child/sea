package cmd

import (
	"context"
	bg "sea/internal/background"
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
	// start background tasks
	g.Log().Info(ctx, "Starting background worker ...")
	bg.Client = g.Client()
	bg.Init()
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
	v1Water := func(group *ghttp.RouterGroup) {
		group.Bind(handler.Water)
	}
	v1Admin := func(group *ghttp.RouterGroup) {
		group.Bind(handler.Admin)
	}
	// main router
	root := s.Group(rootURL.String())
	root.Middleware(
		handler.HeaderMiddleware,
	)
	root.Bind(handler.Version)
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
