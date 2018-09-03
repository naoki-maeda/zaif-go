package private

type ActivePositionsParams struct {
	Type         string `url:"type"`               // margin or futures
	GroupID      int    `url:"group_id,omitempty"` // type = futuresの場合必須
	CurrencyPair string `url:"currency_pair,omitempty"`
}

type ActivePositionsResponse struct {
	GroupId          int     `json:"group_id"`
	CurrencyPair     string  `json:"currency_pair"`
	Action           string  `json:"action"`
	Amount           float64 `json:"amount"`
	Price            float64 `json:"price"`
	Limit            float64 `json:"limit"`
	Stop             float64 `json:"stop"`
	TermEnd          string  `json:"term_end"`
	Leverage         float64 `json:"leverage"`
	FeeSpent         float64 `json:"fee_spent"`
	Timestamp        string  `json:"timestamp"`
	PriceAvg         float64 `json:"price_avg"`
	AmountDone       float64 `json:"amount_done"`
	CloseAvg         float64 `json:"close_avg"`
	CloseDone        float64 `json:"close_done"`
	DepositJPY       float64 `json:"deposit_jpy"`
	DepositBTC       float64 `json:"deposit_btc"`
	DepositXEM       float64 `json:"deposit_xem"`
	DepositMONA      float64 `json:"deposit_mona"`
	DepositPriceJPY  float64 `json:"deposit_price_jpy"`
	DepositPriceBTC  float64 `json:"deposit_price_btc"`
	DepositPriceXEM  float64 `json:"deposit_price_xem"`
	DepositPriceMONA float64 `json:"deposit_price_mona"`
	Swap             float64 `json:"swap"`
}

type ActivePositionsAPIResponse struct {
	ApiResponse
	Response map[string]ActivePositionsResponse `json:"return"`
}

func (api *ApiClient) ActivePositions(param ActivePositionsParams) (map[string]ActivePositionsResponse, error) {
	var res ActivePositionsAPIResponse
	if err := api.Request("active_positions", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
