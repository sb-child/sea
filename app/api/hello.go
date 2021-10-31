package api

import (
	"sea/app/service"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

var Hello = helloApi{}

type helloApi struct{}

func (*helloApi) Index(r *ghttp.Request) {
	service.Water.ReGenWaterID()
	w, _ := service.Water.GetSelfWater()
	k, _ := crypto.GenerateKey("name", "email", "rsa", 4096)
	ks, _ := k.Armor()
	armored, _ := helper.SignCleartextMessageArmored(ks, nil, w.WaterId)
	r.Response.WriteJson(g.MapStrStr{
		"WaterId": w.WaterId,
		"SeaId":   gconv.String(armored),
	})
}
