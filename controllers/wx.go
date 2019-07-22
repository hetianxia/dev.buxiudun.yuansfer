package controllers

type WxController struct {
	Controller
}

// Get the response for veirfy-service from wechat
func (w *WxController) Get() {
	echoStr := w.Input().Get("echostr")
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
