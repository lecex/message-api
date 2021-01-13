package handler

import (
	"context"

	client "github.com/lecex/core/client"

	pb "github.com/lecex/message-api/proto/template"
)

// Template 权限结构
type Template struct {
	ServiceName string
}

// List 权限列表
func (srv *Template) List(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Templates.List", req, res)
}

// Get 获取权限
func (srv *Template) Get(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Templates.Get", req, res)
}

// Create 创建权限
func (srv *Template) Create(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Templates.Create", req, res)
}

// Update 更新权限
func (srv *Template) Update(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Templates.Update", req, res)
}

// Delete 删除权限
func (srv *Template) Delete(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Templates.Delete", req, res)
}
