package public

type TickerMethod struct {
	Last   float64 `json:"last"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Vwap   float64 `json:"vwap"`
	Volume float64 `json:"volume"`
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
}

func (api *ApiClient) Ticker(currencyPair string) (*TickerMethod, error) {
	var ticker *TickerMethod
	err := api.GetRequest("ticker", currencyPair, &ticker)
	return ticker, err
}
