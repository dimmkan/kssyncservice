package services

type GetAllServicesResponse struct {
	Count int `json:"count"`
	Text string `json:"text"`
	Data []Service `json:"data"`
}