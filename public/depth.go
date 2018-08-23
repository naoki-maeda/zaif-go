package public

type Asks []float64
type Bids []float64

type DepthMethod struct {
	Asks []Asks `json:"asks"`
	Bids []Bids `json:"bids"`
}

func (api *ApiClient) Depth(currencyPair string) (*DepthMethod, error) {
	var depth *DepthMethod
	err := api.GetRequest("depth", currencyPair, &depth)
	return depth, err
}
