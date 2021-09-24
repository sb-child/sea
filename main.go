package main

import (
	_ "sea/boot"
	_ "sea/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
