package mysunxapi

type PublicRestAPI int

const (
	// Basic
	PublicRestPublicContractInfo       PublicRestAPI = iota // 获取合约信息
	PublicRestPublicIndex                                   // 获取合约指数信息
	PublicRestPublicRiskLimit                               // 获取合约最高限价和最低限价
	PublicRestPublicFundingRate                             // 获取合约的资金费率
	PublicRestPublicFundingRateHistory                      // 获取合约的历史资金费率
	PublicRestPublicPriceLimit                              // 获取合约最高限价和最低限价
	PublicRestPublicMultiAssetsMargin                       // 查询联合保证金支持币种

	// Market
	PublicRestMarketDepth        // GET 获取行情深度数据
	PublicRestMarketHistoryKline // GET 获取K线数据
	PublicRestMarketDetailMerged // GET 获取聚合行情
	PublicRestMarketTrade        // GET 获取最新成交
	PublicRestMarketBBO          // GET 获取市场最优挂单
	PublicRestMarketHistoryTrade // GET 批量获取最近的交易记录
)

var PublicRestAPIMap = map[PublicRestAPI]string{
	// Public
	PublicRestPublicContractInfo:       "/sapi/v1/public/contract_info",        // GET 获取合约信息
	PublicRestPublicIndex:              "/sapi/v1/public/index",                // GET 获取合约指数信息
	PublicRestPublicRiskLimit:          "/sapi/v1/public/risk/limit",           // GET 获取合约最高限价和最低限价
	PublicRestPublicFundingRate:        "/sapi/v1/public/funding_rate",         // GET 获取合约的资金费率
	PublicRestPublicFundingRateHistory: "/sapi/v1/public/funding_rate_history", // GET 获取合约的历史资金费率
	PublicRestPublicPriceLimit:         "/sapi/v1/public/price_limit",          // GET 获取合约最高限价和最低限价
	PublicRestPublicMultiAssetsMargin:  "/sapi/v1/public/multi_assets_margin",  // GET 查询联合保证金支持币种

	// Market
	PublicRestMarketDepth:        "/sapi/v1/market/depth",         // GET 获取行情深度数据
	PublicRestMarketHistoryKline: "/sapi/v1/market/history/kline", // GET 获取K线数据
	PublicRestMarketDetailMerged: "/sapi/v1/market/detail/merged", // GET 获取聚合行情
	PublicRestMarketTrade:        "/sapi/v1/market/trade",         // GET 获取最新成交
	PublicRestMarketBBO:          "/sapi/v1/market/bbo",           // GET 获取市场最优挂单
	PublicRestMarketHistoryTrade: "/sapi/v1/market/history/trade", // GET 批量获取最近的交易记录
}
