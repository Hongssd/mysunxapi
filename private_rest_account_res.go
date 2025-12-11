package mysunxapi

type PrivateRestAccountBalanceRes struct {
	State   string `json:"state"`
	Equity  string `json:"equity"`
	Details []struct {
		Currency              string `json:"currency"`
		Equity                string `json:"equity"`
		Available             string `json:"available"`
		ProfitUnreal          string `json:"profit_unreal"`
		InitialMargin         string `json:"initial_margin"`
		MaintenanceMargin     string `json:"maintenance_margin"`
		MaintenanceMarginRate string `json:"maintenance_margin_rate"`
	} `json:"details"`
	InitialMargin         string `json:"initial_margin"`
	MaintenanceMargin     string `json:"maintenance_margin"`
	MaintenanceMarginRate string `json:"maintenance_margin_rate"`
	ProfitUnreal          string `json:"profit_unreal"`
	AvailableMargin       string `json:"available_margin"`
	VoucherValue          string `json:"voucher_value"`
	CreatedTime           int64  `json:"created_time"`
	UpdatedTime           int64  `json:"updated_time"`
}

type PrivateRestAccountBillRecordResRow struct {
	QueryId           int64  `json:"query_id"` // 下次查询ID
	Id                int64  `json:"id"`
	Ts                int64  `json:"ts"`                  // 创建时间
	Asset             string `json:"asset"`               // 币种
	ContractCode      string `json:"contract_code"`       // 合约代码
	MarginAccount     string `json:"margin_account"`      // 保证金账户
	FaceMarginAccount string `json:"face_margin_account"` // 对手方保证金账户
	Type              int    `json:"type"`                // 交易类型
	Amount            string `json:"amount"`              // 金额（计价货币）
}

type PrivateRestAccountBillRecordRes []PrivateRestAccountBillRecordResRow

type PrivateRestAccountFeeRateResRow struct {
	Symbol        string `json:"symbol"`          // 品种代码
	ContractCode  string `json:"contract_code"`   // 合约代码
	OpenMakerFee  string `json:"open_maker_fee"`  // 开仓挂单的手续费费率，小数形式
	OpenTakerFee  string `json:"open_taker_fee"`  // 开仓吃单的手续费费率，小数形式
	CloseMakerFee string `json:"close_maker_fee"` // 平仓挂单的手续费费率，小数形式
	CloseTakerFee string `json:"close_taker_fee"` // 平仓吃单的手续费费率，小数形式
	FeeAsset      string `json:"fee_asset"`       // 手续费币种 USDT...
	ContractType  string `json:"contract_type"`   // 合约类型
	Pair          string `json:"pair"`            // 交易对 swap（永续）、this_week（当周）、next_week（次周）、quarter（当季）、next_quarter（次季）
	BusinessType  string `json:"business_type"`   // 业务类型 如：“BTC-USDT”
	DeliveryFee   string `json:"delivery_fee"`    // 交割手续费 futures：交割、swap：永续
	// TradePartition string `json:"trade_partition"` // 交易分区
}
type PrivateRestAccountFeeRateRes []PrivateRestAccountFeeRateResRow
