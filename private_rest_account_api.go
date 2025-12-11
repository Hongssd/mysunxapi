package mysunxapi

// Account
// GET 查询账户余额
func (client *PrivateRestClient) NewPrivateRestAccountBalance() *PrivateRestAccountBalanceAPI {
	return &PrivateRestAccountBalanceAPI{
		client: client,
		req:    &PrivateRestAccountBalanceReq{},
	}
}

func (api *PrivateRestAccountBalanceAPI) Do() (*SunxRestRes[PrivateRestAccountBalanceRes], error) {
	url := sunxHandlerRequestAPIWithSignature(api.client.c, REST, GET, api.req, PrivateRestAPIMap[PrivateRestAccountBalance])
	return sunxCallApi[PrivateRestAccountBalanceRes](url, NIL_REQBODY, GET)
}

// POST 组合查询用户财务记录
func (client *PrivateRestClient) NewPrivateRestAccountBillRecord() *PrivateRestAccountBillRecordAPI {
	return &PrivateRestAccountBillRecordAPI{
		client: client,
		req:    &PrivateRestAccountBillRecordReq{},
	}
}

func (api *PrivateRestAccountBillRecordAPI) Do() (*SunxRestRes[PrivateRestAccountBillRecordRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestAccountBillRecordReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestAccountBillRecord])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestAccountBillRecordRes](url, reqBody, POST)
}

// POST 查询用户当前的手续费费率
func (client *PrivateRestClient) NewPrivateRestAccountFeeRate() *PrivateRestAccountFeeRateAPI {
	return &PrivateRestAccountFeeRateAPI{
		client: client,
		req:    &PrivateRestAccountFeeRateReq{},
	}
}

func (api *PrivateRestAccountFeeRateAPI) Do() (*SunxRestRes[PrivateRestAccountFeeRateRes], error) {
	url := sunxHandlerRequestAPIWithSignature[PrivateRestAccountFeeRateReq](api.client.c, REST, POST, nil, PrivateRestAPIMap[PrivateRestAccountFeeRate])
	reqBody, err := json.Marshal(api.req)
	if err != nil {
		return nil, err
	}
	return sunxCallApi[PrivateRestAccountFeeRateRes](url, reqBody, POST)
}
