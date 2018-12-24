package private

type ActivePositionsParams struct {
	Type         string `url:"type"`     // margin or futures
	GroupID      int    `url:"group_id"` // type = futuresの場合必須
	CurrencyPair string `url:"currency_pair"`
}

type ActivePositionsResponse struct {
	GroupId          int     `json:"group_id"`
	CurrencyPair     string  `json:"currency_pair"`
	Action           string  `json:"action"`
	Amount           float64 `json:"amount"`
	Price            float64 `json:"price"`
	Limit            float64 `json:"limit"`
	Stop             float64 `json:"stop"`
	TermEnd          string  `json:"term_end"`
	Leverage         float64 `json:"leverage"`
	FeeSpent         float64 `json:"fee_spent"`
	Timestamp        string  `json:"timestamp"`
	PriceAvg         float64 `json:"price_avg"`
	AmountDone       float64 `json:"amount_done"`
	CloseAvg         float64 `json:"close_avg"`
	CloseDone        float64 `json:"close_done"`
	DepositJPY       float64 `json:"deposit_jpy"`
	DepositBTC       float64 `json:"deposit_btc"`
	DepositXEM       float64 `json:"deposit_xem"`
	DepositMONA      float64 `json:"deposit_mona"`
	DepositPriceJPY  float64 `json:"deposit_price_jpy"`
	DepositPriceBTC  float64 `json:"deposit_price_btc"`
	DepositPriceXEM  float64 `json:"deposit_price_xem"`
	DepositPriceMONA float64 `json:"deposit_price_mona"`
	Swap             float64 `json:"swap"`
}

type ActivePositionsAPIResponse struct {
	ApiResponse
	Response map[string]ActivePositionsResponse `json:"return"`
}

func (api *ApiClient) ActivePositions(param ActivePositionsParams) (map[string]ActivePositionsResponse, error) {
	var res ActivePositionsAPIResponse
	values := SetParam(param)
	if err := api.Request("active_positions", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type CancelOrderParams struct {
	OrderID      int    `url:"order_id"`
	CurrencyPair string `url:"currency_pair"`
	IsToken      bool   `url:"is_token"`
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
	values := SetParam(param)
	if err := api.Request("cancel_order", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type CancelPositionParams struct {
	Type       string `url:"type"`     // margin or futures
	GroupID    int    `url:"group_id"` // type = futuresの場合必須
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
	values := SetParam(param)
	if err := api.Request("cancel_position", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type ChangePositionParams struct {
	Type       string  `url:"type"`     // margin or futures
	GroupID    int     `url:"group_id"` // type = futuresの場合必須
	LeverageId int     `url:"leverage_id"`
	Price      float64 `url:"price"`
	Limit      float64 `url:"limit"`
	Stop       float64 `url:"stop"`
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
	values := SetParam(param)
	if err := api.Request("change_position", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type CreatePositionParams struct {
	Type         string  `url:"type"`     // margin or futures
	GroupID      int     `url:"group_id"` // type = futuresの場合必須
	CurrencyPair string  `url:"currency_pair"`
	Action       string  `url:"action"`
	Price        float64 `url:"price"`
	Amount       float64 `url:"amount"`
	Leverage     float64 `url:"leverage"`
	Limit        float64 `url:"limit"`
	Stop         float64 `url:"stop"`
}

type CreatePositionResponse struct {
	LeverageId       int     `json:"leverage_id"`
	Timestamp        string  `json:"timestamp"`
	TermEnd          string  `json:"term_end"`
	PriceAvg         float64 `json:"price_avg"`
	AmountDone       float64 `json:"amount_done"`
	DepositJPY       float64 `json:"deposit_jpy"`
	DepositBTC       float64 `json:"deposit_btc"`
	DepositXEM       float64 `json:"deposit_xem"`
	DepositMONA      float64 `json:"deposit_mona"`
	DepositPriceJPY  float64 `json:"deposit_price_jpy"`
	DepositPriceBTC  float64 `json:"deposit_price_btc"`
	DepositPriceXEM  float64 `json:"deposit_price_xem"`
	DepositPriceMONA float64 `json:"deposit_price_mona"`
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

type CreatePositionAPIResponse struct {
	ApiResponse
	Response *CreatePositionResponse `json:"return"`
}

func (api *ApiClient) CreatePosition(param CreatePositionParams) (*CreatePositionResponse, error) {
	var res CreatePositionAPIResponse
	values := SetParam(param)
	if err := api.Request("create_position", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type GetPositionsParams struct {
	Type         string `url:"type"` // margin or futures
	GroupID      int    `url:"group_id"`
	From         int    `url:"from"`
	Count        int    `url:"count"`
	FromID       int    `url:"from_id"`
	EndID        int    `url:"end_id"`
	Order        string `url:"order"` // ASC or DESC
	Since        int    `url:"since"`
	End          int    `url:"end"`
	CurrencyPair string `url:"currency_pair"`
}

type GetPositionsResponse struct {
	GroupId           int     `json:"group_id"`
	CurrencyPair      string  `json:"currency_pair"`
	Action            string  `json:"action"`
	Amount            float64 `json:"amount"`
	Price             float64 `json:"price"`
	Limit             float64 `json:"limit"`
	Stop              float64 `json:"stop"`
	TermEnd           string  `json:"term_end"`
	Leverage          float64 `json:"leverage"`
	FeeSpent          float64 `json:"fee_spent"`
	Timestamp         string  `json:"timestamp"`
	TimestampClosed   string  `json:"timestamp_closed"`
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
	values := SetParam(param)
	if err := api.Request("get_positions", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type PositionHistoryParams struct {
	Type       string `url:"type"`     // margin or futures
	GroupID    int    `url:"group_id"` // type = futuresの場合必須
	LeverageId int    `url:"leverage_id"`
}

type PositionHistoryResponse struct {
	GroupId       int     `json:"group_id"`
	CurrencyPair  string  `json:"currency_pair"`
	Action        string  `json:"action"`
	Amount        float64 `json:"amount"`
	Price         float64 `json:"price"`
	Timestamp     string  `json:"timestamp"`
	Stop          float64 `json:"stop"`
	YourAction    string  `json:"your_action"`
	BidLeverageId string  `json:"bid_leverage_id"`
	AskLeverageId string  `json:"ask_leverage_id"`
}

type PositionHistoryAPIResponse struct {
	ApiResponse
	Response map[string]PositionHistoryResponse `json:"return"`
}

func (api *ApiClient) PositionHistory(param PositionHistoryParams) (map[string]PositionHistoryResponse, error) {
	var res PositionHistoryAPIResponse
	values := SetParam(param)
	if err := api.Request("position_history", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
