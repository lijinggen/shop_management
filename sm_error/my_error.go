package sm_error

import (
	"github.com/go-playground/validator/v10"
	"github.com/shop_management/sm_error/error_code"
	"github.com/shop_management/util"
	"strings"
)

var ErrMap map[int]string

func init() {
	ErrMap = make(map[int]string)
	ErrMap[error_code.UserPhoneExists] = "手机号已经存在"
}

// define 000 00000
type Error struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

func (e *Error) Error() string {
	return util.MarshalToStringNoErr(e)
}

func NewHttpError(errorCode int, errMsg ...string) *Error {
	fErrMsg := ErrMap[errorCode]
	if len(errMsg) != 0 {
		fErrMsg = strings.Join(errMsg, "; ")
	}
	return &Error{
		ErrorCode: errorCode,
		ErrorMsg:  fErrMsg,
	}
}

func NewParamHttpError(err error) *Error {
	errors, ok := err.(validator.ValidationErrors)
	if !ok {
		return &Error{
			ErrorCode: error_code.ReqParamError,
			ErrorMsg:  err.Error(),
		}
	}
	result := ""
	for _, fieldError := range errors {
		rawResult := ""
		fields := strings.Split(fieldError.StructNamespace(), ".")
		if len(fields) > 0 {
			rawResult = fields[len(fields)-1] + " entered incorrectly; "
		}
		result += rawResult
	}
	return &Error{
		ErrorCode: error_code.ReqParamError,
		ErrorMsg:  result,
	}
}

func GetErrMsg(errorCode int) string {
	return ErrMap[errorCode]
}
