package public

type FuturesLastPriceMethod struct {
	LastPrice float64 `json:"last_price"`
}

func (api *ApiClient) FuturesLastPrice(groupId string, currencyPair string) (*FuturesLastPriceMethod, error) {
	var lastPrice *FuturesLastPriceMethod
	err := api.GetRequest("last_price", groupId+"/"+currencyPair, &lastPrice)
	return lastPrice, err
}
