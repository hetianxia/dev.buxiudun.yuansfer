package controllers

import (
	"encoding/json"
	"log"

	yuan "github.com/yuansfer/golang_sdk"
)

const (
	instoreAddScanner = "app-instore/v2/add-external"
)

type ScannerAddController struct {
	Controller
}

type InstoreAddScanner struct {
	MerchantNo   string `json:"merchantNo"`
	StoreNo      string `json:"storeNo"`
	VerifySign   string `json:"verifySign"`
	Amount       string `json:"amount"`
	StoreAdminNo string `json:"storeAdminNo"`
}

func (this *ScannerAddController) Get() {
	this.Data["IsInstoreAddScanner"] = true
	this.TplName = "instore-add-scanner.tpl"
}

func (this *ScannerAddController) Post() {

	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	storeAdminNo := this.Input().Get("storeAdminNo")

	amt := this.Input().Get("amt")

	req := InstoreAddScanner{
		MerchantNo:   merchantNo,
		StoreNo:      storeNo,
		Amount:       amt,
		StoreAdminNo: storeAdminNo,
	}

	ret, err := req.PostToYuansfer()
	if err != nil {
		log.Println(err)
	}
	log.Println(ret)
	var qr yuan.AddResponse

	err = json.Unmarshal([]byte(ret), &qr)
	if err != nil {
		log.Println("Unmarshal Err:%v", err)
	}
	log.Println("resp:", qr)
	resp := qr.Transaction
	this.Data["IsInstoreAdd"] = true
	this.TplName = "instore-add-ret.tpl"

	this.checkData("TransactionNo", resp.TransactionNo)
	this.checkData("OriginalTransactionNo", resp.OriginalTransactionNo)
	this.checkData("Amount", resp.Amount)
	this.checkData("Vendor", resp.Vendor)
	this.checkData("MerchantNo", resp.MerchantNo)

	return
}

func (r InstoreAddScanner) PostToYuansfer() (string, error) {
	var (
		addUrl string
	)

	addUrl = YuansferHost + instoreAddScanner

	values := generateValues(r, YuansferAPI.Token.InstoreToken)
	return postToYuansfer(addUrl, values)
}
