package top

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Client struct {
	AppKey    string
	AppSecret string
	Param     map[string]interface{}
	IsHttps   bool
	IsSanBox  bool
}

const (
	OpenRouterHttpAPI        = "http://gw.api.taobao.com/router/rest"     //正式版API路由地址 http 版本
	OpenRouterHttpsAPI       = "https://eco.taobao.com/router/rest"       //正式版API路由地址 https 版本
	OpenSanBoxRouterHttpAPI  = "http://gw.api.tbsandbox.com/router/rest"  //沙盒版API路由地址 http 版本
	OpenSanBoxRouterHttpsAPI = "https://gw.api.tbsandbox.com/router/rest" //沙盒版API路由地址 https 版本
)

func NewClient(appkey, appsecret string) *Client {
	return &Client{
		AppKey:    appkey,
		AppSecret: appsecret,
	}
}

func (t *Client) Run(api TopApi) ([]byte, error) {
	if msg, ok := api.CheckParam(); !ok {
		return []byte(msg), nil
	}

	param := map[string]interface{}{
		"app_key":     t.AppKey,
		"sign_method": "md5",
		"format":      "json",
		"v":           "2.0",
		"timestamp":   time.Now().Format("2006-01-02 15:04:05"),
	}

	param["method"] = api.ApiName()
	for key, val := range api.Param() {
		param[key] = val
	}

	param["sign"] = sign(param, t.AppSecret)

	var routerApi string
	if t.IsSanBox {
		if t.IsHttps {
			routerApi = OpenSanBoxRouterHttpsAPI
		} else {
			routerApi = OpenSanBoxRouterHttpAPI
		}
	} else {
		if t.IsHttps {
			routerApi = OpenRouterHttpsAPI
		} else {
			routerApi = OpenRouterHttpAPI
		}
	}

	respByte, err := request(routerApi, param)

	return respByte, err
}

func sign(param map[string]interface{}, appSecret string) (str string) {
	keys := []string{}
	for key, _ := range param {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var buff bytes.Buffer
	buff.WriteString(appSecret)
	for i := 0; i < len(keys); i++ {
		buff.WriteString(fmt.Sprintf("%s%v", string(keys[i]), param[keys[i]]))
	}
	buff.WriteString(appSecret)

	s := md5.New()
	s.Write(buff.Bytes())

	return strings.ToUpper(hex.EncodeToString(s.Sum(nil)))
}

func request(url string, param map[string]interface{}) ([]byte, error) {
	var bodyBuff bytes.Buffer

	for k, v := range param {
		bodyBuff.WriteString(fmt.Sprintf("%s=%v&", k, v))
	}

	resp, err := http.Post(url, "application/x-www-form-urlencoded; charset=utf-8", bytes.NewReader(bodyBuff.Bytes()))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
