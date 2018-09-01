package private

type WithdrawHistoryParams struct {
	Currency string `url:"currency"`
	From     int    `url:"from,omitempty"`
	Count    int    `url:"count,omitempty"`
	FromID   int    `url:"from_id,omitempty"`
	EndID    int    `url:"end_id,omitempty"`
	Order    string `url:"order,omitempty"` // ASC or DESC
	Since    int    `url:"since,omitempty"`
	End      int    `url:"end,omitempty"`
}

type WithdrawHistoryResponse struct {
	Timestamp int     `json:"timestamp"`
	Address   string  `json:"address"`
	Amount    float64 `json:"amount"`
	Txid      string  `json:"txid"`
}

type withdrawHistoryAPIResponse struct {
	ApiResponse
	Response map[string]WithdrawHistoryResponse `json:"return"`
}

func (api *ApiClient) WithdrawHistory(param WithdrawHistoryParams) (map[string]WithdrawHistoryResponse, error) {
	var res withdrawHistoryAPIResponse
	if err := api.Request("withdraw_history", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
