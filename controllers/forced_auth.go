package controllers

import (
	"fmt"
	"html/template"
)

const (
	ForcedAuthUrl = "creditpay/v2/credit-force"
)

type ForcedAuthController struct {
	Controller
}

type ForcedAuth struct {
	MerchantNo     string `json:"merchantNo"`
	StoreNo        string `json:"storeNo"`
	AuthCode       string `json:"authCode"`
	CardNumber     string `json:"cardNumber"`
	ExpirationDate string `json:"expirationDate"`
	Amt            string `json:"transactionAmount"`
	Address        string `json:"addressLine1"`
	Zip            string `json:"zip"`
	CardType       string `json:"cardType"`

	RetCode string `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
}

func (this *ForcedAuthController) Get() {
	this.Data["IsForced"] = true
	this.TplName = "forced-auth.tpl"
}

func (this *ForcedAuthController) Post() {
	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	authCode := this.Input().Get("authCode")
	cardNumber := this.Input().Get("cardNumber")
	expirationDate := this.Input().Get("expirationDate")
	transactionAmount := this.Input().Get("amt")
	addr := this.Input().Get("addr")
	zip := this.Input().Get("zip")
	cardType := this.Input().Get("cardType")

	req := &ForcedAuth{
		Amt:            transactionAmount,
		ExpirationDate: expirationDate,
		CardNumber:     cardNumber,
		AuthCode:       authCode,
		MerchantNo:     merchantNo,
		StoreNo:        storeNo,
		Address:        addr,
		Zip:            zip,
		CardType:       cardType,
	}

	resp, err := req.PostToYuansfer()

	if err != nil {
		fmt.Println(err)
		this.Ctx.WriteString("something wrong happened")
	}

	t := template.New("forced-template")
	t, _ = t.Parse(resp)
	_ = t.Execute(this.Ctx.ResponseWriter, resp)

	return
}

func (f ForcedAuth) PostToYuansfer() (string, error) {

	values := generateValues(f, YuansferAPI.Token.SecurepayToken)
	forcedAuthUrl := YuansferHost + ForcedAuthUrl

	return postToYuansfer(forcedAuthUrl, values)
}
