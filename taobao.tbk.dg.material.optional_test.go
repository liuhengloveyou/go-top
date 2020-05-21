package top_test

import (
	"fmt"
	"testing"

	"github.com/liuhengloveyou/go-top"
)

func TestTaobaoTbkDgMaterialOptional(t *testing.T) {
	var topClient *top.Client = top.NewClient(AppKey, AppSecret)

	var api = top.NewTaobaoTbkDgMaterialOptional()
	api.SetParam("adzone_id", "xxxxxx")
	api.SetParam("site_id", "xxxxxx")
	api.SetParam("q", "格调华夫饼整箱早餐原味蛋糕")

	d, err := topClient.Run(api)

	fmt.Println(">>>>>>", string(d), err)

}
