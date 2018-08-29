package public

type SwapHistoryMethod struct {
	Timestamp   int     `json:"timestamp"`
	SwapRateBid float64 `json:"swap_rate_bid"`
	SwapRateAsk float64 `json:"swap_rate_ask"`
}

func (api *ApiClient) SwapHistory(groupId string, currencyPair string) ([]SwapHistoryMethod, error) {
	var swapHistory []SwapHistoryMethod
	err := api.GetRequest("swap_history", groupId+"/"+currencyPair, &swapHistory)
	return swapHistory, err
}
