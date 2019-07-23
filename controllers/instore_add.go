package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	yuan "github.com/yuansfer/golang_sdk"
)

type InstoreAddController struct {
	Controller
}

func (this *InstoreAddController) Get() {
	reference := fmt.Sprintf("seq_%d", time.Now().Unix())
	reference = md5Token(reference)
	this.Data["reference"] = reference
	this.Data["IsInstoreAdd"] = true
	this.TplName = "instore-add.tpl"
}

func (this *InstoreAddController) Post() {

	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	storeAdminNo := this.Input().Get("storeAdminNo")
	reference := this.Input().Get("reference")
	if "" == reference {
		reference = fmt.Sprintf("SEQ_%d", time.Now().Unix())
		reference = md5Token(reference)
	}
	amt := this.Input().Get("amt")

	req := yuan.InstoreAdd{
		MerchantNo:   merchantNo,
		StoreNo:      storeNo,
		Amount:       amt,
		StoreAdminNo: storeAdminNo,
		Reference:    reference,
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
