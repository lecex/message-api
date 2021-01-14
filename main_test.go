package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/lecex/message-api/handler"

	conPB "github.com/lecex/message-api/proto/config"
	mesPB "github.com/lecex/message-api/proto/message"
)

func TestMessageConfigUpdate(t *testing.T) {
	req := &conPB.Request{
		Config: &conPB.Config{
			Sms: &conPB.Sms{
				Drive: "cloopen",
				Cloopen: &conPB.Cloopen{
					AppID:        "8a48b551506fd26f01509405471a6db8",
					AccountSid:   "aaf98f895069246a01506a9770ea0268",
					AccountToken: "3fd8b18597d346c48631821abc00b138",
				},
			},
			Wechat: &conPB.Wechat{
				Appid:  "8a48b551506fd26f01509405471a6db8",
				Secret: "aaf98f895069246a01506a9770ea0268",
			},
		},
	}
	res := &conPB.Response{}
	h := handler.Config{}
	err := h.Update(context.TODO(), req, res)
	t.Log(req, res, err)
}

func TestMessageSend(t *testing.T) {
	req := &mesPB.Request{
		Addressee: "13954386521",
		Event:     "register_verify",
		Type:      "wechat",
		QueryParams: `
			{
				"datas":[
					"654321",
					"5"
				]
			}
		`,
	}
	res := &mesPB.Response{}
	h := handler.Message{}
	err := h.Send(context.TODO(), req, res)
	t.Log(req, res, err)
}

func TestSmsMessageSend(t *testing.T) {
	req := &mesPB.Request{
		Addressee: "13954386521",
		Event:     "register_verify",
		Type:      "sms",
		QueryParams: `
			{
				"datas":[
					"654321",
					"5"
				]
			}
		`,
	}
	res := &mesPB.Response{}
	h := handler.Message{}
	err := h.Send(context.TODO(), req, res)
	fmt.Println(req, res, err)
	t.Log(req, res, err)
}
