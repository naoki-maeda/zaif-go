package public

type FuturesLastPriceMethod struct {
	GroupId   int     `json:"group_id"`
	LastPrice float64 `json:"currency_pair"`
}

func (api *ApiClient) FuturesLastPrice(groupId string, currencyPair string) (*FuturesLastPriceMethod, error) {
	var lastPrice *FuturesLastPriceMethod
	err := api.GetRequest("last_price", groupId+"/"+currencyPair, &lastPrice)
	return lastPrice, err
}
