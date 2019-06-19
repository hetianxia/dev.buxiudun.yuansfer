package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"

	yuan "github.com/yuansfer/golang_sdk"
)

var (
	YuansferAPI  = yuan.YuansferAPI
	YuansferHost = YuansferAPI.Host[0]
)

func struct2Map(obj yuan.Yuansfer) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Tag.Get("json")] = v.Field(i).String()
	}
	return data
}

func md5Token(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	md5Data := md5.Sum([]byte(""))
	return hex.EncodeToString(md5Data)
}

func generateValues(req yuan.Yuansfer, token string) url.Values {
	values := url.Values{}
	data := struct2Map(req)
	pre := map2Str(data) + md5Token(token)
	values.Add("verifySign", md5Token(pre))

	for key, value := range data {
		if value != "" {
			values.Add(key, value)
		}
	}

	return values
}

func map2Str(m map[string]string) string {
	var keys []string
	for k := range m {
		if m[k] != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	dec := ""
	for _, key := range keys {
		dec += key + "=" + m[key] + "&"
	}

	return dec
}

func postToYuansfer(addr string, values url.Values) (string, error) {
	var (
		err  error
		resp *http.Response
	)
	resp, err = http.PostForm(addr, values)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), err
}

func values2Map(m url.Values) map[string]string {
	var r = make(map[string]string)
	m.Del("verifySign")
	for k := range m {
		r[k] = m[k][0]
	}
	return r
}

//VerifySignNotify checks the parameters from Yuansfer with the value of verifySign.
func VerifySignNotify(values url.Values, token string) (m map[string]string, r bool) {
	verifySign := values.Get("verifySign")
	m = values2Map(values)
	pre := map2Str(m) + md5Token(token)
	vs := md5Token(pre)

	return m, vs == verifySign
}
