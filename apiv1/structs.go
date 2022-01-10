package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type VersionReq struct {
	g.Meta `path:"/ver" method:"get"`
}
type VersionRes struct {
	// g.Meta    `mime:"application/json"`
	BuildInfo map[string]string `json:"info"`
}

// water api auth join

type WaterApiJoinStep1Req struct {
	g.Meta          `path:"/auth/join/1" method:"post"`
	SenderPublicKey string `json:"sender" v:"required"` // a 4096 bits rsa public key from sender(client)
}
type WaterApiJoinStep1Res struct {
	g.Meta                     `mime:"application/json"`
	EncryptedReceiverPublicKey string `json:"receiver"` // a encrypted pack, sender can't be decrypted if haven't a private key
	ReturnCode                 int    `json:"returnCode"`
}

type WaterApiJoinStep2Req struct {
	g.Meta                `path:"/auth/join/2" method:"post"`
	EncryptedRandomString string `json:"random" v:"required"` // a encrypted pack for receiver
}
type WaterApiJoinStep2Res struct {
	g.Meta     `mime:"application/json"`
	ReturnCode int `json:"returnCode"`
}

type AdminWaterAddReq struct {
	g.Meta `path:"/water/add" method:"post"`
	URL    string `json:"url" v:"required"`
}
type AdminWaterAddRes struct {
	g.Meta     `mime:"application/json"`
	ReturnCode int `json:"returnCode"`
}

type AdminWaterDeleteReq struct {
	g.Meta `path:"/water/delete" method:"post"`
}
type AdminWaterDeleteRes struct {
	g.Meta `mime:"application/json"`
}

type AdminWaterEditReq struct {
	g.Meta `path:"/water/edit" method:"post"`
}
type AdminWaterEditRes struct {
	g.Meta `mime:"application/json"`
}

type AdminWaterQueryReq struct {
	g.Meta `path:"/water/query" method:"post"`
}
type AdminWaterQueryRes struct {
	g.Meta `mime:"application/json"`
}
