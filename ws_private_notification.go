package mysunxapi

import "fmt"

type WsNotificationCommonReq struct {
	Op    string `json:"op"`
	Cid   string `json:"cid"`
	Topic string `json:"topic"`
}

func (ws *PrivateWsStreamClient) SubscribeAccount(isSubscribe bool) (*Subscription[WsNotificationCommonReq, WsAccountRes], error) {
	subKey := "account"

	if isSubscribe {
		if existingSub, ok := ws.accountSubMap.Load(subKey); ok {
			return existingSub, nil
		}
	}

	reqId := node.Generate().String()
	subReq := []*WsNotificationCommonReq{}
	if isSubscribe {
		subReq = append(subReq, &WsNotificationCommonReq{
			Op:    "sub",
			Cid:   reqId,
			Topic: subKey,
		})
	} else {
		subReq = append(subReq, &WsNotificationCommonReq{
			Op:    "unsub",
			Cid:   reqId,
			Topic: subKey,
		})
	}
	sub, err := subscribe[WsNotificationCommonReq, WsAccountRes](&ws.WsStreamClient, subReq, reqId)
	if err != nil {
		log.Error("SubscribeAccount error: ", err)
		return nil, err
	}

	ws.accountSubMap.Store(subKey, sub)

	return sub, nil
}

type WsPositionsReq struct {
	WsNotificationCommonReq
	ContractCode string `json:"contract_code"`
}

// 订阅持仓变动数据 positions.$contract_code
func (ws *PrivateWsStreamClient) SubscribePositions(contractCodes []string, isSubscribe bool) (*Subscription[WsPositionsReq, WsPositions], error) {
	if len(contractCodes) == 0 {
		return nil, fmt.Errorf("contractCode is required")
	}

	subKey := "positions"

	if isSubscribe {
		if existingSub, ok := ws.positionsSubMap.Load(subKey); ok {
			return existingSub, nil
		}
	}

	reqId := node.Generate().String()
	subReqs := []*WsPositionsReq{}
	if isSubscribe {
		for _, contractCode := range contractCodes {
			subReqs = append(subReqs, &WsPositionsReq{
				WsNotificationCommonReq: WsNotificationCommonReq{
					Op:    "sub",
					Cid:   reqId,
					Topic: subKey,
				},
				ContractCode: contractCode,
			})
		}
	} else {
		for _, contractCode := range contractCodes {
			subReqs = append(subReqs, &WsPositionsReq{
				WsNotificationCommonReq: WsNotificationCommonReq{
					Op:    "unsub",
					Cid:   reqId,
					Topic: subKey,
				},
				ContractCode: contractCode,
			})
		}
	}
	sub, err := subscribe[WsPositionsReq, WsPositions](&ws.WsStreamClient, subReqs, reqId)
	if err != nil {
		log.Error("SubscribePositions error: ", err)
		return nil, err
	}

	ws.positionsSubMap.Store(subKey, sub)

	return sub, nil
}

type WsMatchOrdersReq struct {
	WsNotificationCommonReq
	ContractCode string `json:"contract_code"`
}

// 订阅撮合后的订单数据  match_orders.$contract_code
func (ws *PrivateWsStreamClient) SubscribeMatchOrders(contractCodes []string, isSubscribe bool) (*Subscription[WsMatchOrdersReq, WsMatchOrders], error) {
	if len(contractCodes) == 0 {
		return nil, fmt.Errorf("contractCode is required")
	}

	subKey := "match_orders"

	if isSubscribe {
		if existingSub, ok := ws.matchOrdersSubMap.Load(subKey); ok {
			return existingSub, nil
		}
	}

	reqId := node.Generate().String()
	subReqs := []*WsMatchOrdersReq{}
	if isSubscribe {
		for _, contractCode := range contractCodes {
			subReqs = append(subReqs, &WsMatchOrdersReq{
				WsNotificationCommonReq: WsNotificationCommonReq{
					Op:    "sub",
					Cid:   reqId,
					Topic: subKey,
				},
				ContractCode: contractCode,
			})
		}
	} else {
		for _, contractCode := range contractCodes {
			subReqs = append(subReqs, &WsMatchOrdersReq{
				WsNotificationCommonReq: WsNotificationCommonReq{
					Op:    "unsub",
					Cid:   reqId,
					Topic: subKey,
				},
				ContractCode: contractCode,
			})
		}
	}
	sub, err := subscribe[WsMatchOrdersReq, WsMatchOrders](&ws.WsStreamClient, subReqs, reqId)
	if err != nil {
		log.Error("SubscribeMatchOrders error: ", err)
		return nil, err
	}

	ws.matchOrdersSubMap.Store(subKey, sub)

	return sub, nil
}

type WsTradeReq struct {
	WsNotificationCommonReq
	ContractCode string `json:"contract_code"`
}

// 订阅成交变动数据 trade.$contract_code
func (ws *PrivateWsStreamClient) SubscribeTrade(contractCodes []string, isSubscribe bool) (*Subscription[WsTradeReq, WsTrade], error) {
	if len(contractCodes) == 0 {
		return nil, fmt.Errorf("contractCode is required")
	}

	subKey := "trade"

	if isSubscribe {
		if existingSub, ok := ws.tradeSubMap.Load(subKey); ok {
			return existingSub, nil
		}
	}

	reqId := node.Generate().String()
	subReqs := []*WsTradeReq{}

	if isSubscribe {
		for _, contractCode := range contractCodes {
			subReqs = append(subReqs, &WsTradeReq{
				WsNotificationCommonReq: WsNotificationCommonReq{
					Op:    "sub",
					Cid:   reqId,
					Topic: subKey,
				},
				ContractCode: contractCode,
			})
		}
	} else {
		for _, contractCode := range contractCodes {
			subReqs = append(subReqs, &WsTradeReq{
				WsNotificationCommonReq: WsNotificationCommonReq{
					Op:    "unsub",
					Cid:   reqId,
					Topic: subKey,
				},
				ContractCode: contractCode,
			})
		}
	}
	sub, err := subscribe[WsTradeReq, WsTrade](&ws.WsStreamClient, subReqs, reqId)
	if err != nil {
		log.Error("SubscribeTrade error: ", err)
		return nil, err
	}

	ws.tradeSubMap.Store(subKey, sub)

	return sub, nil
}

type WsOrdersReq struct {
	WsNotificationCommonReq
	ContractCode string `json:"contract_code"`
}

// 订单数据 orders.$contract_code
func (ws *PrivateWsStreamClient) SubscribeOrders(contractCodes []string, isSubscribe bool) (*Subscription[WsOrdersReq, WsOrders], error) {
	if len(contractCodes) == 0 {
		return nil, fmt.Errorf("contractCode is required")
	}

	subKey := "orders"

	if isSubscribe {
		if existingSub, ok := ws.ordersSubMap.Load(subKey); ok {
			return existingSub, nil
		}
	}

	reqId := node.Generate().String()
	subReqs := []*WsOrdersReq{}

	if isSubscribe {
		for _, contractCode := range contractCodes {
			subReqs = append(subReqs, &WsOrdersReq{
				WsNotificationCommonReq: WsNotificationCommonReq{
					Op:    "sub",
					Cid:   reqId,
					Topic: subKey,
				},
				ContractCode: contractCode,
			})
		}
	} else {
		for _, contractCode := range contractCodes {
			subReqs = append(subReqs, &WsOrdersReq{
				WsNotificationCommonReq: WsNotificationCommonReq{
					Op:    "unsub",
					Cid:   reqId,
					Topic: subKey,
				},
				ContractCode: contractCode,
			})
		}
	}
	sub, err := subscribe[WsOrdersReq, WsOrders](&ws.WsStreamClient, subReqs, reqId)
	if err != nil {
		log.Error("SubscribeOrders error: ", err)
		return nil, err
	}

	ws.ordersSubMap.Store(subKey, sub)

	return sub, nil
}
