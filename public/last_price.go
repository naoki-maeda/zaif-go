package public

type LastPrice struct {
	LastPrice float64 `json:"last_price"`
}

func (api *ApiClient) LastPrice(currencyPair string) (*LastPrice, error) {
	var lastPrice *LastPrice
	err := api.Request("last_price", currencyPair, &lastPrice)
	return lastPrice, err
}
