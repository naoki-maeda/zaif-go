package public

type CurrenciesMethod struct {
	Id      int    `json:"id"`
	TokenID int    `json:"token_id"`
	Name    string `json:"name"`
	IsToken bool   `json:"is_token"`
}

func (api *ApiClient) Currencies(currency string) ([]CurrenciesMethod, error) {
	var currencies []CurrenciesMethod
	err := api.GetRequest("currencies", currency, &currencies)
	return currencies, err
}

type CurrencyPairsMethod struct {
	Name         string  `json:"name"`
	Title        string  `json:"title"`
	CurrencyPair string  `json:"currency_pair"`
	Description  string  `json:"description"`
	IsToken      bool    `json:"is_token"`
	EventNumber  int     `json:"event_number"`
	Seq          int     `json:"seq"`
	ItemUnitMin  float64 `json:"item_unit_min"`
	ItemUnitStep float64 `json:"item_unit_step"`
	ItemJapanese string  `json:"item_japanese"`
	AuxUnitMin   float64 `json:"aux_unit_min"`
	AuxUnitStep  float64 `json:"aux_unit_step"`
	AuxUnitPoint int     `json:"aux_unit_point"`
	AuxJapanese  string  `json:"aux_japanese"`
}

func (api *ApiClient) CurrencyPairs(currencyPair string) ([]CurrencyPairsMethod, error) {
	var currencyPairs []CurrencyPairsMethod
	err := api.GetRequest("currency_pairs", currencyPair, &currencyPairs)
	return currencyPairs, err
}

type Asks []float64
type Bids []float64

type DepthMethod struct {
	Asks []Asks `json:"asks"`
	Bids []Bids `json:"bids"`
}

func (api *ApiClient) Depth(currencyPair string) (*DepthMethod, error) {
	var depth *DepthMethod
	err := api.GetRequest("depth", currencyPair, &depth)
	return depth, err
}

type LastPriceMethod struct {
	LastPrice float64 `json:"last_price"`
}

func (api *ApiClient) LastPrice(currencyPair string) (*LastPriceMethod, error) {
	var lastPrice *LastPriceMethod
	err := api.GetRequest("last_price", currencyPair, &lastPrice)
	return lastPrice, err
}

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

type TradesMethod struct {
	Date         int     `json:"date"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	Tid          int     `json:"tid"`
	CurrencyPair string  `json:"currency_pair"`
	TradeType    string  `json:"trade_type"`
}

func (api *ApiClient) Trades(currencyPair string) ([]TradesMethod, error) {
	var trades []TradesMethod
	err := api.GetRequest("trades", currencyPair, &trades)
	return trades, err
}
