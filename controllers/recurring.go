package controllers

import (
	"fmt"
	"html/template"
)

const (
	RecurringPayUrl    = "creditpay/v2/credit-recurring"
	RecurringSearchUrl = "creditpay/v2/search-recurring"
	RecurringUpdateUrl = "creditpay/v2/update-recurring"
)

type RecurringController struct {
	Controller
}

type Recurring struct {
	ScheduleNo   string `json:"scheduleNo"`
	PaymentDate  string `json:"paymentDate"`
	MerchantNo   string `json:"merchantNo"`
	PaymentCount string `json:"paymentCount"`
	Status       string `json:"status"`
	StoreNo      string `json:"storeNo"`
	RetCode      string `json:"ret_code"`

	router string
}

func (this *RecurringController) Get() {
	this.Data["IsRecurring"] = true
	this.TplName = "recurring.tpl"
}

func (this *RecurringController) Pay() {
	merchantNo := this.Input().Get("merchantNo")
	storeNo := this.Input().Get("storeNo")
	scheduleNo := this.Input().Get("scheduleNo")
	paymentDate := this.Input().Get("paymentDate")

	req := &Recurring{
		ScheduleNo:  scheduleNo,
		PaymentDate: paymentDate,
		MerchantNo:  merchantNo,
		StoreNo:     storeNo,
		router:      "pay",
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
	var (
		recurringUrl string
	)
	if "pay" == r.router {
		recurringUrl = YuansferHost + RecurringPayUrl
	} else if "search" == r.router {
		recurringUrl = YuansferHost + RecurringSearchUrl
	} else if "update" == r.router {
		recurringUrl = YuansferHost + RecurringUpdateUrl
	}
	r.router = ""
	values := generateValues(r, YuansferAPI.Token.SecurepayToken)
	return postToYuansfer(recurringUrl, values)
}

func (r *RecurringController) Search() {
	merchantNo := r.Input().Get("merchantNo")
	storeNo := r.Input().Get("storeNo")
	scheduleNo := r.Input().Get("scheduleNo")

	req := &Recurring{
		ScheduleNo: scheduleNo,
		MerchantNo: merchantNo,
		StoreNo:    storeNo,
		router:     "search",
	}

	resp, err := req.PostToYuansfer()

	if err != nil {
		fmt.Println(err)
		r.Ctx.WriteString("something wrong happened")
	}

	t := template.New("recurring-template")
	t, _ = t.Parse(resp)
	_ = t.Execute(r.Ctx.ResponseWriter, resp)

	return
}

func (r *RecurringController) Update() {
	var (
		resp string
		err  error
	)
	merchantNo := r.Input().Get("merchantNo")
	storeNo := r.Input().Get("storeNo")
	scheduleNo := r.Input().Get("scheduleNo")
	status := r.Input().Get("stat")
	paymentCount := r.Input().Get("count")

	req := &Recurring{
		ScheduleNo:   scheduleNo,
		MerchantNo:   merchantNo,
		StoreNo:      storeNo,
		Status:       status,
		PaymentCount: paymentCount,
		router:       "update",
	}

	resp, err = req.PostToYuansfer()

	if err != nil {
		fmt.Println(err)
		r.Ctx.WriteString("something wrong happened")
	}

	t := template.New("recurring-template")
	t, _ = t.Parse(resp)
	_ = t.Execute(r.Ctx.ResponseWriter, resp)

	return
}
