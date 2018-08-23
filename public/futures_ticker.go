package public

type FuturesTickerMethod struct {
	Last        float64 `json:"last"`
	High        float64 `json:"high"`
	Low         float64 `json:"low"`
	Vwap        float64 `json:"vwap"`
	Volume      float64 `json:"volume"`
	Bid         float64 `json:"bid"`
	Ask         float64 `json:"ask"`
	SwapRateBid float64 `json:"swap_rate_bid"`
	SwapRateAsk float64 `json:"swap_rate_ask"`
}

func (api *ApiClient) FuturesTicker(groupId string, currencyPair string) (*FuturesTickerMethod, error) {
	var ticker *FuturesTickerMethod
	err := api.GetRequest("ticker", groupId+"/"+currencyPair, &ticker)
	return ticker, err
}
