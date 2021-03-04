package handler

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	client "github.com/lecex/core/client"
	pb "github.com/lecex/message-api/proto/message"
	"github.com/lecex/message-api/providers/redis"
)

// 默认验证码过期时间
var expireTime int64 = 5

// Message 权限结构
type Message struct {
	ServiceName string
}

// VerifySend 手机短信验证码发送
func (srv *Message) VerifySend(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	if req.Addressee == "" {
		return errors.New("Empty Addressee")
	}
	if req.Event == "" {
		return errors.New("Empty Event")
	}
	vrify, err := srv.Verify(req.Addressee)
	if err != nil {
		return err
	}
	req.Type = "sms"
	req.QueryParams = `{
			"datas":[
				"` + vrify + `",
				"` + strconv.FormatInt(expireTime, 10) + `"
			]
		}`
	err = client.Call(ctx, srv.ServiceName, "Message.Send", req, res)
	if err != nil {
		return err
	}
	return err
}

// Verify 生成验证码并保存到 redis 【临时存放 以后邮件使用生成验证码可以使用】
func (srv *Message) Verify(Addressee string) (key string, vrify string, err error) {
	key = "verify_" + Addressee
	vrify = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	redis := redis.NewClient()
	// 过期时间默认5分钟
	err = redis.SetNX(key, vrify, time.Duration(expireTime)*time.Minute).Err()
	return key, vrify, err
}
