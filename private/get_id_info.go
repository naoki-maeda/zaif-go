package private

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
	if err := api.Request("get_id_info", nil, &res); err != nil {
		return nil, err
	}
	if res.Success == 0 {
		return nil, ApiError{Message: res.Error}
	}
	return res.Response, nil
}
