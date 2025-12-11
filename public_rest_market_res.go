package mysunxapi

type PublicRestMarketDepthResMiddle struct {
	Asks    [][]float64 `json:"asks"`    // 卖盘
	Bids    [][]float64 `json:"bids"`    // 买盘
	Ch      string      `json:"ch"`      // 数据所属的 channel，格式： market.period
	Id      int64       `json:"id"`      // 消息id
	Mrid    int64       `json:"mrid"`    // 订单ID
	Ts      int64       `json:"ts"`      // 消息生成时间，单位：毫秒.
	Version int64       `json:"version"` // 版本
}

func (m *PublicRestMarketDepthResMiddle) ConvertToRes() *PublicRestMarketDepthRes {
	var asks, bids []PriceLevel
	for _, ask := range m.Asks {
		asks = append(asks, PriceLevel{Price: ask[0], Volume: ask[1]})
	}
	for _, bid := range m.Bids {
		bids = append(bids, PriceLevel{Price: bid[0], Volume: bid[1]})
	}
	return &PublicRestMarketDepthRes{
		Asks:    asks,
		Bids:    bids,
		Ch:      m.Ch,
		Id:      m.Id,
		Mrid:    m.Mrid,
		Ts:      m.Ts,
		Version: m.Version,
	}
}

type PublicRestMarketDepthRes struct {
	Asks    []PriceLevel `json:"asks"`
	Bids    []PriceLevel `json:"bids"`
	Ch      string       `json:"ch"`
	Id      int64        `json:"id"`
	Mrid    int64        `json:"mrid"`
	Ts      int64        `json:"ts"`
	Version int64        `json:"version"`
}

type PublicRestMarketKlineCommon struct {
	Id            int64   `json:"id"`             // K线ID,也就是K线时间戳，K线起始时间
	Vol           float64 `json:"vol"`            // 成交量(张)。 值是买卖双边之和
	Count         float64 `json:"count"`          // 成交笔数。 值是买卖双边之和
	Open          float64 `json:"open"`           // 开盘价
	Close         float64 `json:"close"`          // 收盘价,当K线为最晚的一根时，是最新成交价
	Low           float64 `json:"low"`            // 最低价
	High          float64 `json:"high"`           // 最高价
	Amount        float64 `json:"amount"`         // 成交量(币), 即 (成交量(张) * 单张合约面值)。 值是买卖双边之和
	TradeTurnover float64 `json:"trade_turnover"` // 成交额，即 sum（每一笔成交张数 * 合约面值 * 成交价格）。 值是买卖双边之和
}
type PublicRestMarketHistoryKlineRes []PublicRestMarketKlineCommon

type PublicRestMarketDetailMergedRes struct {
	Id            int64      `json:"id"`             // K线ID,也就是K线时间戳，K线起始时间
	Vol           string     `json:"vol"`            // 成交量(张)。 值是买卖双边之和
	Count         float64    `json:"count"`          // 成交笔数。 值是买卖双边之和
	Open          string     `json:"open"`           // 开盘价
	Close         string     `json:"close"`          // 收盘价,当K线为最晚的一根时，是最新成交价
	Low           string     `json:"low"`            // 最低价
	High          string     `json:"high"`           // 最高价
	Amount        string     `json:"amount"`         // 成交量(币), 即 (成交量(张) * 单张合约面值)。 值是买卖双边之和
	TradeTurnover string     `json:"trade_turnover"` // 成交额，即 sum（每一笔成交张数 * 合约面值 * 成交价格）。 值是买卖双边之和
	Ask           PriceLevel `json:"ask"`            // 卖一价
	Bid           PriceLevel `json:"bid"`            // 买一价
	Ts            int64      `json:"ts"`             // 时间戳
}
type PublicRestMarketDetailMergedResMiddle struct {
	Id            int64     `json:"id"`             // K线ID,也就是K线时间戳，K线起始时间
	Vol           string    `json:"vol"`            // 成交量(张)。 值是买卖双边之和
	Count         float64   `json:"count"`          // 成交笔数。 值是买卖双边之和
	Open          string    `json:"open"`           // 开盘价
	Close         string    `json:"close"`          // 收盘价,当K线为最晚的一根时，是最新成交价
	Low           string    `json:"low"`            // 最低价
	High          string    `json:"high"`           // 最高价
	Amount        string    `json:"amount"`         // 成交量(币), 即 (成交量(张) * 单张合约面值)。 值是买卖双边之和
	TradeTurnover string    `json:"trade_turnover"` // 成交额，即 sum（每一笔成交张数 * 合约面值 * 成交价格）。 值是买卖双边之和
	Ask           []float64 `json:"ask"`            // 卖一价
	Bid           []float64 `json:"bid"`            // 买一价
	Ts            int64     `json:"ts"`             // 时间戳
}

func (m *PublicRestMarketDetailMergedResMiddle) ConvertToRes() *PublicRestMarketDetailMergedRes {
	return &PublicRestMarketDetailMergedRes{
		Id:            m.Id,
		Vol:           m.Vol,
		Count:         m.Count,
		Open:          m.Open,
		Close:         m.Close,
		Low:           m.Low,
		High:          m.High,
		Amount:        m.Amount,
		TradeTurnover: m.TradeTurnover,
		Ask:           PriceLevel{Price: m.Ask[0], Volume: m.Ask[1]},
		Bid:           PriceLevel{Price: m.Bid[0], Volume: m.Bid[1]},
		Ts:            m.Ts,
	}
}

type PublicRestMarketTradeRes struct {
	Id   int64 `json:"id"` // 订单唯一id（品种唯一）
	Ts   int64 `json:"ts"` // 最新成交时间
	Data []struct {
		Amount        string `json:"amount"`    // 成交量(张)。 值是买卖双边之和
		Direction     string `json:"direction"` // 买卖方向，即taker(主动成交)的方向
		Id            int64  `json:"id"`
		Price         string `json:"price"`          // 成交价
		Ts            int64  `json:"ts"`             // 成交时间
		Quantity      string `json:"quantity"`       // 成交量（币）
		ContractCode  string `json:"contract_code"`  // 合约代码
		BusinessType  string `json:"business_type"`  // 业务类型
		TradeTurnover string `json:"trade_turnover"` // 成交额（计价币种）
		// TradePartition string `json:"trade_partition"` // 参数已废弃
	} `json:"data"`
}

type PublicRestMarketBBOResRow struct {
	ContractCode string     `json:"contract_code"` // 合约代码 或 合约标识 永续："BTC-USDT" ... ，交割：“BTC-USDT-210625”... 或 BTC-USDT-CW（当周合约标识）、BTC-USDT-NW（次周合约标识）、BTC-USDT-CQ（当季合约标识）、BTC-USDT-NQ（次季合约标识）
	BusinessType string     `json:"business_type"` // 业务类型 futures：交割、swap：永续、all：全部
	Mrid         int64      `json:"mrid"`
	Ask          PriceLevel `json:"ask"` // [卖1价,卖1量(张)]
	Bid          PriceLevel `json:"bid"` // [买1价,买1量(张)]
	Ts           int64      `json:"ts"`  // 系统检测orderbook时间点，单位：毫秒
}
type PublicRestMarketBBORes []PublicRestMarketBBOResRow

type PublicRestMarketBBOResMiddleRow struct {
	ContractCode string    `json:"contract_code"` // 合约代码 或 合约标识 永续："BTC-USDT" ... ，交割：“BTC-USDT-210625”... 或 BTC-USDT-CW（当周合约标识）、BTC-USDT-NW（次周合约标识）、BTC-USDT-CQ（当季合约标识）、BTC-USDT-NQ（次季合约标识）
	BusinessType string    `json:"business_type"` // 业务类型 futures：交割、swap：永续、all：全部
	Mrid         int64     `json:"mrid"`
	Ask          []float64 `json:"ask"` // [卖1价,卖1量(张)]
	Bid          []float64 `json:"bid"` // [买1价,买1量(张)]
	Ts           int64     `json:"ts"`  // 系统检测orderbook时间点，单位：毫秒
}
type PublicRestMarketBBOResMiddle []PublicRestMarketBBOResMiddleRow

func (m *PublicRestMarketBBOResMiddle) ConvertToRes() *PublicRestMarketBBORes {
	res := PublicRestMarketBBORes{}
	for _, middleRow := range *m {
		res = append(res, PublicRestMarketBBOResRow{
			ContractCode: middleRow.ContractCode,
			BusinessType: middleRow.BusinessType,
			Mrid:         middleRow.Mrid,
			Ask:          PriceLevel{Price: middleRow.Ask[0], Volume: middleRow.Ask[1]},
			Bid:          PriceLevel{Price: middleRow.Bid[0], Volume: middleRow.Bid[1]},
			Ts:           middleRow.Ts,
		})
	}
	return &res
}

type PublicRestMarketHistoryTradeResRow struct {
	Id   int64 `json:"id"` // 订单唯一id（品种唯一）
	Ts   int64 `json:"ts"` // 成交时间
	Data []struct {
		Amount        float64 `json:"amount"`         // 成交量(张)。 值是买卖双边之和
		Quantity      float64 `json:"quantity"`       // 成交量（币）
		TradeTurnover float64 `json:"trade_turnover"` // 成交额（计价币种）
		Ts            int64   `json:"ts"`             // 成交时间
		Id            int64   `json:"id"`             // 订单唯一id（品种唯一）
		Price         float64 `json:"price"`          // 成交价
		Direction     string  `json:"direction"`      // 买卖方向，即taker(主动成交)的方向
	} `json:"data"`
}
type PublicRestMarketHistoryTradeRes []PublicRestMarketHistoryTradeResRow
