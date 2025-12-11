package mysunxapi

// GET 获取合约信息
func (client *PublicRestClient) NewPublicRestPublicContractInfo() *PublicRestPublicContractInfoAPI {
	return &PublicRestPublicContractInfoAPI{
		client: client,
		req:    &PublicRestPublicContractInfoReq{},
	}
}

func (api *PublicRestPublicContractInfoAPI) Do() (*SunxRestRes[PublicRestPublicContractInfoRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(REST, api.req, PublicRestAPIMap[PublicRestPublicContractInfo])
	return sunxCallApi[PublicRestPublicContractInfoRes](url, NIL_REQBODY, GET)
}

// GET 获取合约指数信息
func (client *PublicRestClient) NewPublicRestPublicIndex() *PublicRestPublicIndexAPI {
	return &PublicRestPublicIndexAPI{
		client: client,
		req:    &PublicRestPublicIndexReq{},
	}
}

func (api *PublicRestPublicIndexAPI) Do() (*SunxRestRes[PublicRestPublicIndexRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(REST, api.req, PublicRestAPIMap[PublicRestPublicIndex])
	return sunxCallApi[PublicRestPublicIndexRes](url, NIL_REQBODY, GET)
}

// GET 获取合约最高限价和最低限价
func (client *PublicRestClient) NewPublicRestPublicRiskLimit() *PublicRestPublicRiskLimitAPI {
	return &PublicRestPublicRiskLimitAPI{
		client: client,
		req:    &PublicRestPublicRiskLimitReq{},
	}
}

func (api *PublicRestPublicRiskLimitAPI) Do() (*SunxRestRes[PublicRestPublicRiskLimitRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(REST, api.req, PublicRestAPIMap[PublicRestPublicRiskLimit])
	return sunxCallApi[PublicRestPublicRiskLimitRes](url, NIL_REQBODY, GET)
}

// GET 获取合约的资金费率
func (client *PublicRestClient) NewPublicRestPublicFundingRate() *PublicRestPublicFundingRateAPI {
	return &PublicRestPublicFundingRateAPI{
		client: client,
		req:    &PublicRestPublicFundingRateReq{},
	}
}

func (api *PublicRestPublicFundingRateAPI) Do() (*SunxRestRes[PublicRestPublicFundingRateRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(REST, api.req, PublicRestAPIMap[PublicRestPublicFundingRate])
	return sunxCallApi[PublicRestPublicFundingRateRes](url, NIL_REQBODY, GET)
}

// GET 获取合约的历史资金费率
func (client *PublicRestClient) NewPublicRestPublicFundingRateHistory() *PublicRestPublicFundingRateHistoryAPI {
	return &PublicRestPublicFundingRateHistoryAPI{
		client: client,
		req:    &PublicRestPublicFundingRateHistoryReq{},
	}
}

// GET 获取合约最高限价和最低限价
func (client *PublicRestClient) NewPublicRestPublicPriceLimit() *PublicRestPublicPriceLimitAPI {
	return &PublicRestPublicPriceLimitAPI{
		client: client,
		req:    &PublicRestPublicPriceLimitReq{},
	}
}

func (api *PublicRestPublicPriceLimitAPI) Do() (*SunxRestRes[PublicRestPublicPriceLimitRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(REST, api.req, PublicRestAPIMap[PublicRestPublicPriceLimit])
	return sunxCallApi[PublicRestPublicPriceLimitRes](url, NIL_REQBODY, GET)
}

func (api *PublicRestPublicFundingRateHistoryAPI) Do() (*SunxRestRes[PublicRestPublicFundingRateHistoryRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(REST, api.req, PublicRestAPIMap[PublicRestPublicFundingRateHistory])
	return sunxCallApi[PublicRestPublicFundingRateHistoryRes](url, NIL_REQBODY, GET)
}

// GET 查询联合保证金支持币种
func (client *PublicRestClient) NewPublicRestPublicMultiAssetsMargin() *PublicRestPublicMultiAssetsMarginAPI {
	return &PublicRestPublicMultiAssetsMarginAPI{
		client: client,
		req:    &PublicRestPublicMultiAssetsMarginReq{},
	}
}

func (api *PublicRestPublicMultiAssetsMarginAPI) Do() (*SunxRestRes[PublicRestPublicMultiAssetsMarginRes], error) {
	url := sunxHandlerRequestAPIWithoutSignature(REST, api.req, PublicRestAPIMap[PublicRestPublicMultiAssetsMargin])
	return sunxCallApi[PublicRestPublicMultiAssetsMarginRes](url, NIL_REQBODY, GET)
}
