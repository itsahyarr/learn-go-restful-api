package web

type WebResponse struct {
	Code   int         `json:"id"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
