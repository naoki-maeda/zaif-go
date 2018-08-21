package private

type WithdrawParams struct {
	Currency string  `url:"currency"`
	Address  string  `url:"address"`
	Message  string  `url:"message,omitempty"`
	Amount   float64 `url:"amount"`
	OptFee   float64 `url:"opt_fee,omitempty"`
}

type WithdrawResponse struct {
	Id    int     `json:"id"`
	Txid  string  `json:"txid"`
	Fee   float64 `json:"fee"`
	Funds struct {
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

type withdrawAPIResponse struct {
	ApiResponse
	Response *WithdrawResponse `json:"return"`
}

func (api *ApiClient) Withdraw(param WithdrawParams) (*WithdrawResponse, error) {
	var res withdrawAPIResponse
	if err := api.Request("withdraw", param, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
