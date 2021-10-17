package tools

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResultApi(Code int, Message string, Data interface{}) interface{} {
	result := Result{
		Code:    Code,
		Message: Message,
		Data:    Data,
	}
	return result
}
