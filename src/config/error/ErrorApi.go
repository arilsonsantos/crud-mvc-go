package errApi

import "fmt"

type ErrApi struct {
	HttpCode        int    `json:"httpCode"`
	ErrorCode       string `json:"errorCode"`
	Message         string `json:"message"`
	DetailedMessage string `json:"detailedMessage"`
}

func NewErrApi(httpCode int, errorCode string, message string, detailedMessage string) *ErrApi {
	return &ErrApi{
		HttpCode:        httpCode,
		ErrorCode:       errorCode,
		Message:         message,
		DetailedMessage: detailedMessage,
	}
}

func NewBadRequest(message string) *ErrApi {
	return NewErrApi(
		400,
		"MS-GO-400",
		"Bad Request",
		message)
}

func NewInternalServerError(message string) *ErrApi {
	return NewErrApi(
		500,
		"MS-GO-500",
		"Internal Server Error",
		message)
}

func (e *ErrApi) Error() string {
	return fmt.Sprintf("Error ao processar, code=%d, message=%s", e.HttpCode, e.DetailedMessage)
}
