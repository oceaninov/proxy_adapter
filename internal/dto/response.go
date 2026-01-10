package dto

const (
	RespSuccessStatus = "success"
	RespFailedStatus  = "failed"
)

type (
	HttpResponse struct {
		ProcessTime string      `json:"process_time"`
		Status      string      `json:"status"`
		Code        string      `json:"code"`
		Message     string      `json:"message"`
		Error       *string     `json:"error,omitempty"`
		Data        interface{} `json:"data,omitempty"`
	}
)

func SuccessHttpResponse(code, message string, data interface{}) HttpResponse {
	return HttpResponse{
		Status:  RespSuccessStatus,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func FailedHttpResponse(code, message string, data interface{}) HttpResponse {
	return HttpResponse{
		Status:  RespFailedStatus,
		Code:    code,
		Message: message,
		Data:    data,
	}
}
