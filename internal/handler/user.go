package handler

import (
	"context"

	"sea/apiv1"
)

var (
	User = hUser{}
)

type hUser struct{}

func (h *hUser) Login(ctx context.Context, req *apiv1.UserLoginReq) (res *apiv1.UserLoginRes, err error) {
	res, err = apiv1.UserAuth.Login(ctx, req)
	return
}
