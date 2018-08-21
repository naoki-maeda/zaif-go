package private

type TradeParams struct {
	CurrencyPair string  `url:"currency_pair"`
	Action       string  `url:"action"`
	Price        float64 `url:"price"`
	Amount       float64 `url:"amount"`
	Limit        float64 `url:"limit,omitempty"`
	Comment      string  `url:"comment,omitempty"`
}

type TradeResponse struct {
	Received float64 `json:"received"`
	Remains  float64 `json:"remains"`
	OrderId  int     `json:"order_id"`
	Funds    struct {
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

type TradeAPIResponse struct {
	ApiResponse
	Response *TradeResponse `json:"return"`
}

func (api *ApiClient) Trade(param TradeParams) (*TradeResponse, error) {
	var res TradeAPIResponse
	if err := api.Request("trade", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
