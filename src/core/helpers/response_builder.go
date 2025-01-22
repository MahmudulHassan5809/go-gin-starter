package helpers

import (
	"gin_starter/src/core/errors"
	"net/http"
)


type BaseResponse struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	Code     int         `json:"code"`
	MetaInfo interface{} `json:"meta_info,omitempty"`
	ValidationErrors interface{} `json:"validation_errors,omitempty"` // New field
}


type SuccessResponse[T any] struct {
	BaseResponse
	Data T `json:"data,omitempty"`
}


type ErrorResponse struct {
	BaseResponse
	Data interface{} `json:"data,omitempty"`
}


type ResponseOptions struct {
	Message  string
	MetaInfo interface{}
	Error    *errors.CustomError
	ValidationErrors interface{}
}


func WithError(err *errors.CustomError) func(*ResponseOptions) {
	return func(opt *ResponseOptions) {
		opt.Error = err
	}
}

func WithMessage(message string) func(*ResponseOptions) {
	return func(opt *ResponseOptions) {
		opt.Message = message
	}
}

func WithMetaInfo(metaInfo interface{}) func(*ResponseOptions) {
	return func(opt *ResponseOptions) {
		opt.MetaInfo = metaInfo
	}
}




func buildBaseResponse(success bool, defaultMessage string, defaultCode int, opts ...func(*ResponseOptions)) BaseResponse {
	options := ResponseOptions{
		Message:  defaultMessage,
		MetaInfo: map[string]interface{}{},
		
	}

	for _, opt := range opts {
		opt(&options)
	}

	message := options.Message
	code := defaultCode
	var validationErrors interface{} = map[string]interface{}{}

	if options.Error != nil {
		if success {
			message = options.Message
			
		} else {
			message = options.Error.Msg
			code = options.Error.Status
			validationErrors = options.Error.ValidationErrors
		}
	}

	return BaseResponse{
		Success:  success,
		Message:  message,
		Code:     code,
		MetaInfo: options.MetaInfo,
		ValidationErrors: validationErrors,
	}
}


func BuildSuccessResponse[T any](data T, opts ...func(*ResponseOptions)) SuccessResponse[T] {
	return SuccessResponse[T]{
		BaseResponse: buildBaseResponse(true, "OK", http.StatusOK, opts...),
		Data:         data,
	}
}

func BuildErrorResponse(opts ...func(*ResponseOptions)) ErrorResponse {
	return ErrorResponse{
		BaseResponse: buildBaseResponse(false, "An unexpected error occurred", http.StatusBadRequest, opts...),
		Data:         map[string]string{},
	}
}
