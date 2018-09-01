package private

type ActiveOrdersParams struct {
	CurrencyPair string `url:"currency_pair,omitempty"`
	IsToken      bool   `url:"is_token,omitempty"`      // 非推奨
	IsTokenBoth  bool   `url:"is_token_both,omitempty"` // 非推奨
}

type ActiveOrdersResponse struct {
	CurrencyPair string  `json:"currency_pair"`
	Action       string  `json:"action"`
	Amount       float64 `json:"amount"`
	Price        float64 `json:"price"`
	Timestamp    string  `json:"timestamp"`
}

type activeOrdersApiResponse struct {
	ApiResponse
	Response map[string]ActiveOrdersResponse `json:"return"`
}

func (api *ApiClient) ActiveOrders(param ActiveOrdersParams) (map[string]ActiveOrdersResponse, error) {
	var res activeOrdersApiResponse
	if err := api.Request("active_orders", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
