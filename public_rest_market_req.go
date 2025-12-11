package mysunxapi

type PublicRestMarketDepthReq struct {
	ContractCode *string `json:"contract_code"`
	Type         *string `json:"type"`
}
type PublicRestMarketDepthAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketDepthReq
}

// contract_code string false 合约代码 或 合约标识 永续："BTC-USDT" ... ，交割：“BTC-USDT-210625”... 或 BTC-USDT-CW（当周合约标识）、BTC-USDT-NW（次周合约标识）、BTC-USDT-CQ（当季合约标识）、BTC-USDT-NQ（次季合约标识）
func (api *PublicRestMarketDepthAPI) ContractCode(contractCode string) *PublicRestMarketDepthAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// type string true 深度类型 (150档数据) step0, step1, step2, step3, step4, step5, step14, step15, step16, step17（合并深度1-5,14-17）；step0时，不合并深度, (20档数据) step6, step7, step8, step9, step10, step11, step12, step13, step18, step19（合并深度7-13,18-19）；step6时，不合并深度
func (api *PublicRestMarketDepthAPI) Type(t string) *PublicRestMarketDepthAPI {
	api.req.Type = GetPointer(t)
	return api
}

type PublicRestMarketHistoryKlineReq struct {
	ContractCode *string `json:"contract_code"`
	Period       *string `json:"period"`
	Size         *int    `json:"size"`
	From         *int64  `json:"from"`
	To           *int64  `json:"to"`
}
type PublicRestMarketHistoryKlineAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketHistoryKlineReq
}

// contract_code string false 合约代码 或 合约标识 永续："BTC-USDT" ... ，交割：“BTC-USDT-210625”... 或 BTC-USDT-CW（当周合约标识）、BTC-USDT-NW（次周合约标识）、BTC-USDT-CQ（当季合约标识）、BTC-USDT-NQ（次季合约标识）
func (api *PublicRestMarketHistoryKlineAPI) ContractCode(contractCode string) *PublicRestMarketHistoryKlineAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// period string true K线类型
func (api *PublicRestMarketHistoryKlineAPI) Period(period string) *PublicRestMarketHistoryKlineAPI {
	api.req.Period = GetPointer(period)
	return api
}

// size int false 获取数量，默认150 [1,2000]
func (api *PublicRestMarketHistoryKlineAPI) Size(size int) *PublicRestMarketHistoryKlineAPI {
	api.req.Size = GetPointer(size)
	return api
}

// from long false 开始时间戳 10位 单位S
func (api *PublicRestMarketHistoryKlineAPI) From(from int64) *PublicRestMarketHistoryKlineAPI {
	api.req.From = GetPointer(from)
	return api
}

// to long false 结束时间戳 10位 单位S
func (api *PublicRestMarketHistoryKlineAPI) To(to int64) *PublicRestMarketHistoryKlineAPI {
	api.req.To = GetPointer(to)
	return api
}

type PublicRestMarketDetailMergedReq struct {
	ContractCode *string `json:"contract_code"`
}
type PublicRestMarketDetailMergedAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketDetailMergedReq
}

// contract_code string false 合约代码 或 合约标识 永续："BTC-USDT" ... ，交割：“BTC-USDT-210625”... 或 BTC-USDT-CW（当周合约标识）、BTC-USDT-NW（次周合约标识）、BTC-USDT-CQ（当季合约标识）、BTC-USDT-NQ（次季合约标识）
func (api *PublicRestMarketDetailMergedAPI) ContractCode(contractCode string) *PublicRestMarketDetailMergedAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

type PublicRestMarketTradeReq struct {
	ContractCode *string `json:"contract_code"`
	BusinessType *string `json:"business_type"`
}
type PublicRestMarketTradeAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketTradeReq
}

// contract_code string false 合约代码 或 合约标识 永续："BTC-USDT" ... ，交割：“BTC-USDT-210625”... 或 BTC-USDT-CW（当周合约标识）、BTC-USDT-NW（次周合约标识）、BTC-USDT-CQ（当季合约标识）、BTC-USDT-NQ（次季合约标识）
func (api *PublicRestMarketTradeAPI) ContractCode(contractCode string) *PublicRestMarketTradeAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// business_type string false 业务类型 futures：交割、swap：永续、all：全部
func (api *PublicRestMarketTradeAPI) BusinessType(businessType string) *PublicRestMarketTradeAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

type PublicRestMarketBBOReq struct {
	ContractCode *string `json:"contract_code"`
	BusinessType *string `json:"business_type"`
}
type PublicRestMarketBBOAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketBBOReq
}

// contract_code string false 合约代码 或 合约标识 永续："BTC-USDT" ... ，交割：“BTC-USDT-210625”... 或 BTC-USDT-CW（当周合约标识）、BTC-USDT-NW（次周合约标识）、BTC-USDT-CQ（当季合约标识）、BTC-USDT-NQ（次季合约标识）
func (api *PublicRestMarketBBOAPI) ContractCode(contractCode string) *PublicRestMarketBBOAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// business_type string false 业务类型 futures：交割、swap：永续、all：全部
func (api *PublicRestMarketBBOAPI) BusinessType(businessType string) *PublicRestMarketBBOAPI {
	api.req.BusinessType = GetPointer(businessType)
	return api
}

type PublicRestMarketHistoryTradeReq struct {
	ContractCode *string `json:"contract_code"`
	Size         *int    `json:"size"`
}
type PublicRestMarketHistoryTradeAPI struct {
	client *PublicRestClient
	req    *PublicRestMarketHistoryTradeReq
}

// contract_code string true 合约代码 或 合约标识 永续："BTC-USDT" ... ，交割：“BTC-USDT-210625”... 或 BTC-USDT-CW（当周合约标识）、BTC-USDT-NW（次周合约标识）、BTC-USDT-CQ（当季合约标识）、BTC-USDT-NQ（次季合约标识）
func (api *PublicRestMarketHistoryTradeAPI) ContractCode(contractCode string) *PublicRestMarketHistoryTradeAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// size int true 获取数量，默认150 [1,2000]
func (api *PublicRestMarketHistoryTradeAPI) Size(size int) *PublicRestMarketHistoryTradeAPI {
	api.req.Size = GetPointer(size)
	return api
}