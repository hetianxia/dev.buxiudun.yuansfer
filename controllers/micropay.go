package controllers

import (
	"fmt"
	"html/template"
	"log"
	"time"

	yuan "github.com/yuansfer/golang_sdk"
)

type MicropayController struct {
	Controller
}

func (this *MicropayController) Get() {
	reference := fmt.Sprintf("seq_%d", time.Now().Unix())
	reference = md5Token(reference)
	this.Data["reference"] = reference
	this.Data["IsMicroPay"] = true
	this.TplName = "micropay.tpl"
}

func (this *MicropayController) Post() {
	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	amt := this.Input().Get("amt")
	rmbAmt := this.Input().Get("rmbAmt")
	reference := this.Input().Get("reference")
	description := this.Input().Get("description")
	note := this.Input().Get("note")
	ipnUrl := this.Input().Get("ipnUrl")
	openid := this.Input().Get("openid")
	vendor := this.Input().Get("vendor")
	terminal := this.Input().Get("terminal")

	if "" == reference {
		reference = fmt.Sprintf("seq_%d", time.Now().Unix())
	}
	req := &yuan.Micropay{
		MerchantNo:  merchantNo,
		StoreNo:     storeNo,
		Currency:    "USD",
		Amount:      amt,
		RmbAmount:   rmbAmt,
		Vendor:      vendor,
		Reference:   reference,
		IpnURL:      ipnUrl,
		Description: description,
		Note:        note,
		Openid:      openid,
		Terminal:    terminal,
	}

	resp, err := req.PostToYuansfer()
	if err != nil {
		log.Println(err)
		this.Ctx.WriteString("something wrong happened")
	}
	t := template.New("secure pay template")
	t, _ = t.Parse(resp)
	_ = t.Execute(this.Ctx.ResponseWriter, resp)

	return
}
