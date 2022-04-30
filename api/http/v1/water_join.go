package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type WaterApiJoinStep1Req struct {
	g.Meta          `path:"/auth/join/1" method:"post"`
	SenderPublicKey string `p:"sender" v:"required"` // a 4096 bits rsa public key from sender(client)
}
type WaterApiJoinStep1Res struct {
	g.Meta                     `mime:"application/json"`
	EncryptedReceiverPublicKey string `json:"receiver"` // a encrypted pack, sender can't be decrypted if haven't a private key
	ReturnCode                 int    `json:"returnCode"`
}

type WaterApiJoinStep2Req struct {
	g.Meta                `path:"/auth/join/2" method:"post"`
	EncryptedRandomString string `p:"random" v:"required"` // a encrypted pack for receiver
}
type WaterApiJoinStep2Res struct {
	g.Meta     `mime:"application/json"`
	ReturnCode int `json:"returnCode"`
}
