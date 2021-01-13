package config

import (
	"github.com/lecex/core/config"
	"github.com/lecex/core/env"
	PB "github.com/lecex/user/proto/permission"
)

// 	Conf 配置
// 	Name // 服务名称
//	Method // 方法
//	Auth // 是否认证授权
//	Policy // 是否认证权限
//	Name // 权限名称
//	Description // 权限解释
var Conf config.Config = config.Config{
	Name:    env.Getenv("MICRO_API_NAMESPACE", "go.micro.api.") + "message-api",
	Version: "latest",
	Service: map[string]string{
		"message": env.Getenv("USER_SERVICE", "go.micro.srv.message"),
		"user":    env.Getenv("USER_SERVICE", "go.micro.srv.user"),
	},
	Permissions: []*PB.Permission{ // 权限管理
		{Service: "message-api", Method: "Configs.Get", Auth: true, Policy: true, Name: "配置获取", Description: "消息相关配置获取"},
		{Service: "message-api", Method: "Configs.Update", Auth: true, Policy: true, Name: "配置更新", Description: "消息相关配置更新"},
		// 消息发送
		{Service: "message-api", Method: "Message.Send", Auth: false, Policy: false, Name: "消息发送", Description: "消息短信消息微信消息邮件等发送"},
		// 模板管理
		{Service: "message-api", Method: "Template.Create", Auth: true, Policy: true, Name: "创建模板", Description: "创建新模板权限。"},
		{Service: "message-api", Method: "Template.Delete", Auth: true, Policy: true, Name: "删除模板", Description: "删除模板。"},
		{Service: "message-api", Method: "Template.Update", Auth: true, Policy: true, Name: "更新模板", Description: "更新模板信息。"},
		{Service: "message-api", Method: "Template.Get", Auth: true, Policy: true, Name: "查询模板", Description: "查询模板信息权限。"},
		{Service: "message-api", Method: "Template.List", Auth: true, Policy: true, Name: "模板列表", Description: "查询模板列表"},
	},
}
