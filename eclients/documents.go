package eclients

type DocumentFetchQuery struct {
	Type   string `json:"type"`
	Series string `json:"series"`
	Number string `json:"number"`
}
