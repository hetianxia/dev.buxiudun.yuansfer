package controllers

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Cache = redis.Conn

var (
	Conn          Cache
	WxAccessToken string
	WxJSAPITicket string
	AppID         string
	Secret        string

	UrlGetToken  = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	UrlGetTicket = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)

func init() {
	var err error
	if nil == Conn {
		Conn, err = redis.Dial("tcp", bxdConf.CacheAddr)
		if err != nil {
			log.Println("连接REDIS服务器失败:\n %v\n", err)
			os.Exit(99)
		}
		if _, err = Conn.Do("AUTH", bxdConf.CacheAuth); err != nil {
			Conn.Close()
			log.Println("密码错误! %v\n", err)
			os.Exit(99)
		}
		if nil == Conn {
			log.Println("Redis 初始化失败")
		} else {
			log.Println("Redis 初始化成功")
		}
	}

	if "" == WxJSAPITicket || "" == WxAccessToken {
		log.Println("Load configuration failed.")
		os.Exit(98)
	}
}

type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

func getToken() string {
	token, err := redis.String(Conn.Do("GET", WxAccessToken))
	if nil != err {
		if "redigo: nil returned" == err.Error() {
			token = requestToken()
			if "" != token {
				Conn.Do("SET", WxAccessToken, token)
				n, _ := Conn.Do("EXPIRE", WxAccessToken, 7000)
				if n == int64(1) {
					fmt.Println("success")
				} else {
					fmt.Println("set token expire time err")
				}
				return token
			}
		}
	}

	return token
}

func requestToken() string {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	urlGetToken := fmt.Sprintf(UrlGetToken, AppID, Secret)
	resp, err := client.Get(urlGetToken)
	if err != nil {
		fmt.Println("requestToken error:", err.Error())
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("requestToken result:", string(body))
	var accessToken AccessToken
	err = json.Unmarshal(body, &accessToken)
	if nil != err {
		fmt.Println("access token Unmarshal err:", err.Error())
		return ""
	}

	return accessToken.Token
}

type JSAPITicket struct {
	Code      int    `json:"errcode"`
	Msg       string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

func getTicket() string {
	ticket, err := redis.String(Conn.Do("GET", WxJSAPITicket))
	if nil != err {
		if "redigo: nil returned" == err.Error() {
			ticket = requestJSAPITicket()
			if "" != ticket {
				Conn.Do("SET", WxJSAPITicket, ticket)
				n, _ := Conn.Do("EXPIRE", WxJSAPITicket, 7000)
				if n == int64(1) {
					fmt.Println("success")
				} else {
					fmt.Println("set ticket expire time err")
				}
				return ticket
			}
		}
	}

	return ticket
}

func requestJSAPITicket() string {
	accessToken := getToken()
	if "" == accessToken {
		fmt.Println("Accesss Token not found")
		return ""
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	addr := fmt.Sprintf(UrlGetTicket, accessToken)
	resp, err := client.Get(addr)
	if err != nil {
		fmt.Println("requestTicket error:", err.Error())
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("requestTicket result:", string(body))
	var jsapiTicket JSAPITicket
	err = json.Unmarshal(body, &jsapiTicket)
	if nil != err {
		fmt.Println("JSAPITicket Unmarshal err:", err.Error())
		return ""
	}

	return jsapiTicket.Ticket
}

type signature struct {
	Timestamp   string `json:"timestamp"`
	NonceStr    string `json:"noncestr"`
	Signature   string `json:"signature"`
	JSAPITicket string `json:"jsapi_ticket"`
	Url         string `json:"url"`

	AppId string `json:"appid"`
}

func genSignature(addr string) string {
	ticket := getTicket()
	if "" == ticket {
		fmt.Println("Ticket not found.")
		return ""
	}
	timestamp := time.Now().Unix()
	s := signature{
		Timestamp:   fmt.Sprintf("%d", timestamp),
		NonceStr:    fmt.Sprintf("seq_%d", timestamp),
		JSAPITicket: ticket,
		Url:         addr,
	}
	m := parseSignature(s)
	if nil == m {
		fmt.Println("gen Map error")
		return ""
	}
	b := map2byte(m)
	sign := Sha1(b)
	fmt.Println("signature:", sign)
	s.Signature = sign
	s.Url = addr
	s.JSAPITicket = ""
	s.AppId = AppID
	data, err := json.Marshal(&s)
	if nil != err {
		fmt.Println("gen JSON data err:", err.Error())
		return ""
	}
	fmt.Println("response:", string(data))

	return string(data)
}
