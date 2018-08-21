package public

type Trades struct {
	Date         int     `json:"date"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	Tid          int     `json:"tid"`
	CurrencyPair string  `json:"currency_pair"`
	TradeType    string  `json:"trade_type"`
}

func (api *ApiClient) Trades(currencyPair string) ([]Trades, error) {
	var trades []Trades
	err := api.Request("trades", currencyPair, &trades)
	return trades, err
}
