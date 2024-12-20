package utils

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Success(data interface{}) Response {
	return Response{Code: 200, Msg: "success", Data: data}
}

func Error(code int, msg string) Response {
	return Response{Code: code, Msg: msg}
}
