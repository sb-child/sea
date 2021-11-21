package main

import (
	"context"
	serviceWater "sea/app/service/water"
	_ "sea/boot"
	_ "sea/router"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gbuild"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	// display welcome info
	buildInfo := gbuild.Info()
	g.Log().Infof(ctx, "gf version: %s", buildInfo["gf"])
	g.Log().Infof(ctx, "go version: %s", buildInfo["go"])
	g.Log().Infof(ctx, "build at: %s", buildInfo["time"])
	g.Log().Infof(ctx, "commit: %s", buildInfo["git"])
	g.Log().Info(ctx, "Starting the chat server - Sea ...")
	// check the key pair exists
	k, err := serviceWater.WaterKey.GetSelfKey(context.Background())
	if err != nil {
		g.Log().Warningf(ctx, "Failed to get the key pair: %s, will generate...", err.Error())
		ks, err := serviceWater.GenerateKey()
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
		kid, _ := serviceWater.GetKeyID(ks)
		g.Log().Infof(ctx, "Generated the key pair: %s", kid)
		_, err = serviceWater.WaterKey.AddKey(context.Background(), ks, true)
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
		g.Log().Infof(ctx, "Added the key pair to database.")
	} else {
		g.Log().Infof(ctx, "Got the key pair: %s", k.GetKeyID())
	}
	g.Log().Info(ctx, "Finally, start the web server...")
	// run web server
	g.Server().Run()
}
