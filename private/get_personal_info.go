package private

type GetPersonalInfoResponse struct {
	RankingNickname string `json:"ranking_nickname"`
	IconPath        string `json:"icon_path"`
}

type GetPersonalInfoApiResponse struct {
	ApiResponse
	Response map[string]GetPersonalInfoResponse `json:"return"`
}

func (api *ApiClient) GetPersonalInfo() (map[string]GetPersonalInfoResponse, error) {
	var res GetPersonalInfoApiResponse
	if err := api.Request("get_personal_info", nil, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
