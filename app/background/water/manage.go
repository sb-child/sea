package background

import (
	"context"
	bg "sea/app/background"

	"github.com/gogf/gf/v2/net/gclient"
)

var (
	WaterManager = waterManager{
		cli: &bg.Client,
	}
)

type waterManager struct {
	cli **gclient.Client
}

func (m *waterManager) GetClient() *gclient.Client {
	return *m.cli
}

func (m *waterManager) AddWater(ctx context.Context, url string) error {
	m.GetClient().PostVar(ctx, url)
	return nil
}
