package e

var MsgFlags = map[int]string{
	SUCCESS:        "success",
	ERROR:          "fail",
	INVALID_PARAMS: "参数错误",

	ERROR_EXIST_USER:     "用户名已存在",
	ERROR_NOT_EXIST_USER: "该用户不存在",
	ERROR_INCORRECT_PWD:  "用户存在但密码错误",
	ERROR_PWD_NOT_EQUAL:  "两次密码不一致",

	ErrorAuthCheckTokenFail:    "Token鉴权失败",
	ErrorAuthCheckTokenTimeout: "Token已超时",
	ErrorAuthToken:             "Token生成失败",

	PAGE_NOT_FOUND: "Page not found",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
