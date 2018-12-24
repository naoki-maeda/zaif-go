package public

type FuturesDepthMethod struct {
	Asks []Asks `json:"asks"`
	Bids []Bids `json:"bids"`
}

func (api *ApiClient) FuturesDepth(groupId string, currencyPair string) (*FuturesDepthMethod, error) {
	var depth *FuturesDepthMethod
	err := api.GetRequest("depth", groupId+"/"+currencyPair, &depth)
	return depth, err
}

type GroupsMethod struct {
	Id             int    `json:"id"`
	CurrencyPair   string `json:"currency_pair"`
	StartTimestamp int    `json:"start_timestamp"`
	EndTimestamp   int    `json:"end_timestamp"`
	UseSwap        bool   `json:"use_swap"`
}

func (api *ApiClient) Groups(groupId string) ([]GroupsMethod, error) {
	var groups []GroupsMethod
	err := api.GetRequest("groups", groupId, &groups)
	return groups, err
}

type FuturesLastPriceMethod struct {
	LastPrice float64 `json:"last_price"`
}

func (api *ApiClient) FuturesLastPrice(groupId string, currencyPair string) (*FuturesLastPriceMethod, error) {
	var lastPrice *FuturesLastPriceMethod
	err := api.GetRequest("last_price", groupId+"/"+currencyPair, &lastPrice)
	return lastPrice, err
}

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
