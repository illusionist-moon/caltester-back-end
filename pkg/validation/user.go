package validation

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

type UserLogin struct {
	UserName string `json:"username" binding:"required,alphanumunicode,min=1,max=20" msg:"UserName: 用户名长度在1-20之间，仅包含Unicode字符与数字"`
	Password string `json:"password" binding:"required,min=8,noCRLF" msg:"Password: 用户密码长度不小于8，不能包含换行符和回车"`
}

type UserRegister struct {
	UserName   string `json:"username" binding:"required,alphanumunicode,min=1,max=20" msg:"UserName: 用户名长度在1-20之间，仅包含Unicode字符与数字"`
	Password   string `json:"password" binding:"required,min=8,noCRLF" msg:"Password: 用户密码长度不小于8，不能包含换行符和回车"`
	RePassword string `json:"re-password" binding:"required,eqfield=Password" msg:"RePassword: 两次密码输入需要一致"`
}

// user defined validator
// 定义密码验证
func noCRLF(fl validator.FieldLevel) bool {
	passwd := fl.Field().String()
	if strings.Contains(passwd, "\r") || strings.Contains(passwd, "\n") {
		return false
	}
	return true
}
