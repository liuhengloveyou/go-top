package top_test

import (
	"fmt"
	"testing"

	"github.com/liuhengloveyou/go-top"
)

const (
	AppKey    = "xxxxxx"
	AppSecret = "xxxxxx"
)

func TestTaobaoTbkTpwdCreate(t *testing.T) {
	var topClient *top.Client = top.NewClient(AppKey, AppSecret)

	var api = top.NewTaobaoTbkTpwdCreate()
	api.SetParam("text", "长度大于5个字符")
	api.SetParam("url", "https://uland.taobao.com/")

	d, err := topClient.Run(api)

	fmt.Println(">>>>>>", string(d), err)
}
