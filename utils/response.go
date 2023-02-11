package utils

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponseFormatter(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func ReturnSuccess(data interface{}) Response {
	return ResponseFormatter("success", 200, "success", data)
}

func ReturnError(err error) Response {
	return ResponseFormatter("error", 500, "Internal server error", err)
}
