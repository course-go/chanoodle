package common

type Response struct {
	Data  any    `json:"data,omitzero"`
	Error string `json:"error,omitzero"`
}

func NewDataResponse(data any) Response {
	return Response{
		Data: data,
	}
}

func NewErrorResponse(err error) Response {
	return Response{
		Error: err.Error(),
	}
}
