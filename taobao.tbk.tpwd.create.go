package top

type TaobaoTbkTpwdCreate struct {
	param map[string]interface{}
}

func NewTaobaoTbkTpwdCreate() *TaobaoTbkTpwdCreate {
	return &TaobaoTbkTpwdCreate{
		param: map[string]interface{}{},
	}

}

func (t *TaobaoTbkTpwdCreate) ApiName() (s string) {
	return "taobao.tbk.tpwd.create"
}

func (t *TaobaoTbkTpwdCreate) SetParam(k string, v interface{}) {
	t.param[k] = v
}

func (t *TaobaoTbkTpwdCreate) Param() map[string]interface{} {
	return t.param
}

func (t *TaobaoTbkTpwdCreate) CheckParam() (msg string, ok bool) {
	if text, ok := t.param["text"]; !ok || len(text.(string)) <= 5 {
		return "Missing required arguments:text", false
	}

	if _, ok := t.param["url"]; !ok {
		return "Missing required arguments:url", false
	}

	return "", true
}
