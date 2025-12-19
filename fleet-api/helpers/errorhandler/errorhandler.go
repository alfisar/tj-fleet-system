package errorhandler

import (
	"fleet-api/domain"

	"github.com/valyala/fasthttp"
)

const (
	// code series 1xx for validation data
	// Error code for invalid input ex : email format not valid, type data not valid, params not valid
	ErrCodeInvalidInput int = 101

	// code series 2xx for error DB

	// Error code for invalid connection DB
	ErrCodeConnection int = 201

	// Error code for data is empty or not found
	ErrCodeDataNotFound int = 202

	// Error code for save data
	ErrCodeInsert int = 203

	// Error code for update data
	ErrCodeUpdate int = 204

	// Error code for delete data
	ErrCodeDelete int = 205

	// Error code for get data
	ErrCodeGet int = 206

	// Error code for data is empty or not found
	ErrCodeBlocked int = 207

	// code series 5xx for error internal ex: error hashing

	// Error code for generate token
	ErrCodeGenerateToken int = 501

	// Error code for parsing data
	ErrCodeParsing int = 502

	// Error code for panic
	ErrCodePanic int = 503

	// Error code for internal server error
	ErrCodeInternalServer int = 504

	// Error MSG fo conncention is nil
	ErrMsgConnEmpty string = "connection is nil"

	// Error message for invalid data query
	ErrInvalidDataQuery string = "Invalid Data Query"
)

func ErrValidation(err error) (result domain.ErrorData) {
	result.Status = "error"
	result.Code = ErrCodeInvalidInput
	result.HTTPCode = fasthttp.StatusBadRequest
	result.Message = "Invalid data input"
	result.Errors = err.Error()
	return
}

func ErrRecordNotFound() (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:   "error",
		HTTPCode: fasthttp.StatusBadRequest,
		Code:     ErrCodeDataNotFound,
		Message:  "Data not found",
	}

	return
}

func ErrGetData(err error) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:   "error",
		Code:     ErrCodeGet,
		HTTPCode: fasthttp.StatusBadRequest,
		Message:  "Failed get data",
		Errors:   err.Error(),
	}

	return
}

func ErrInternal(code int, err error) (result domain.ErrorData) {
	result = domain.ErrorData{
		Status:   "error",
		Code:     code,
		HTTPCode: fasthttp.StatusInternalServerError,
		Message:  "Internal errors",
		Errors:   err.Error(),
	}

	return
}
