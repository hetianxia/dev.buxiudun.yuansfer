package controllers

import (
	"fmt"
	"html/template"
)

const (
	RecurringUrl = "creditpay/v2/credit-recurring"
)

type RecurringController struct {
	Controller
}

type Recurring struct {
	ScheduleNo  string `json:"scheduleNo"`
	PaymentDate string `json:"paymentDate"`
	MerchantNo  string `json:"merchantNo"`
	StoreNo     string `json:"storeNo"`
	RetCode     string `json:"ret_code"`
}

func (this *RecurringController) Get() {
	this.Data["IsRecurring"] = true
	this.TplName = "recurring.tpl"
}

func (this *RecurringController) Post() {
	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	scheduleNo := this.Input().Get("scheduleNo")
	paymentDate := this.Input().Get("paymentDate")

	req := &Recurring{
		ScheduleNo:  scheduleNo,
		PaymentDate: paymentDate,
		MerchantNo:  merchantNo,
		StoreNo:     storeNo,
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

func (r Recurring) PostToYuansfer() (string, error) {

	values := generateValues(r, YuansferAPI.Token.SecurepayToken)
	recurringUrl := YuansferHost + RecurringUrl

	return postToYuansfer(recurringUrl, values)
}
