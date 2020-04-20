package base

//定义Code的常用的别名
type ResCode int

const (
	ResCodeOk                  ResCode = 1000
	ResCodeValidationError     ResCode = 2000
	ResCodeRequestParamsError  ResCode = 2100
	ResCodeInnerServerError    ResCode = 5000
	ResCodeBussError           ResCode = 6000 //业务上的异常
	ResCodeBissTransferFailure ResCode = 6010
)

//定义web response的结构体

type Res struct {
	Code    ResCode     `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
