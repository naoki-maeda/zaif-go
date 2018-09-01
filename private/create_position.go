package private

type CreatePositionParams struct {
	Type         string  `url:"type"`               // margin or futures
	GroupID      int     `url:"group_id,omitempty"` // type = futuresの場合必須
	CurrencyPair string  `url:"currency_pair"`
	Action       string  `url:"action"`
	Price        float64 `url:"price"`
	Amount       float64 `url:"amount"`
	Leverage     float64 `url:"leverage"`
	Limit        float64 `url:"limit,omitempty"`
	Stop         float64 `url:"stop,omitempty"`
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
	if err := api.Request("create_position", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
