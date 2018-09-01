package public

type GroupsMethod struct {
	Id             int    `json:"id"`
	CurrencyPair   string `json:"currency_pair"`
	StartTimestamp string `json:"start_timestamp"`
	EndTimestamp   string `json:"end_timestamp"`
	UseSwap        bool   `json:"use_swap"`
}

func (api *ApiClient) Groups(groupId string) ([]GroupsMethod, error) {
	var groups []GroupsMethod
	err := api.GetRequest("groups", groupId, &groups)
	return groups, err
}
