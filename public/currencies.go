package public

type Currencies struct {
	Id      int    `json:"id"`
	TokenID int    `json:"token_id"`
	Name    string `json:"name"`
	IsToken bool   `json:"is_token"`
}

func (api *ApiClient) Currencies(currency string) ([]Currencies, error) {
	var currencies []Currencies
	err := api.Request("currencies", currency, &currencies)
	return currencies, err
}
