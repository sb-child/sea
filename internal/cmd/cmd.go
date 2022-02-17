package cmd

import (
	"context"
	bg "sea/internal/background"
	"sea/internal/handler"
	"sea/internal/service"

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
	g.Log().Info(ctx, "Checking main key pair ...")
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
		g.Log().Warning(ctx, "(debug build, use `gf build` for release)")
	} else {
		g.Log().Infof(ctx, "built at %s in %s with gf%s", buildInfo["time"], buildInfo["go"], buildInfo["gf"])
		g.Log().Infof(ctx, "commit: %s", buildInfo["git"])
	}
}

func checkKeyPair(ctx context.Context) {
	k, err := service.WaterKey.GetSelfKey(context.Background())
	if err != nil {
		g.Log().Warningf(ctx, "Failed to get key pair: %s, will generate...", err.Error())
		ks, err := service.GenerateKey()
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
		kid, _ := service.GetKeyID(&ks.PublicKey)
		g.Log().Infof(ctx, "Generated key pair: %s", kid)
		_, err = service.WaterKey.AddSelfKey(context.Background(), ks)
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
		g.Log().Infof(ctx, "Added key pair to database.")
	} else {
		g.Log().Infof(ctx, "Got key pair: %s", k.GetKeyID())
	}
	k, _ = service.WaterKey.GetSelfKey(ctx) // refresh the key
	if kp, err := k.GetPublicKey(); (err == nil) && (kp != nil) {
		kid, err := service.GetKeyID(kp)
		if (kid == "") || (err != nil) || (k.GetKeyID() != kid) {
			g.Log().Warningf(ctx, "Key ID should be: %s", kid)
			g.Log().Warningf(ctx, "but got         : %s", k.GetKeyID())
			g.Log().Fatal(ctx, "Key ID mismatch, please remove this key pair and generate a new one.")
		}
		g.Log().Info(ctx, "Public key test passed.")
	} else {
		g.Log().Fatal(ctx, err)
	}
	if kp, err := k.GetPrivateKey(); (err == nil) && (kp.Validate() == nil) {
		g.Log().Info(ctx, "Private key test passed.")
	} else {
		g.Log().Fatal(ctx, err)
	}
	// ks, _ := service.GenerateKey()
	// kss, _ := service.PackPublicKey(&ks.PublicKey)
	// g.Log().Debug(ctx, kss)
}

func routes(ctx context.Context, parser *gcmd.Parser) (err error) {
	s := g.Server()
	s.SetErrorStack(true)
	rootURL, err := g.Config().Get(context.Background(), "water.root")
	if err != nil {
		rootURL = gvar.New("/")
	}
	v1Water := func(group *ghttp.RouterGroup) {
		group.Bind(handler.Water)
	}
	v1User := func(group *ghttp.RouterGroup) {
		group.Bind(handler.User)
	}
	v1Admin := func(group *ghttp.RouterGroup) {
		group.Bind(handler.Admin)
	}
	v1Stream := func(group *ghttp.RouterGroup) {
		group.Bind(handler.WaterStream)
	}
	// main router
	root := s.Group(rootURL.String())
	root.Middleware(
		ghttp.MiddlewareHandlerResponse,
		handler.HeaderMiddleware,
	)
	root.Bind(handler.Version)
	root.Group("/v1", func(group *ghttp.RouterGroup) {
		group.Group("/water", v1Water)
		group.Group("/user", v1User)
		group.Group("/admin", v1Admin)
		group.Group("/stream", v1Stream)
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
