package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserLoginReq struct {
	g.Meta   `path:"/login" method:"post"`
	Username string `p:"username" v:"required"`
	Password string `p:"password" v:"required"`
}
type UserLoginRes struct {
	g.Meta     `mime:"application/json"`
	ReturnCode int `json:"returnCode"`
}
