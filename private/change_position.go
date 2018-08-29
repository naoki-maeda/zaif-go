package private

type ChangePositionParams struct {
	Type       string  `url:"type"`               // margin or futures
	GroupID    int     `url:"group_id,omitempty"` // type = futuresの場合必須
	LeverageId int     `url:"leverage_id"`
	Price      float64 `url:"price"`
	Limit      float64 `url:"limit,omitempty"`
	Stop       float64 `url:"stop,omitempty"`
}

type ChangePositionResponse struct {
	LeverageId       int     `json:"leverage_id"`
	TimestampClosed  int     `json:"timestamp_closed"`
	PriceAvg         float64 `json:"price_avg"`
	AmountDone       float64 `json:"amount_done"`
	CloseAvg         float64 `json:"close_avg"`
	CloseDone        float64 `json:"close_done"`
	RefundedBtc      float64 `json:"refunded_btc"`
	RefundedPriceBtc float64 `json:"refunded_price_btc"`
	Swap             float64 `json:"swap"`
	GuardFee         float64 `json:"guard_fee"`
}

type ChangePositionAPIResponse struct {
	ApiResponse
	Response *ChangePositionResponse `json:"return"`
}

func (api *ApiClient) ChangePosition(param ChangePositionParams) (*ChangePositionResponse, error) {
	var res ChangePositionAPIResponse
	if err := api.Request("change_position", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
