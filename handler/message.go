package handler

import (
	"context"

	client "github.com/lecex/core/client"

	pb "github.com/lecex/message-api/proto/message"
)

// Message 权限结构
type Message struct {
	ServiceName string
}

// Send 消息发送
func (srv *Message) Send(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	return client.Call(ctx, srv.ServiceName, "Messages.Send", req, res)
}
