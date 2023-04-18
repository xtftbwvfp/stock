// 业务错误码定义

package response

import (
	"github.com/ykstudy/gocore/utils"
	"strings"
)

// 错误码中的 code 定义
const (
	failure = iota - 1
	success
	invalidParam
	notFound
	unknownError
)

// 错误码对象定义
var (
	CodeSuccess       = utils.NewErrCode(success, "Success")
	CodeFailure       = utils.NewErrCode(failure, "Failure")
	CodeInvalidParam  = utils.NewErrCode(invalidParam, "Invalid Param")
	CodeNotFound      = utils.NewErrCode(notFound, "Not Fount")
	CodeInternalError = utils.NewErrCode(unknownError, "Unknown Error")
)

// IsInvalidParamError 判断错误信息中是否包含:参数错误
func IsInvalidParamError(err error) bool {
	return strings.Contains(err.Error(), "Invalid Param")
}
