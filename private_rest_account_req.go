package mysunxapi

type PrivateRestAccountBalanceReq struct{}
type PrivateRestAccountBalanceAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountBalanceReq
}

type PrivateRestAccountBillRecordReq struct {
	Contract  *string `json:"contract"`   // 合约代码 支持大小写，示例："BTC-USDT" ...
	MarAcct   *string `json:"mar_acct"`   // 保证金账户 支持范围：USDT 默认值：BTC-USDT,"USDT"(查询全仓时使用)...
	Type      *string `json:"type"`       // 不填查询全部类型,【查询多类型中间用，隔开】 范围：3:平多; 4:平空; 5:开仓手续费-吃单; 6:开仓手续费-挂单; 7:平仓手续费-吃单; 8:平仓手续费-挂单; 9:交割平多; 10:交割平空; 11:交割手续费; 12:强制平多; 13:强制平空; 14:从币币转入; 15:转出至币币; 19:穿仓分摊; 26:系统; 28:活动奖励; 29:返利; 30:资金费-收入; 31:资金费-支出; 34:转出到子账号合约账户; 35:从子账号合约账户转入; 36:转出到母账号合约账户; 37:从母账号合约账户转入;38:从其他保证金账户转入 ;39:转出到其他保证金账户 ;85（系统垫资账户-用户U本位账户（转出垫资账户)） ;86（用户U本位账户-系统垫资发放账户 (转入垫资账户)）
	StartTime *int64  `json:"start_time"` // 查询开始时间, 以数据按创建时间进行查询 取值范围 [((end-time) – 48h), (end-time)] ，查询窗口最大为48小时，窗口平移范围为最近90天。
	EndTime   *int64  `json:"end_time"`   // 查询结束时间, 以数据按创建时间进行查询 取值范围 [(present-90d), present] ，查询窗口最大为48小时，窗口平移范围为最近90天。
	Direct    *string `json:"direct"`     // 查询方向, 方向为next时，数据按照时间正序排列返回，方向为prev时，数据按照时间倒序返回 prev表示向前查询，next表示向后查询。
	FromId    *int64  `json:"from_id"`    // 如果是向前(prev)查询，则赋值为上一次查询结果中得到的最小query_id ；如果是向后(next)查询，则赋值为上一次查询结果中得到的最大query_id
}
type PrivateRestAccountBillRecordAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountBillRecordReq
}

// contract string false 合约代码 支持大小写，示例："BTC-USDT" ...
func (api *PrivateRestAccountBillRecordAPI) Contract(contract string) *PrivateRestAccountBillRecordAPI {
	api.req.Contract = GetPointer(contract)
	return api
}

// mar_acct string false 保证金账户 支持范围：USDT 默认值：BTC-USDT,"USDT"(查询全仓时使用)...
func (api *PrivateRestAccountBillRecordAPI) MarAcct(marAcct string) *PrivateRestAccountBillRecordAPI {
	api.req.MarAcct = GetPointer(marAcct)
	return api
}

// type string false 不填查询全部类型,【查询多类型中间用，隔开】
// 范围：3:平多; 4:平空; 5:开仓手续费-吃单; 6:开仓手续费-挂单; 7:平仓手续费-吃单; 8:平仓手续费-挂单; 9:交割平多; 10:交割平空; 11:交割手续费; 12:强制平多; 13:强制平空; 14:从币币转入; 15:转出至币币; 19:穿仓分摊; 26:系统; 28:活动奖励; 29:返利; 30:资金费-收入; 31:资金费-支出; 34:转出到子账号合约账户; 35:从子账号合约账户转入; 36:转出到母账号合约账户; 37:从母账号合约账户转入;38:从其他保证金账户转入 ;39:转出到其他保证金账户 ;85（系统垫资账户-用户U本位账户（转出垫资账户)） ;86（用户U本位账户-系统垫资发放账户 (转入垫资账户)）
func (api *PrivateRestAccountBillRecordAPI) Type(t string) *PrivateRestAccountBillRecordAPI {
	api.req.Type = GetPointer(t)
	return api
}

// start_time long false 查询开始时间, 以数据按创建时间进行查询 取值范围 [((end-time) – 48h), (end-time)] ，查询窗口最大为48小时，窗口平移范围为最近90天。
func (api *PrivateRestAccountBillRecordAPI) StartTime(startTime int64) *PrivateRestAccountBillRecordAPI {
	api.req.StartTime = GetPointer(startTime)
	return api
}

// end_time long false 查询结束时间, 以数据按创建时间进行查询 取值范围 [(present-90d), present] ，查询窗口最大为48小时，窗口平移范围为最近90天。
func (api *PrivateRestAccountBillRecordAPI) EndTime(endTime int64) *PrivateRestAccountBillRecordAPI {
	api.req.EndTime = GetPointer(endTime)
	return api
}

// direct string false 查询方向, 方向为next时，数据按照时间正序排列返回，方向为prev时，数据按照时间倒序返回 prev表示向前查询，next表示向后查询。
func (api *PrivateRestAccountBillRecordAPI) Direct(direct string) *PrivateRestAccountBillRecordAPI {
	api.req.Direct = GetPointer(direct)
	return api
}

// from_id long false 如果是向前(prev)查询，则赋值为上一次查询结果中得到的最小query_id ；如果是向后(next)查询，则赋值为上一次查询结果中得到的最大query_id
func (api *PrivateRestAccountBillRecordAPI) FromId(fromId int64) *PrivateRestAccountBillRecordAPI {
	api.req.FromId = GetPointer(fromId)
	return api
}

type PrivateRestAccountFeeRateReq struct {
	ContractCode *string `json:"contract_code"` // 合约代码 支持大小写，示例："BTC-USDT" ...
	Pair         *string `json:"pair"`          // 交易对 支持大小写，示例："BTC-USDT" ...
}
type PrivateRestAccountFeeRateAPI struct {
	client *PrivateRestClient
	req    *PrivateRestAccountFeeRateReq
}

// contract_code string false 合约代码 支持大小写，示例："BTC-USDT" ...
func (api *PrivateRestAccountFeeRateAPI) ContractCode(contractCode string) *PrivateRestAccountFeeRateAPI {
	api.req.ContractCode = GetPointer(contractCode)
	return api
}

// pair string false 交易对 支持大小写，示例："BTC-USDT" ...
func (api *PrivateRestAccountFeeRateAPI) Pair(pair string) *PrivateRestAccountFeeRateAPI {
	api.req.Pair = GetPointer(pair)
	return api
}
