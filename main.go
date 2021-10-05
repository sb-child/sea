package main

import (
	"github.com/gogf/gf/frame/g"
	_ "github.com/lib/pq"
	_ "sea/boot"
	_ "sea/router"
)

/*
db:
keys:
	id   key_type            content
	num  pubkey|privkey|sea  ...
*/

func main() {
	/*
		sea key
		(sea id = hash(sea key))

		water key
			private key
			public key
		(water id = hash(water public key))
	*/
	g.Server().Run()
}
