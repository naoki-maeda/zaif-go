package public

type CurrenciesMethod struct {
	Id      int    `json:"id"`
	TokenID int    `json:"token_id"`
	Name    string `json:"name"`
	IsToken bool   `json:"is_token"`
}

func (api *ApiClient) Currencies(currency string) ([]CurrenciesMethod, error) {
	var currencies []CurrenciesMethod
	err := api.GetRequest("currencies", currency, &currencies)
	return currencies, err
}
