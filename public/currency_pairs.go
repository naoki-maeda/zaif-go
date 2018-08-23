package public

type CurrencyPairsMethod struct {
	Name         string  `json:"name"`
	Title        string  `json:"title"`
	CurrencyPair string  `json:"currency_pair"`
	Description  string  `json:"description"`
	IsToken      bool    `json:"is_token"`
	EventNumber  int     `json:"event_number"`
	Seq          int     `json:"seq"`
	ItemUnitMin  float64 `json:"item_unit_min"`
	ItemUnitStep float64 `json:"item_unit_step"`
	ItemJapanese string  `json:"item_japanese"`
	AuxUnitMin   float64 `json:"aux_unit_min"`
	AuxUnitStep  float64 `json:"aux_unit_step"`
	AuxUnitPoint int     `json:"aux_unit_point"`
	AuxJapanese  string  `json:"aux_japanese"`
}

func (api *ApiClient) CurrencyPairs(currencyPair string) ([]CurrencyPairsMethod, error) {
	var currencyPairs []CurrencyPairsMethod
	err := api.GetRequest("currency_pairs", currencyPair, &currencyPairs)
	return currencyPairs, err
}
