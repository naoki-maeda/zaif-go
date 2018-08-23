package public

type FuturesDepthMethod struct {
	Asks []Asks `json:"asks"`
	Bids []Bids `json:"bids"`
}

func (api *ApiClient) FuturesDepth(groupId string, currencyPair string) (*FuturesDepthMethod, error) {
	var depth *FuturesDepthMethod
	err := api.GetRequest("depth", groupId + "/" + currencyPair, &depth)
	return depth, err
}
