package common

// 状态码
const (
	OK               = 200 // Success
	NotLoggedIn      = 700 // 未登录
	ParameterIllegal = 701 // 参数不合法
	Unauthorized     = 743 // 未授权
	ServerError      = 750 // 系统错误
)

// GetErrorMessage 根据错误码 获取错误信息
func GetErrorMessage(code uint64, message string) string {
	var codeMessage string
	codeMap := map[uint64]string{
		OK:               "Success",
		NotLoggedIn:      "未登录",
		ParameterIllegal: "参数不合法",
		Unauthorized:     "未授权",
		ServerError:      "系统错误",
	}

	if message == "" {
		if value, ok := codeMap[code]; ok {
			// 存在
			codeMessage = value
		} else {
			codeMessage = "未定义错误类型!"
		}
	} else {
		codeMessage = message
	}

	return codeMessage
}
