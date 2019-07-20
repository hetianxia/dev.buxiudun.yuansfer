package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

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
	IpnURL       string `json:"ipnUrl"`
	Reference    string `json:"reference"`
}

func (this *ScannerAddController) Get() {
	this.Data["IsInstoreAddScanner"] = true
	this.TplName = "instore-add-scanner.tpl"
}

func (this *ScannerAddController) Post() {

	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	storeAdminNo := this.Input().Get("storeAdminNo")
	ipnURL := this.Input().Get("ipnUrl")
	amt := this.Input().Get("amt")
	reference := this.Input().Get("reference")

	if "" == reference {
		reference = fmt.Sprintf("seq_%d", time.Now().Unix())
		reference = md5Token(reference)
	}

	req := InstoreAddScanner{
		MerchantNo:   merchantNo,
		StoreNo:      storeNo,
		Amount:       amt,
		StoreAdminNo: storeAdminNo,
		IpnURL:       ipnURL,
		Reference:    reference,
	}

	ret, err := req.PostToYuansfer()
	if err != nil {
		log.Println(err)
	}
	log.Println("instore add scanner result:", ret)
	var qr yuan.AddResponse

	err = json.Unmarshal([]byte(ret), &qr)
	if err != nil {
		log.Println("Unmarshal Err:%v", err)
	}
	resp := qr.Transaction
	log.Println("resp:", qr)
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
	fmt.Println("instore add scanner addUrl:", addUrl)
	values := generateValues(r, YuansferAPI.Token.InstoreToken)
	fmt.Println("values:", values)
	return postToYuansfer(addUrl, values)
}
