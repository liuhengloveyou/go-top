package top

type TaobaoTbkDgMaterialOptional struct {
	param map[string]interface{}
}

func NewTaobaoTbkDgMaterialOptional() *TaobaoTbkDgMaterialOptional {
	return &TaobaoTbkDgMaterialOptional{
		param: map[string]interface{}{
			"fields":                "user_id,shop_title,shop_type,seller_nick,pict_url,shop_url", //需返回的字段列表
			"sort":                  "commission_rate_des",                                        //排序_des（降序），排序_asc（升序），佣金比率（commission_rate）， 商品数量（auction_count），销售总数量（total_auction）
			"start_credit":          1,                                                            //信用等级下限，1~20
			"end_credit":            20,                                                           //信用等级上限，1~20
			"start_commission_rate": 1,                                                            //淘客佣金比率下限，1~10000
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
