package helper

import "go-fiber-svelte/internal/lang"

var Res res

type res struct{}

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

func (res) Success(msg string, status ...int) *Response {
	return &Response{Message: msg}
}

func (res) SuccessData(data interface{}, msg string, status ...int) *Response {
	return &Response{Message: msg, Data: data}
}

func (res) Error(msg string, errors interface{}, status ...int) *Response {
	return &Response{Message: msg, Errors: errors}
}

func (res) Paginate(data interface{}, meta *Meta, msg string) *Response {
	return &Response{Message: msg, Data: data, Meta: meta}
}

func (res) Catch(err error) *Response {
	return &Response{Message: lang.T.Get().SOMETHING_WENT_WRONG}
}

func (res) Validate(err error) *Response {
	return &Response{Message: "Validation error", Errors: err.Error()}
}
