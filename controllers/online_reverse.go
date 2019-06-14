package controllers

import (
	"fmt"
	"html/template"
	// yuan "github.com/yuansfer/golang_sdk"
)

const (
	ReverseUrl = "appTransaction/v2/securepayVoid"
)

type OnlineReverse struct {
	MerchantNo    string `json:"merchantNo"`
	StoreNo       string `json:"storeNo"`
	TransactionNo string `json:"transactionNo"`
	VoidAmount    string `json:"voidAmount"`
	Reference     string `json:"reference"`
}

type ReverseController struct {
	Controller
}

func (this *ReverseController) Get() {
	this.Data["IsOnlineReverse"] = true
	this.TplName = "online-reverse.tpl"
}

func (this *ReverseController) Post() {

	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	voidAmount := this.Input().Get("voidAmount")
	transactionNo := this.Input().Get("transactionNo")
	reference := this.Input().Get("reference")

	req := OnlineReverse{
		MerchantNo:    merchantNo,
		StoreNo:       storeNo,
		TransactionNo: transactionNo,
		VoidAmount:    voidAmount,
		Reference:     reference,
	}

	resp, err := req.PostToYuansfer()
	if err != nil {
		fmt.Println(err)
		this.Ctx.WriteString("something wrong happened")
	}
	t := template.New("void-template")
	t, _ = t.Parse(resp)
	_ = t.Execute(this.Ctx.ResponseWriter, resp)

	return
}

func (r OnlineReverse) PostToYuansfer() (string, error) {

	values := generateValues(r, YuansferAPI.Token.SecurepayToken)
	fmt.Println("online reverse request values:")
	fmt.Println(values)
	reverseUrl := YuansferHost + ReverseUrl
	fmt.Println("online reverse request url:", reverseUrl)
	return postToYuansfer(reverseUrl, values)
}
