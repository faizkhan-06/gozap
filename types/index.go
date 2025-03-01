package types

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

type UrlData struct {
	LongUrl string `json:"long_url"`
}