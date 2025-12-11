package mysunxapi

// GET 查询当前持仓
func (client *PrivateRestClient) NewPrivateRestTradePositionOpens() *PrivateRestTradePositionOpensAPI {
	return &PrivateRestTradePositionOpensAPI{
		client: client,
		req:    &PrivateRestTradePositionOpensReq{},
	}
}

func (api *PrivateRestTradePositionOpensAPI) Do() (*SunxRestRes[PrivateRestTradePositionOpensRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestTradePositionOpens])
	return sunxCallApi[PrivateRestTradePositionOpensRes](url, NIL_REQBODY, GET)
}

// GET 查询杠杆等级列表
func (client *PrivateRestClient) NewPrivateRestTradePositionLeverGet() *PrivateRestTradePositionLeverGetAPI {
	return &PrivateRestTradePositionLeverGetAPI{
		client: client,
		req:    &PrivateRestTradePositionLeverGetReq{},
	}
}

func (api *PrivateRestTradePositionLeverGetAPI) Do() (*SunxRestRes[PrivateRestTradePositionLeverGetRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestTradePositionLeverGet])
	return sunxCallApi[PrivateRestTradePositionLeverGetRes](url, NIL_REQBODY, GET)
}

// POST 设置杠杆等级
func (client *PrivateRestClient) NewPrivateRestTradePositionLeverPost() *PrivateRestTradePositionLeverPostAPI {
	return &PrivateRestTradePositionLeverPostAPI{
		client: client,
		req:    &PrivateRestTradePositionLeverPostReq{},
	}
}

func (api *PrivateRestTradePositionLeverPostAPI) Do() (*SunxRestRes[PrivateRestTradePositionLeverPostRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestTradePositionLeverPostReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestTradePositionLeverPost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestTradePositionLeverPostRes](url, reqBody, POST)
}

// GET 查询持仓模式
func (client *PrivateRestClient) NewPrivateRestTradePositionModeGet() *PrivateRestTradePositionModeGetAPI {
	return &PrivateRestTradePositionModeGetAPI{
		client: client,
		req:    &PrivateRestTradePositionModeGetReq{},
	}
}

func (api *PrivateRestTradePositionModeGetAPI) Do() (*SunxRestRes[PrivateRestTradePositionModeGetRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestTradePositionModeGet])
	return sunxCallApi[PrivateRestTradePositionModeGetRes](url, NIL_REQBODY, GET)
}

// POST 设置持仓模式
func (client *PrivateRestClient) NewPrivateRestTradePositionModePost() *PrivateRestTradePositionModePostAPI {
	return &PrivateRestTradePositionModePostAPI{
		client: client,
		req:    &PrivateRestTradePositionModePostReq{},
	}
}

func (api *PrivateRestTradePositionModePostAPI) Do() (*SunxRestRes[PrivateRestTradePositionModePostRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestTradePositionModePostReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestTradePositionModePost])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestTradePositionModePostRes](url, reqBody, POST)
}

// GET 查询持仓风险限额
func (client *PrivateRestClient) NewPrivateRestTradePositionRiskLimit() *PrivateRestTradePositionRiskLimitAPI {
	return &PrivateRestTradePositionRiskLimitAPI{
		client: client,
		req:    &PrivateRestTradePositionRiskLimitReq{},
	}
}

func (api *PrivateRestTradePositionRiskLimitAPI) Do() (*SunxRestRes[PrivateRestTradePositionRiskLimitRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestTradePositionRiskLimit])
	return sunxCallApi[PrivateRestTradePositionRiskLimitRes](url, NIL_REQBODY, GET)
}

// POST 用户持仓量限制的查询
func (client *PrivateRestClient) NewPrivateRestTradePositionPositionLimit() *PrivateRestTradePositionPositionLimitAPI {
	return &PrivateRestTradePositionPositionLimitAPI{
		client: client,
		req:    &PrivateRestTradePositionPositionLimitReq{},
	}
}

func (api *PrivateRestTradePositionPositionLimitAPI) Do() (*SunxRestRes[PrivateRestTradePositionPositionLimitRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestTradePositionPositionLimitReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestTradePositionPositionLimit])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestTradePositionPositionLimitRes](url, reqBody, POST)
}