package background

import (
	"context"
	bg "sea/app/background"

	"github.com/gogf/gf/v2/net/gclient"
)

var (
	client = bg.GetCliect()
	WaterManager = waterManager{
		cli: client,
	}
)


type waterManager struct {
	cli *gclient.Client
}

func (m *waterManager) AddWater(ctx context.Context, url string) (error) {
	m.cli.PostVar(ctx, url, nil)
	return nil
}
