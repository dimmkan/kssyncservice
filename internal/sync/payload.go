package sync

type ServicesResponse struct {
	Count int `json:"count"`
	Text string `json:"text"`
	Data []Tmp_Ksservice `json:"data"`
}