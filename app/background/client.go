package background

import (
	"time"

	"github.com/gogf/gf/v2/net/gclient"
)

var Client *gclient.Client

func Init() {
	Client = Client.Timeout(5 * time.Second)
}
