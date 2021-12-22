package background

import (
	"github.com/gogf/gf/v2/net/gclient"
)

var Client *gclient.Client

func BindCliect(cli *gclient.Client) {
	Client = cli
}

func GetCliect() *gclient.Client {
	return Client
}
