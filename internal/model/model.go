package model

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Trace   string      `json:"trace"`
	Data    interface{} `json:"data"`
}
