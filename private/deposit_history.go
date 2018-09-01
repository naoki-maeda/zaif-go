package private

type DepositHistoryParams struct {
	Currency string `url:"currency"`
	From     int    `url:"from,omitempty"`
	Count    int    `url:"count,omitempty"`
	FromID   int    `url:"from_id,omitempty"`
	EndID    int    `url:"end_id,omitempty"`
	Order    string `url:"order,omitempty"` // ASC or DESC
	Since    int    `url:"since,omitempty"`
	End      int    `url:"end,omitempty"`
}

type DepositHistoryResponse struct {
	Timestamp int     `json:"timestamp"`
	Address   string  `json:"address"`
	Amount    float64 `json:"amount"`
	Txid      string  `json:"txid"`
}

type depositHistoryApiResponse struct {
	ApiResponse
	Response map[string]DepositHistoryResponse `json:"return"`
}

func (api *ApiClient) DepositHistory(param DepositHistoryParams) (map[string]DepositHistoryResponse, error) {
	var res depositHistoryApiResponse
	if err := api.Request("deposit_history", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
