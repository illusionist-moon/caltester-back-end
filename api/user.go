package api

import (
	"ChildrenMath/models"
	"ChildrenMath/pkg/e"
	"ChildrenMath/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Login(ctx *gin.Context) {
	username := strings.TrimSpace(ctx.PostForm("username"))
	password := ctx.PostForm("password")

	// 判断
	// 用户名格式：长度在 1~20 之间，不能包含空格
	// 用户密码格式：长度大于等于 8
	if err := ctx.ShouldBind(&models.User{
		UserName: username,
		Password: password,
	}); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"msg":  err.Error(),
		})
		return
	}

	getPassword, exists := models.Exists(username)
	if !exists {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ErrorNotExistUser,
			"msg":  e.GetMsg(e.ErrorNotExistUser),
		})
		return
	}

	// 校验密码
	if password != getPassword {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ErrorIncorrectPwd,
			"msg":  e.GetMsg(e.ErrorIncorrectPwd),
		})
		return
	}

	token, err := util.GenerateToken(username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ErrorAuthToken,
			"msg":  e.GetMsg(e.ErrorAuthToken),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":  e.Success,
		"msg":   e.GetMsg(e.Success),
		"token": token,
	})
}

func Register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	rePassword := ctx.PostForm("re-password")

	// 判断
	// 用户名格式：长度在 1~20 之间，不能包含空格
	// 用户密码格式：长度大于等于 8
	if err := ctx.ShouldBind(&models.User{
		UserName: username,
		Password: password,
	}); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"msg":  err.Error(),
		})
		return
	}

	// 判断用户是否存在
	_, exists := models.Exists(username)
	if exists {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ErrorExistUser,
			"msg":  e.GetMsg(e.ErrorExistUser),
		})
		return
	}

	if password != rePassword {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ErrorPwdNotEqual,
			"msg":  e.GetMsg(e.ErrorPwdNotEqual),
		})
		return
	}

	//user := &models.User{
	//	UserName: username,
	//	Password: password,
	//	Points:   0,
	//}
	//err := models.DB.Create(user).Error
	//if err != nil {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"code": e.Error,
	//		"msg":  "Create User Error: " + err.Error(),
	//	})
	//	return
	//}

	err := models.CreateUser(username, password)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.Error,
			"msg":  "Create User Error: " + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  e.GetMsg(e.Success),
	})
}

func Logout(ctx *gin.Context) {
	code := e.Success
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
