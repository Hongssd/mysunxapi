package mysunxapi

type PublicRestPublicContractInfoReq struct {
	ContractCode      *string `json:"contract_code"`       // 合约代码
	SupportMarginMode *string `json:"support_margin_mode"` // 合约支持的保证金模式
	Pair              *string `json:"pair"`                // 交易对
	ContractType      *string `json:"contract_type"`       // 合约类型  swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
	BusinessType      *string `json:"business_type"`       // 业务类型 futures：交割、swap：永续、all：全部
}

type PublicRestPublicContractInfoAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicContractInfoReq
}

// string false 合约代码 永续："BTC-USDT"... ，交割："BTC-USDT-210625"...
func (api *PublicRestPublicContractInfoAPI) ContractCode(contractCode string) *PublicRestPublicContractInfoAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// string false 合约支持的保证金模式 cross：仅支持全仓模式；isolated：仅支持逐仓模式；all：全逐仓都支持
func (api *PublicRestPublicContractInfoAPI) SupportMarginMode(supportMarginMode string) *PublicRestPublicContractInfoAPI {
	api.req.SupportMarginMode = GetPointer(supportMarginMode)
	return api
}

// string false 交易对 BTC-USDT
func (api *PublicRestPublicContractInfoAPI) Pair(pair string) *PublicRestPublicContractInfoAPI {
	api.req.Pair = GetPointer(pair)
	return api
}

// string false 合约类型  swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
func (api *PublicRestPublicContractInfoAPI) ContractType(contractType string) *PublicRestPublicContractInfoAPI {
	api.req.ContractType = GetPointer(contractType)
	return api
}

// string false 业务类型 futures：交割、swap：永续、all：全部
func (api *PublicRestPublicContractInfoAPI) BusinessType(businessType string) *PublicRestPublicContractInfoAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

type PublicRestPublicIndexReq struct {
	ContractCode *string `json:"contract_code"`
}

type PublicRestPublicIndexAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicIndexReq
}

// string false 指数代码 BTC-USDT,"ETH-USDT"...
func (api *PublicRestPublicIndexAPI) ContractCode(contractCode string) *PublicRestPublicIndexAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

type PublicRestPublicRiskLimitReq struct {
	ContractCode *string `json:"contract_code"`
	Pair         *string `json:"pair"`
	ContractType *string `json:"contract_type"`
	BusinessType *string `json:"business_type"`
}

type PublicRestPublicRiskLimitAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicRiskLimitReq
}

// contract_code string false 合约代码 永续："BTC-USDT"... ，交割："BTC-USDT-210625"...
func (api *PublicRestPublicRiskLimitAPI) ContractCode(contractCode string) *PublicRestPublicRiskLimitAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// pair string false 交易对 BTC-USDT
func (api *PublicRestPublicRiskLimitAPI) Pair(pair string) *PublicRestPublicRiskLimitAPI {
	api.req.Pair = GetPointer(pair)
	return api
}

// contract_type string false 合约类型 swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
func (api *PublicRestPublicRiskLimitAPI) ContractType(contractType string) *PublicRestPublicRiskLimitAPI {
	api.req.ContractType = GetPointer(contractType)
	return api
}

// business_type string false 业务类型，不填默认永续 futures：交割、swap：永续、all：全部
func (api *PublicRestPublicRiskLimitAPI) BusinessType(businessType string) *PublicRestPublicRiskLimitAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

type PublicRestPublicFundingRateReq struct {
	ContractCode *string `json:"contract_code"`
}

type PublicRestPublicFundingRateAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicFundingRateReq
}

// contract_code string true 合约代码 永续："BTC-USDT"... ，交割："BTC-USDT-210625"...
func (api *PublicRestPublicFundingRateAPI) ContractCode(contractCode string) *PublicRestPublicFundingRateAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

type PublicRestPublicFundingRateHistoryReq struct {
	PageIndex    *int    `json:"page_index"`
	PageSize     *int    `json:"page_size"`
	ContractCode *string `json:"contract_code"`
}
type PublicRestPublicFundingRateHistoryAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicFundingRateHistoryReq
}

// page_index int false 页码，不填默认第1页
func (api *PublicRestPublicFundingRateHistoryAPI) PageIndex(pageIndex int) *PublicRestPublicFundingRateHistoryAPI {
	api.req.PageIndex = GetPointer(pageIndex)
	return api
}

// page_size int false 不填默认20，不得多于50
func (api *PublicRestPublicFundingRateHistoryAPI) PageSize(pageSize int) *PublicRestPublicFundingRateHistoryAPI {
	api.req.PageSize = GetPointer(pageSize)
	return api
}

// contract_code string true 合约代码 BTC-USDT等
func (api *PublicRestPublicFundingRateHistoryAPI) ContractCode(contractCode string) *PublicRestPublicFundingRateHistoryAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

type PublicRestPublicPriceLimitReq struct {
	ContractCode *string `json:"contract_code"` // false 合约代码
	Pair         *string `json:"pair"`          // false 交易对
	ContractType *string `json:"contract_type"` // false 合约类型
	BusinessType *string `json:"business_type"` // false 业务类型，不填默认永续
}
type PublicRestPublicPriceLimitAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicPriceLimitReq
}

// contract_code string false 合约代码
func (api *PublicRestPublicPriceLimitAPI) ContractCode(contractCode string) *PublicRestPublicPriceLimitAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// pair string false 交易对
func (api *PublicRestPublicPriceLimitAPI) Pair(pair string) *PublicRestPublicPriceLimitAPI {
	api.req.Pair = GetPointer(pair)
	return api
}

// contract_type string false 合约类型
func (api *PublicRestPublicPriceLimitAPI) ContractType(contractType string) *PublicRestPublicPriceLimitAPI {
	api.req.ContractType = GetPointer(contractType)
	return api
}

// business_type string false 业务类型，不填默认永续
func (api *PublicRestPublicPriceLimitAPI) BusinessType(businessType string) *PublicRestPublicPriceLimitAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

type PublicRestPublicMultiAssetsMarginReq struct{}
type PublicRestPublicMultiAssetsMarginAPI struct {
	client *PublicRestClient
	req    *PublicRestPublicMultiAssetsMarginReq
}
