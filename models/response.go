package models

type Response struct {
	Status  int         `json:"status"`
	Massage string      `json:"message"`
	Data    interface{} `json:"data"`
}
