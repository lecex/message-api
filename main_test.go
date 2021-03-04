package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/lecex/message-api/handler"

	mesPB "github.com/lecex/message-api/proto/message"
)

// func TestMessageConfigUpdate(t *testing.T) {
// 	req := &conPB.Request{
// 		Config: &conPB.Config{
// 			Sms: &conPB.Sms{
// 				Drive: "cloopen",
// 				Cloopen: &conPB.Cloopen{
// 					AppID:        "",
// 					AccountSid:   "",
// 					AccountToken: "",
// 				},
// 			},
// 			Wechat: &conPB.Wechat{
// 				Appid:  "",
// 				Secret: "",
// 			},
// 		},
// 	}
// 	res := &conPB.Response{}
// 	h := handler.Config{}
// 	err := h.Update(context.TODO(), req, res)
// 	t.Log(req, res, err)
// }

// func TestMessageSend(t *testing.T) {
// 	req := &mesPB.Request{
// 		Addressee: "13954386521",
// 		Event:     "register_verify",
// 		Type:      "wechat",
// 		QueryParams: `
// 			{
// 				"datas":[
// 					"654321",
// 					"5"
// 				]
// 			}
// 		`,
// 	}
// 	res := &mesPB.Response{}
// 	h := handler.Message{}
// 	err := h.Send(context.TODO(), req, res)
// 	t.Log(req, res, err)
// }

func TestSmsMessageSend(t *testing.T) {
	req := &mesPB.Request{
		Addressee: "13954386521",
		Event:     "sms_captcha",
	}
	res := &mesPB.Response{}
	h := handler.Message{"go.micro.srv.message"}
	err := h.SendCaptcha(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
