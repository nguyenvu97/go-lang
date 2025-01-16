package dto

import "encoding/json"

type Response struct {
	Success    bool `json:"success"`
	Message    string `json:"message"`
	Code       int16 `json:"code"`
	Timestamp  int64`json:"timestamp"`
	Result     json.RawMessage `json:"result"`
}

type ResponseWithInterface struct {
	Success    bool `json:"success"`
	Message    string `json:"message"`
	Code       int16 `json:"code"`
	Timestamp  int64`json:"timestamp"`
	Result     interface{} `json:"result"`
}