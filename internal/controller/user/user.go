package user

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"learn_golang/api/user"
)

type Controller struct{}

func Router() *Controller {
	return &Controller{}
}

func (c *Controller) Get(ctx context.Context, req *user.GetReq) (res *user.GetResp, err error) {
	r := g.RequestFromCtx(ctx)
	userId := r.Get("id", 1).Int()
	res = &user.GetResp{
		Id:   userId,
		Name: "张三",
		Age:  18,
	}
	return
}

func (c *Controller) Add(ctx context.Context, req *user.AddReq) (res *user.AddResp, err error) {
	return
}

func (c *Controller) Update(ctx context.Context, req *user.UpdateReq) (res *user.UpdateResp, err error) {
	return
}

func (c *Controller) Delete(ctx context.Context, req *user.DeleteReq) (res *user.Delete, err error) {
	return
}
