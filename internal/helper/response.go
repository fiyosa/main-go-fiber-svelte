package helper

import "go-fiber-svelte/internal/lang"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type Meta struct {
	Total     int `json:"total"`
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
}

func ResSuccess(msg string, status ...int) *Response {
	return &Response{Message: msg}
}

func ResSuccessData(data interface{}, msg string, status ...int) *Response {
	return &Response{Message: msg, Data: data}
}

func ResError(msg string, errors interface{}, status ...int) *Response {
	return &Response{Message: msg, Errors: errors}
}

func ResPaginate(data interface{}, meta *Meta, msg string) *Response {
	return &Response{Message: msg, Data: data, Meta: meta}
}

func ResCatch(err error) *Response {
	return &Response{Message: lang.T("internal_error")}
}

func ResValidate(err error) *Response {
	return &Response{Message: lang.T("validation_error"), Errors: err.Error()}
}
