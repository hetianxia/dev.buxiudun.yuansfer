package controllers

import (
	"fmt"
	"html/template"
)

type InstoreDetailController struct {
	Controller
}

const (
	instoreDetailUrl = "app-instore/v2/detail"
)

type InstoreDetail struct {
	MerchantNo    string `json:"merchantNo"`
	StoreNo       string `json:"storeNo"`
	VerifySign    string `json:"verifySign"`
	StoreAdminNo  string `json:"storeAdminNo"`
	TransactionNo string `json:"transactionNo"`
}

func (this *InstoreDetailController) Get() {
	this.Data["IsInstoreDetail"] = true
	this.TplName = "instore-detail.tpl"
}

func (this *InstoreDetailController) Post() {
	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	storeAdminNo := this.Input().Get("storeAdminNo")
	transNo := this.Input().Get("transNo")

	req := InstoreDetail{
		MerchantNo:    merchantNo,
		StoreNo:       storeNo,
		StoreAdminNo:  storeAdminNo,
		TransactionNo: transNo,
	}

	resp, err := req.PostToYuansfer()

	if err != nil {
		fmt.Println(err)
		this.Ctx.WriteString("something wrong happened")
	}

	t := template.New("recurring-template")
	t, _ = t.Parse(resp)
	_ = t.Execute(this.Ctx.ResponseWriter, resp)

	return
}

func (r InstoreDetail) PostToYuansfer() (string, error) {
	var (
		detailUrl string
	)

	detailUrl = YuansferHost + instoreDetailUrl

	values := generateValues(r, YuansferAPI.Token.InstoreToken)
	return postToYuansfer(detailUrl, values)
}
