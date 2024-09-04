package res

type ErrorCode int

const (
	SettingsError ErrorCode = 1001 //系统错误
	AgrumentError ErrorCode = 1002 // 参数错误
)

// 错误状态码
var (
	ErrorMap = map[ErrorCode]string{
		SettingsError: "系统错误",
		AgrumentError: "参数错误",
	}
)
