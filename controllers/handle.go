package controllers

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
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
	tk := md5Token(token)
	buf := map2Str(data)
	buf.WriteString(tk)
	values.Add("verifySign", md5Token(buf.String()))

	for key, value := range data {
		if value != "" {
			values.Add(key, value)
		}
	}

	return values
}

func map2Str(m map[string]string) bytes.Buffer {
	var keys []string
	for k := range m {
		if m[k] != "" {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)

	var buf bytes.Buffer
	for _, key := range keys {
		buf.WriteString(key)
		buf.WriteString("=")
		buf.WriteString(m[key])
		buf.WriteString("&")
	}

	return buf
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
	tk := md5Token(token)
	buf := map2Str(m)
	buf.WriteString(tk)
	vs := md5Token(buf.String())

	return m, vs == verifySign
}

func parseSignature(obj signature) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Tag.Get("json")] = v.Field(i).String()
	}
	return data
}

func map2byte(m map[string]string) []byte {
	var keys []string
	for k := range m {
		if m[k] != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	des := ""
	for _, key := range keys {
		des += key + "=" + m[key] + "&"
	}
	b := []byte(des)[0 : len(des)-1]
	return b
}

func Sha1(data []byte) string {
	sha1 := sha1.New()
	sha1.Write(data)
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}
