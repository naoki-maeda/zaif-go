package private

type CancelPositionParams struct {
	Type       string `url:"type"`               // margin or futures
	GroupID    int    `url:"group_id,omitempty"` // type = futuresの場合必須
	LeverageId int    `url:"leverage_id"`
}

type CancelPositionResponse struct {
	LeverageId       int     `json:"leverage_id"`
	FeeSpent         float64 `json:"fee_spent"`
	TimestampClosed  string  `json:"timestamp_closed"`
	PriceAvg         float64 `json:"price_avg"`
	AmountDone       float64 `json:"amount_done"`
	CloseAvg         float64 `json:"close_avg"`
	CloseDone        float64 `json:"close_done"`
	RefundedBtc      float64 `json:"refunded_btc"`
	RefundedPriceBtc float64 `json:"refunded_price_btc"`
	Swap             float64 `json:"swap"`
	GuardFee         float64 `json:"guard_fee"`
	Funds            struct {
		BTC    float64 `json:"btc"`
		JPY    float64 `json:"jpy"`
		MONA   float64 `json:"mona"`
		XEM    float64 `json:"xem"`
		ZAIF   float64 `json:"ZAIF"`
		XCP    float64 `json:"XCP"`
		SJCX   float64 `json:"SJCX"`
		BCY    float64 `json:"BITCRISTALS"`
		PEPE   float64 `json:"PEPECASH"`
		FSCC   float64 `json:"FSCC"`
		CICC   float64 `json:"CICC"`
		NCXC   float64 `json:"NCXC"`
		JPYZ   float64 `json:"JPYZ"`
		ETH    float64 `json:"ETH"`
		BCH    float64 `json:"BCH"`
		ERC20  float64 `json:"ERC20.CMS"`
		MOSAIC float64 `json:"MOSAIC.CMS"`
	} `json:"funds"`
}

type CancelPositionAPIResponse struct {
	ApiResponse
	Response *CancelPositionResponse `json:"return"`
}

func (api *ApiClient) CancelPosition(param CancelPositionParams) (*CancelPositionResponse, error) {
	var res CancelPositionAPIResponse
	if err := api.Request("cancel_position", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
