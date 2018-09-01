package private

type GetPositionsParams struct {
	Type         string `url:"type"` // margin or futures
	GroupID      int    `url:"group_id,omitempty"`
	From         int    `url:"from,omitempty"`
	Count        int    `url:"count,omitempty"`
	FromID       int    `url:"from_id,omitempty"`
	EndID        int    `url:"end_id,omitempty"`
	Order        string `url:"order,omitempty"` // ASC or DESC
	Since        int    `url:"since,omitempty"`
	End          int    `url:"end,omitempty"`
	CurrencyPair string `url:"currency_pair,omitempty"`
}

type GetPositionsResponse struct {
	GroupId           int     `json:"group_id"`
	CurrencyPair      string  `json:"currency_pair"`
	Action            string  `json:"action"`
	Amount            float64 `json:"amount"`
	Price             float64 `json:"price"`
	Limit             float64 `json:"limit"`
	Stop              float64 `json:"stop"`
	TermEnd           int     `json:"term_end"`
	Leverage          float64 `json:"leverage"`
	FeeSpent          float64 `json:"fee_spent"`
	Timestamp         int     `json:"timestamp"`
	PriceAvg          float64 `json:"price_avg"`
	AmountDone        float64 `json:"amount_done"`
	CloseAvg          float64 `json:"close_avg"`
	CloseDone         float64 `json:"close_done"`
	RefundedJPY       float64 `json:"refunded_jpy"`
	RefundedBTC       float64 `json:"refunded_btc"`
	RefundedXEM       float64 `json:"refunded_xem"`
	RefundedMONA      float64 `json:"refunded_mona"`
	RefundedPriceJPY  float64 `json:"refunded_price_jpy"`
	RefundedPriceBTC  float64 `json:"refunded_price_btc"`
	RefundedPriceXEM  float64 `json:"refunded_price_xem"`
	RefundedPriceMONA float64 `json:"refunded_price_mona"`
	Swap              float64 `json:"swap"`
	GuardFee          float64 `json:"guard_fee"`
}

type GetPositionsAPIResponse struct {
	ApiResponse
	Response map[string]GetPositionsResponse `json:"return"`
}

func (api *ApiClient) GetPositions(param GetPositionsParams) (map[string]GetPositionsResponse, error) {
	var res GetPositionsAPIResponse
	if err := api.Request("get_positions", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
