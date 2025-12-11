package mysunxapi

type PublicRestPublicContractInfoResRow struct {
	Symbol            string  `json:"symbol"`              // 品种代码
	ContractCode      string  `json:"contract_code"`       // 合约代码
	ContractSize      float64 `json:"contract_size"`       // 合约面值，即1张合约对应多少标的币种（如BTC-USDT合约则面值单位就是BTC）
	PriceTick         float64 `json:"price_tick"`          // 合约价格最小变动精度
	SettlementDate    string  `json:"settlement_date"`     // 合约下次结算时间
	DeliveryTime      string  `json:"delivery_time"`       // 交割时间（合约无需交割时，该字段值为空字符串），单位：毫秒
	CreateDate        string  `json:"create_date"`         // 合约上市日期
	ContractStatus    int     `json:"contract_status"`     // 合约状态
	SupportMarginMode string  `json:"support_margin_mode"` // 合约支持的保证金模式
	ContractType      string  `json:"contract_type"`       // 合约类型
	Pair              string  `json:"pair"`                // 交易对
	BusinessType      string  `json:"business_type"`       // 业务类型
	DeliveryDate      string  `json:"delivery_date"`       // 合约交割日期,永续无需交割时该字段为空字符串
	// Adjust            string          `json:"adjust"`              // 无效字段
	// PriceEstimated    string          `json:"price_estimated"`     // 无效字段
	// OpenType          int             `json:"open_type"`           // 无效字段
	// TradePartition    string          `json:"trade_partition"`     // 参数已废弃
}
type PublicRestPublicContractInfoRes []PublicRestPublicContractInfoResRow

type PublicRestPublicIndexResRow struct {
	ContractCode string  `json:"contract_code"` // 指数代码
	IndexPrice   float64 `json:"index_price"`   // 指数价格
	IndexTs      int64   `json:"index_ts"`      // 响应生成时间点，单位：毫秒
	// TradePartition string          `json:"trade_partition"` // 参数已废弃
}
type PublicRestPublicIndexRes []PublicRestPublicIndexResRow

type PublicRestPublicRiskLimitResRow struct {
	ContractCode          string `json:"contract_code"`           // 合约代码 如 永续："BTC-USDT"... ，交割："BTC-USDT-210625"...
	MarginMode            string `json:"margin_mode"`             // 保证金模式 cross：全仓模式；isolated：逐仓模式；
	Tier                  int    `json:"tier"`                    // View designated position tier
	MaxLever              string `json:"max_lever"`               // Current maximum leverage
	MaintenanceMarginRate string `json:"maintenance_margin_rate"` // Current maintenance margin ratio
	MaxVolume             string `json:"max_volume"`              // Current maximum position Cont
	MinVolume             string `json:"min_volume"`              // Current minimum position Contract
}
type PublicRestPublicRiskLimitRes []PublicRestPublicRiskLimitResRow

type PublicRestPublicFundingRateRes struct {
	Symbol          string `json:"symbol"`            // 品种代码 BTC,"ETH" ...
	ContractCode    string `json:"contract_code"`     // 合约代码 如 永续："BTC-USDT"... ，交割："BTC-USDT-210625"...
	FeeAsset        string `json:"fee_asset"`         // 资金费币种 USDT...
	FundingTime     string `json:"funding_time"`      // 当期资金费率时间
	FundingRate     string `json:"funding_rate"`      // 当期资金费率 （一分钟计算一次）
	EstimatedRate   string `json:"estimated_rate"`    // 下一期预测资金费率（废弃，默认为null）
	NextFundingTime string `json:"next_funding_time"` // 下一期资金费率时间（废弃，默认为null）
	// TradePartition  string `json:"trade_partition"`   // 参数已废弃
}

type PublicRestPublicFundingRateHistoryRes struct {
	TotalPage   int `json:"total_page"`   // 总页数
	CurrentPage int `json:"current_page"` // 当前页
	TotalSize   int `json:"total_size"`   // 总条数
	// TradePartition string `json:"trade_partition"` // 参数已废弃
	Data []struct {
		Symbol          string `json:"symbol"`            // 品种代码 BTC,"ETH" ...
		ContractCode    string `json:"contract_code"`     // 合约代码 如 永续："BTC-USDT"... ，交割："BTC-USDT-210625"...
		FeeAsset        string `json:"fee_asset"`         // 资金费币种 USDT...
		FundingTime     string `json:"funding_time"`      // 当期资金费率时间
		FundingRate     string `json:"funding_rate"`      // 当期资金费率 （一分钟计算一次）
		RealizedRate    string `json:"realized_rate"`     // 实际资金费率（废弃，默认为null）
		AvgPremiumIndex string `json:"avg_premium_index"` // 平均溢价指数
	} `json:"data"`
}

type PublicRestPublicPriceLimitResRow struct {
	Symbol       string  `json:"symbol"`        // 品种代码
	ContractCode string  `json:"contract_code"` // 合约代码
	HighLimit    float64 `json:"high_limit"`    // 最高买价
	LowLimit     float64 `json:"low_limit"`     // 最低卖价
	BusinessType string  `json:"business_type"` // 业务类型
	Pair         string  `json:"pair"`          // 交易对
	ContractType string  `json:"contract_type"` // 合约类型
	// TradePartition    string  `json:"trade_partition"`     // 参数已废弃
}

type PublicRestPublicPriceLimitRes []PublicRestPublicPriceLimitResRow
type PublicRestPublicMultiAssetsMarginRes struct {
	MultiAssets []string `json:"multi_assets"`
}
