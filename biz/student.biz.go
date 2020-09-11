package biz

import (
	"context"
	"github.com/obase/demo/api"
	"github.com/obase/demo/dao"
)

type StudentServiceService struct{}

var _ api.StudentServiceServer = (*StudentServiceService)(nil)

func (s *StudentServiceService) Add(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	return dao.Add(ctx, req)
}
func (s *StudentServiceService) Upd(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	return dao.Upd(ctx, req)
}
func (s *StudentServiceService) Del(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	return dao.Del(ctx, req)
}
func (s *StudentServiceService) Get(ctx context.Context, req *api.Student) (rsp *api.Student, err error) {
	return dao.Get(ctx, req)
}
func (s *StudentServiceService) Page(ctx context.Context, req *api.PageReq) (rsp *api.PageRsp, err error) {
	return dao.Page(ctx, req)
}
