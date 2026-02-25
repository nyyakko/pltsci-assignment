package http_utils

import "time"

type Error struct {
	Code string `json:"code"`
	Message string `json:"message"`
}

type HttpError struct {
	Timestamp time.Time `json:"timestamp"`
	Path string `json:"path"`
	Error Error `json:"error"`
}

func MakeHttpError(error string, errorCode string, message string) *HttpError {
	return &HttpError{
		Timestamp: time.Now(),
		Error: Error {
			Code: errorCode,
			Message: message,
		},
	}
}

func MakeBadRequestError(errorCode string, message string) *HttpError {
	return MakeHttpError("Bad Request", errorCode, message)
}

func MakeInternalServerError(errorCode string, message string) *HttpError {
	return MakeHttpError("Internal Server Error", errorCode, message)
}
