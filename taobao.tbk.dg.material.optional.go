package top

type TaobaoTbkDgMaterialOptional struct {
	param map[string]interface{}
}

func NewTaobaoTbkDgMaterialOptional() *TaobaoTbkDgMaterialOptional {
	return &TaobaoTbkDgMaterialOptional{
		param: map[string]interface{}{
			"platform":      2,                                                            // 链接形式：1：PC，2：无线，默认：１
			"sort":          "tk_total_commi",                                             //排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price）
			"has_coupon":    true,                                                         // 优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
			"fields":        "user_id,shop_title,shop_type,seller_nick,pict_url,shop_url", //需返回的字段列表
			"start_tk_rate": 1,                                                            //淘客佣金比率下限，1~10000
		}}

}

func (t *TaobaoTbkDgMaterialOptional) ApiName() (s string) {
	return "taobao.tbk.dg.material.optional"
}

func (t *TaobaoTbkDgMaterialOptional) SetParam(k string, v interface{}) {
	t.param[k] = v
}

func (t *TaobaoTbkDgMaterialOptional) Param() map[string]interface{} {
	return t.param
}

func (t *TaobaoTbkDgMaterialOptional) CheckParam() (msg string, ok bool) {
	if _, ok := t.param["adzone_id"]; !ok {
		return "Missing required arguments:adzone_id", false
	}

	if _, ok := t.param["site_id"]; !ok {
		return "Missing required arguments:site_id", false
	}

	return "", true
}
