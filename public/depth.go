package public

type Asks []float64
type Bids []float64

type Depth struct {
	Asks []Asks `json:"asks"`
	Bids []Bids `json:"bids"`
}

func (api *ApiClient) Depth(currencyPair string) (*Depth, error) {
	var depth *Depth
	err := api.Request("depth", currencyPair, &depth)
	return depth, err
}
