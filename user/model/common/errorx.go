package common

// 全局的 错误号 类型，用于API调用之间传递
type ServerErrorCode int

// 全局的 错误号 的具体定义
const (
	InsertUserErrorCode = 1000 //插入users出错
	GetUserErrorCode    = 1001 //获取用户信息出错
)

// 内部的错误map，用来对应 错误号和错误信息
var ErrCodeMap = map[ServerErrorCode]string{

	InsertUserErrorCode: "insert into users error",
	GetUserErrorCode:    "get user info error",
}

// Sentinel Error： 即全局定义的Static错误变量
// 注意，这里的全局error是没有保存堆栈信息的，所以需要在初始调用处使用 errors.Wrap
var (
	InsertUserError = NewServerError(InsertUserErrorCode)
	GetUserError = NewServerError(GetUserErrorCode)
)

func NewServerError(code ServerErrorCode) *ServerError {
	return &ServerError{
		Code:    code,
		Message: ErrCodeMap[code],
	}
}

// error的具体实现
type ServerError struct {
	// 对外使用 - 错误码
	Code ServerErrorCode
	// 对外使用 - 错误信息
	Message string
}

func (e *ServerError) Error() string {
	return e.Message
}
