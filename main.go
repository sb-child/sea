package main

import (
	_ "sea/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	"sea/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
