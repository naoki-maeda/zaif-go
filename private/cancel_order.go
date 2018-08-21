package private

type CancelOrderParams struct {
	OrderID      int    `url:"order_id"`
	CurrencyPair string `url:"currency_pair,omitempty"`
	IsToken      bool   `url:"is_token,omitempty"`
}

type CancelOrderResponse struct {
	OrderID int `json:"order_id"`
	Funds   struct {
		BTC    float64 `json:"btc"`
		JPY    float64 `json:"jpy"`
		MONA   float64 `json:"mona"`
		XEM    float64 `json:"xem"`
		ZAIF   float64 `json:"ZAIF"`
		XCP    float64 `json:"XCP"`
		SJCX   float64 `json:"SJCX"`
		BITCR  float64 `json:"BITCRISTALS"`
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

type CancelOrderApiResponse struct {
	ApiResponse
	Response *CancelOrderResponse `json:"return"`
}

func (api *ApiClient) CancelOrder(param CancelOrderParams) (*CancelOrderResponse, error) {
	var res CancelOrderApiResponse
	if err := api.Request("cancel_order", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
