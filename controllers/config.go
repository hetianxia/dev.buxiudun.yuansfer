package controllers

import (
	// "flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
)

var (
	//Development or Production of Yuansfer Service Address
	yuansferHost string
	//YuansferAPI is the configuration information
	bxdConf yuansferAPI
)

type yuansferAPI struct {
	Host                []string      `toml:"yuansfer_host"`
	OnlinePayment       string        `toml:"online_payment_url"`
	OnlineQuery         string        `toml:"online_query_url"`
	OnlineRefund        string        `toml:"online_refund_url"`
	InstoreAdd          string        `toml:"instore_add_url"`
	InstorePay          string        `toml:"instore_pay_url"`
	InstoreCreateQrcode string        `toml:"instore_create_qrcode"`
	InstoreQuery        string        `toml:"instore_query_url"`
	InstoreRefund       string        `toml:"instore_refund_url"`
	InstoreReverse      string        `toml:"instore_reverse_url"`
	Micropay            string        `toml:"micropay_url"`
	PwdPre              string        `toml:"password_prefix"`
	Token               yuansferToken `toml:"token"`
	CacheAddr           string        `toml:"redis_addr"`
	CacheAuth           string        `toml:"redis_auth"`
	WxAccessTokenKey    string        `toml:"wx_token_key"`
	WxJSAPITicketKey    string        `toml:"wx_ticket_key"`
	AppID               string        `toml:"appid"`
	Secret              string        `toml:"secret"`
}

type yuansferToken struct {
	SecurepayToken string `yaml:"online_token" toml:"online_token"`
	InstoreToken   string `yaml:"instore_token" toml:"instore_token"`
	MicropayToken  string `yaml:"micropay_token" toml:"micropay_token"`
}

const (
	//ConfigFile is the default configuration file name.
	configFile = "./config.toml"
)

func init() {
	env := "dev"
	s := strings.Split(configFile, ".")
	fileType := s[len(s)-1]

	switch fileType {
	case "toml":
		if _, err := toml.DecodeFile(configFile, &bxdConf); err != nil {
			log.Fatalf("Decode toml err:%s", err.Error())
		}
		fmt.Println("bxdConf:", bxdConf)
	case "yml", "yaml":
		data, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Fatalf("read config file %s error:%s", configFile, err.Error())
		}
		err = yaml.Unmarshal([]byte(data), &bxdConf)
		if err != nil {
			log.Fatalf("Unmarshal config file %s error:%s", configFile, err.Error())
		}
	default:
		log.Fatal("Unknown configuration file type.")
	}

	yuansferHost = map[string]string{
		"dev":     bxdConf.Host[0],
		"product": bxdConf.Host[1],
	}[env]

	WxAccessToken = bxdConf.WxAccessTokenKey
	WxJSAPITicket = bxdConf.WxJSAPITicketKey
	AppID = bxdConf.AppID
	Secret = bxdConf.Secret
}
