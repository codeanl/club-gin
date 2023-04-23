package e

var MsgFlags = map[int]string{
	SUCCESS:       "success",
	ERROR:         "error",
	InvalidParams: "参数不全",

	ErrorNotExistUser: "该用户不存在",
	ErrorNotCompare:   "账号密码错误",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
