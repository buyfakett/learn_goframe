package user

import "github.com/gogf/gf/v2/frame/g"

type GetReq struct {
	g.Meta `path:"/get" method:"get"`
	id     int
}

type GetResp struct {
	Id   int
	Name string
	Age  int
}

type AddReq struct {
	g.Meta `path:"/create" method:"post"`
}
type AddResp struct{}

type UpdateReq struct {
	g.Meta `path:"/update" method:"put"`
}
type UpdateResp struct{}

type DeleteReq struct {
	g.Meta `path:"/delete" method:"delete"`
}
type Delete struct{}
