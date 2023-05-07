package routers

import (
	"ChildrenMath/api"
	"ChildrenMath/middlewares"
	"ChildrenMath/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cors())

	gin.SetMode(setting.RunMode)

	// 公有方法
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)

	// 用户组私有
	authUser := r.Group("user")
	authUser.Use(middlewares.AuthUserCheck())
	{
		question := authUser.Group("question")
		{
			question.GET("/", api.GetQuestions)

			question.POST("/judgement")
		}

		authUser.GET("/wrong-list")
		authUser.DELETE("/delete")

		// for test ...
		authUser.POST("/test", api.Test)
	}

	return r
}
