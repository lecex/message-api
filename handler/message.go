package handler

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
	client "github.com/lecex/core/client"
	uuid "github.com/satori/go.uuid"

	pb "github.com/lecex/message-api/proto/message"
	"github.com/lecex/message-api/providers/redis"
)

// 默认验证码过期时间
var expireTime time.Duration = 5

// Message 权限结构
type Message struct {
	ServiceName string
}

// SmsVerifySend 手机短信验证码发送
func (srv *Message) SmsVerifySend(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	if req.Addressee == "" {
		return errors.New("Empty Addressee")
	}
	if req.Event == "" {
		return errors.New("Empty Event")
	}
	uuid, vrify, err := srv.Verify()

	req.Type = "sms"
	req.QueryParams = `
		{
			"datas":[
				"` + vrify + `",
				"` + string(expireTime) + `"
			]
		}
	`
	a := uuid
	return client.Call(ctx, srv.ServiceName, "Message.Send", req, res)
}

// Verify 生成验证码并保存到 redis
func (srv *Message) Verify() (key string, vrify string, err error) {
	key = uuid.NewV4().String()
	vrify = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	redis := redis.NewClient()
	// 过期时间默认30分钟
	err = redis.SetNX(key, vrify, expireTime*time.Minute).Err()
	return key, vrify, err
}
