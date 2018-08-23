package public

type FuturesTradesMethod struct {
	Date         int     `json:"date"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	Tid          int     `json:"tid"`
	CurrencyPair string  `json:"currency_pair"`
	TradeType    string  `json:"trade_type"`
}

func (api *ApiClient) FuturesTrades(groupId string, currencyPair string) ([]FuturesTradesMethod, error) {
	var trades []FuturesTradesMethod
	err := api.GetRequest("trades", groupId+"/"+currencyPair, &trades)
	return trades, err
}
