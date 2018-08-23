package public

type LastPriceMethod struct {
	LastPrice float64 `json:"last_price"`
}

func (api *ApiClient) LastPrice(currencyPair string) (*LastPriceMethod, error) {
	var lastPrice *LastPriceMethod
	err := api.GetRequest("last_price", currencyPair, &lastPrice)
	return lastPrice, err
}
