package model

type Response struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
	Total *int64      `json:"total,omitempty"`
}
