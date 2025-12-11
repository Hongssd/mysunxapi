package mysunxapi

type PrivateRestAPI int

const (
	// Account
	PrivateRestAccountBalance    PrivateRestAPI = iota // 查询账户余额
	PrivateRestAccountBillRecord                       // 组合查询用户财务记录
	PrivateRestAccountFeeRate                          // 查询用户当前的手续费费率

	// Order
	PrivateRestTradeOrderPost         // 下单
	PrivateRestTradeBatchOrders       // 批量下单
	PrivateRestTradeCancelOrder       // 撤单
	PrivateRestTradeCancelBatchOrders // 批量撤单
	PrivateRestTradeCancelAllOrders   // 全部撤单
	PrivateRestTradePosition          // 市价全平
	PrivateRestTradePositionAll       // 一键全平
	PrivateRestTradeOrderOpens        // 查询当前委托
	PrivateRestTradeOrderDetails      // 查询成交明细
	PrivateRestTradeOrderHistory      // 查询历史委托
	PrivateRestTradeOrderGet          // 查询订单信息
	PrivateRestTradeOrderLimit        // 查询用户当前的下单量限制

	// Position
	PrivateRestTradePositionOpens         // 查询当前持仓
	PrivateRestTradePositionLeverGet      // 查询杠杆等级列表
	PrivateRestTradePositionLeverPost     // 设置杠杆等级
	PrivateRestTradePositionModeGet       // 查询持仓模式
	PrivateRestTradePositionModePost      // 设置持仓模式
	PrivateRestTradePositionRiskLimit     // 查询持仓风险限额
	PrivateRestTradePositionPositionLimit // 用户持仓量限制的查询
)

var PrivateRestAPIMap = map[PrivateRestAPI]string{
	// Account
	PrivateRestAccountBalance:    "/sapi/v1/account/balance",     // GET 查询账户余额
	PrivateRestAccountBillRecord: "/sapi/v1/account/bill_record", // POST 组合查询用户财务记录
	PrivateRestAccountFeeRate:    "/sapi/v1/account/fee_rate",    // POST 查询用户当前的手续费费率

	// Order
	PrivateRestTradeOrderPost:         "/sapi/v1/trade/order",               // POST 下单
	PrivateRestTradeBatchOrders:       "/sapi/v1/trade/batch_orders",        // POST 批量下单
	PrivateRestTradeCancelOrder:       "/sapi/v1/trade/cancel_order",        // POST 撤单
	PrivateRestTradeCancelBatchOrders: "/sapi/v1/trade/cancel_batch_orders", // POST 批量撤单
	PrivateRestTradeCancelAllOrders:   "/sapi/v1/trade/cancel_all_orders",   // POST 全部撤单
	PrivateRestTradePosition:          "/sapi/v1/trade/position",            // POST 市价全平
	PrivateRestTradePositionAll:       "/sapi/v1/trade/position_all",        // POST 一键全平
	PrivateRestTradeOrderOpens:        "/sapi/v1/trade/order/opens",         // GET 查询当前委托
	PrivateRestTradeOrderDetails:      "/sapi/v1/trade/order/details",       // GET 查询成交明细
	PrivateRestTradeOrderHistory:      "/sapi/v1/trade/order/history",       // GET 查询历史委托
	PrivateRestTradeOrderGet:          "/sapi/v1/trade/order",               // GET 查询订单信息
	PrivateRestTradeOrderLimit:        "/sapi/v1/trade/order_limit",         // GET 查询用户当前的下单量限制

	// Positon
	PrivateRestTradePositionOpens:         "/sapi/v1/trade/position/opens",    // GET 查询当前持仓
	PrivateRestTradePositionLeverGet:      "/sapi/v1/position/lever",          // GET 查询杠杆等级列表
	PrivateRestTradePositionLeverPost:     "/sapi/v1/position/lever",          // POST 设置杠杆等级
	PrivateRestTradePositionModeGet:       "/sapi/v1/position/mode",           // GET 查询持仓模式
	PrivateRestTradePositionModePost:      "/sapi/v1/position/mode",           // POST 设置持仓模式
	PrivateRestTradePositionRiskLimit:     "/sapi/v1/position/risk/limit",     // GET 查询持仓风险限额
	PrivateRestTradePositionPositionLimit: "/sapi/v1/position/position_limit", // POST 用户持仓量限制的查询
}
