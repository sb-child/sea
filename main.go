package main

import (
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
	_ "sea/internal/packed"

	"sea/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
