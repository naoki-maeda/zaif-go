package private

type TradeHistoryParams struct {
	From         int    `url:"from,omitempty"`
	Count        int    `url:"count,omitempty"`
	FromID       int    `url:"from_id,omitempty"`
	EndID        int    `url:"end_id,omitempty"`
	Order        string `url:"order,omitempty"`
	Since        int    `url:"since,omitempty"`
	End          int    `url:"end,omitempty"`
	CurrencyPair string `url:"currency_pair,omitempty"`
	IsToken      bool   `url:"is_token_both,omitempty"`
}

type TradeHistoryResponse struct {
	CurrencyPair string  `json:"currency_pair"`
	Action       string  `json:"action"`
	Amount       float64 `json:"amount"`
	Price        float64 `json:"price"`
	YourAction   string  `json:"your_action"`
	Bonus        float64 `json:"bonus"`
	Timestamp    int     `json:"timestamp"`
	Comment      string  `json:"comment"`
}

type TradeHistoryAPIResponse struct {
	ApiResponse
	Response map[string]TradeHistoryResponse `json:"return"`
}

func (api *ApiClient) TradeHistory(param TradeHistoryParams) (map[string]TradeHistoryResponse, error) {
	var res TradeHistoryAPIResponse
	if err := api.Request("trade_history", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
