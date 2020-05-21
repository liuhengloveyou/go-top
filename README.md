golang版的淘宝客(tbk)SDK



## API示例

### [taobao.tbk.dg.material.optional( 淘宝客-推广者-物料搜索 )](https://open.taobao.com/api.htm?docId=35896&docType=2)

```go
func TestTaobaoTbkDgMaterialOptional(t *testing.T) {
	var topClient *top.Client = top.NewClient(AppKey, AppSecret)

	var api = top.NewTaobaoTbkDgMaterialOptional()
	api.SetParam("adzone_id", "xxxxxx")
	api.SetParam("site_id", "xxxxxx")
	api.SetParam("q", "格调华夫饼整箱早餐原味蛋糕")

	d, err := topClient.Run(api)

	fmt.Println(">>>>>>", string(d), err)
}
```



###  [taobao.tbk.tpwd.create( 淘宝客-公用-淘口令生成 )](https://open.taobao.com/api.htm?docId=31127&docType=2)

```go
func TestTaobaoTbkTpwdCreate(t *testing.T) {
	var topClient *top.Client = top.NewClient(AppKey, AppSecret)

	var api = top.NewTaobaoTbkTpwdCreate()
	api.SetParam("text", "长度大于5个字符")
	api.SetParam("url", "https://uland.taobao.com/")

	d, err := topClient.Run(api)

	fmt.Println(">>>>>>", string(d), err)
}
```