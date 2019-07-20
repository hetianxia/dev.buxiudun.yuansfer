package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type Controller struct {
	beego.Controller
}

func (c *Controller) checkData(key string, value interface{}) {
	if c.Data[key] = value; c.Data[key] == "" {
		c.Data[key] = "-"
	}
}

func init() {
	//跨域设置
	var FilterGateWay = func(ctx *context.Context) {
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
		//允许访问源
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
		//允许post访问
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,ContentType,Authorization,accept,accept-encoding, authorization, content-type") //header的类型
		ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	}
	beego.InsertFilter("*", beego.BeforeRouter, FilterGateWay)
}
