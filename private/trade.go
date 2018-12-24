package private

import (
	"net/url"
)

type ActiveOrdersParams struct {
	CurrencyPair string `url:"currency_pair"`
	IsToken      bool   `url:"is_token"`      // 非推奨
	IsTokenBoth  bool   `url:"is_token_both"` // 非推奨
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
	values := SetParam(param)
	if err := api.Request("active_orders", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type DepositHistoryParams struct {
	Currency string `url:"currency"`
	From     int    `url:"from"`
	Count    int    `url:"count"`
	FromID   int    `url:"from_id"`
	EndID    int    `url:"end_id"`
	Order    string `url:"order"` // ASC or DESC
	Since    int    `url:"since"`
	End      int    `url:"end"`
}

type DepositHistoryResponse struct {
	Timestamp string  `json:"timestamp"`
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
	values := SetParam(param)
	if err := api.Request("deposit_history", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type WithdrawHistoryParams struct {
	Currency string `url:"currency"`
	From     int    `url:"from"`
	Count    int    `url:"count"`
	FromID   int    `url:"from_id"`
	EndID    int    `url:"end_id"`
	Order    string `url:"order"` // ASC or DESC
	Since    int    `url:"since"`
	End      int    `url:"end"`
}

type WithdrawHistoryResponse struct {
	Timestamp string  `json:"timestamp"`
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
	values := SetParam(param)
	if err := api.Request("withdraw_history", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type TradeHistoryParams struct {
	From         int    `url:"from"`
	Count        int    `url:"count"`
	FromID       int    `url:"from_id"`
	EndID        int    `url:"end_id"`
	Order        string `url:"order"`
	Since        int    `url:"since"`
	End          int    `url:"end"`
	CurrencyPair string `url:"currency_pair"`
	IsToken      bool   `url:"is_token_both"`
}

type TradeHistoryResponse struct {
	CurrencyPair string  `json:"currency_pair"`
	Action       string  `json:"action"`
	Amount       float64 `json:"amount"`
	Price        float64 `json:"price"`
	YourAction   string  `json:"your_action"`
	Bonus        float64 `json:"bonus"`
	Timestamp    string  `json:"timestamp"`
	Comment      string  `json:"comment"`
}

type TradeHistoryAPIResponse struct {
	ApiResponse
	Response map[string]TradeHistoryResponse `json:"return"`
}

func (api *ApiClient) TradeHistory(param TradeHistoryParams) (map[string]TradeHistoryResponse, error) {
	var res TradeHistoryAPIResponse
	values := SetParam(param)
	if err := api.Request("trade_history", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type TradeParams struct {
	CurrencyPair string  `url:"currency_pair"`
	Action       string  `url:"action"`
	Price        float64 `url:"price"`
	Amount       float64 `url:"amount"`
	Limit        float64 `url:"limit"`
	Comment      string  `url:"comment"`
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
	values := SetParam(param)
	if err := api.Request("trade", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type WithdrawParams struct {
	Currency string  `url:"currency"`
	Address  string  `url:"address"`
	Message  string  `url:"message"`
	Amount   float64 `url:"amount"`
	OptFee   float64 `url:"opt_fee"`
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
	values := SetParam(param)
	if err := api.Request("withdraw", values, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type GetIdInfoResponse struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Kana      string `json:"kana"`
	Certified bool   `json:"certified"`
}

type GetIdInfoApiResponse struct {
	ApiResponse
	Response *GetIdInfoResponse `json:"return"`
}

func (api *ApiClient) GetIdInfo() (*GetIdInfoResponse, error) {
	var res GetIdInfoApiResponse
	if err := api.Request("get_id_info", url.Values{}, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type GetInfoResponse struct {
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
	Deposit struct {
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
	} `json:"deposit"`
	Rights struct {
		IdInfo       int `json:"id_info"`
		Info         int `json:"info"`
		PersonalInfo int `json:"personal_info"`
		Trade        int `json:"private"`
		Withdraw     int `json:"withdraw"`
	} `json:"rights"`
	OpenOrders int `json:"open_orders"`
	ServerTime int `json:"server_time"`
	TradeCount int `json:"trade_count"`
}

type getInfoApiResponse struct {
	ApiResponse
	Response *GetInfoResponse `json:"return"`
}

func (api *ApiClient) GetInfo() (*GetInfoResponse, error) {
	var res getInfoApiResponse
	if err := api.Request("get_info", url.Values{}, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type GetInfo2Response struct {
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
	Deposit struct {
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
	} `json:"deposit"`
	Rights struct {
		IdInfo       int `json:"id_info"`
		Info         int `json:"info"`
		PersonalInfo int `json:"personal_info"`
		Trade        int `json:"private"`
		Withdraw     int `json:"withdraw"`
	} `json:"rights"`
	OpenOrders int `json:"open_orders"`
	ServerTime int `json:"server_time"`
	TradeCount int `json:"trade_count"`
}

type getInfo2ApiResponse struct {
	ApiResponse
	Response *GetInfo2Response `json:"return"`
}

func (api *ApiClient) GetInfo2() (*GetInfo2Response, error) {
	var res getInfo2ApiResponse
	if err := api.Request("get_info2", url.Values{}, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}

type GetPersonalInfoResponse struct {
	RankingNickname string `json:"ranking_nickname"`
	IconPath        string `json:"icon_path"`
}

type GetPersonalInfoApiResponse struct {
	ApiResponse
	Response *GetPersonalInfoResponse `json:"return"`
}

func (api *ApiClient) GetPersonalInfo() (*GetPersonalInfoResponse, error) {
	var res GetPersonalInfoApiResponse
	if err := api.Request("get_personal_info", url.Values{}, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
