package api

import (
	"ChildrenMath/models"
	"ChildrenMath/pkg/e"
	"ChildrenMath/pkg/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	var user models.User
	// 按照 用户名 进行查找
	err := models.DB.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		// 记录未找到的错误，即用户不存在
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusOK, gin.H{
				"code": e.ErrorNotExistUser,
				"msg":  e.GetMsg(e.ErrorNotExistUser),
			})
			return
		}
		// 其他错误，未知
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.Error,
			"msg":  "Get User Error: " + err.Error(),
		})
		return
	}

	// 校验密码
	if password != user.Password {
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

	var count int64
	// 判断用户名是否已经存在
	err := models.DB.Where("user_name = ?", username).Model(new(models.User)).Count(&count).Error
	if err != nil {
		// 其他错误，未知
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.Error,
			"msg":  "Get User Error: " + err.Error(),
		})
		return
	}
	if count > 0 {
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

	user := &models.User{
		UserName: username,
		Password: password,
		Points:   0,
	}
	err = models.DB.Create(user).Error
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

func Test(ctx *gin.Context) {
	fmt.Println("ok ok ok !!!")
}

func Logout(ctx *gin.Context) {
	code := e.Success
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
	})
}
