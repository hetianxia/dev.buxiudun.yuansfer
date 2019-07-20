package controllers

import (
	"fmt"
)

type WxController struct {
	Controller
}

// Get to veirfy service from wechat
func (w *WxController) Get() {
	fmt.Println("--WxController--GET-----")
	echoStr := w.Input().Get("echostr")
	fmt.Println("echostr:", echoStr)
	w.Ctx.WriteString(echoStr)
}

func (w *WxController) Post() {

	w.Ctx.WriteString("Hello World")
}

func (w *WxController) Sign() {

	var ret string
	addr := w.Input().Get("addr")
	if "" != addr {
		ret = genSignature(addr)
	}
	w.Ctx.WriteString(ret)
}
