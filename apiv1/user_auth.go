package apiv1

import (
	"context"
)

var UserAuth = userAuth{}

type userAuth struct{}

func (api *userAuth) Login(ctx context.Context, req *UserLoginReq) (*UserLoginRes, error) {
	return &UserLoginRes{}, nil
}
