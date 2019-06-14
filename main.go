package main

import (
	"github.com/astaxie/beego"
	"github.com/hetianxia/buxiudun/controllers"
)

func main() {
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.Graceful = true

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/pay", &controllers.HomeController{})
	beego.Router("/inquire", &controllers.InquireController{})
	beego.Router("/recurring", &controllers.RecurringController{})
	beego.Router("/online-reverse", &controllers.ReverseController{})
	// beego.Router("/exchange-rate", &controllers.ExchangeRateController{})

	beego.Router("/callback", &controllers.CallbackController{})

	// In-store
	//
	beego.Router("/instore-add", &controllers.InstoreAddController{})
	beego.Router("/instore-pay", &controllers.InstorePayController{})
	beego.Router("/instore-qrcode", &controllers.InstoreQrcodeController{})
	beego.Router("/instore-reverse", &controllers.InstoreReverseController{})
	beego.Router("/micropay", &controllers.MicropayController{})
	beego.Router("/refund", &controllers.RefundController{})

	beego.Run()
}
