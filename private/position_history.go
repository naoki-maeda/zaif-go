package private

type PositionHistoryParams struct {
	Type       string `url:"type"`               // margin or futures
	GroupID    int    `url:"group_id,omitempty"` // type = futuresの場合必須
	LeverageId int    `url:"leverage_id"`
}

type PositionHistoryResponse struct {
	GroupId       int     `json:"group_id"`
	CurrencyPair  string  `json:"currency_pair"`
	Action        string  `json:"action"`
	Amount        float64 `json:"amount"`
	Price         float64 `json:"price"`
	Timestamp     int     `json:"timestamp"`
	Stop          float64 `json:"stop"`
	YourAction    string  `json:"your_action"`
	BidLeverageId int     `json:"bid_leverage_id"`
	AskLeverageId int     `json:"ask_leverage_id"`
}

type PositionHistoryAPIResponse struct {
	ApiResponse
	Response map[string]PositionHistoryResponse `json:"return"`
}

func (api *ApiClient) PositionHistory(param PositionHistoryParams) (map[string]PositionHistoryResponse, error) {
	var res PositionHistoryAPIResponse
	if err := api.Request("position_history", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
