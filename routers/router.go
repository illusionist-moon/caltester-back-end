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
	authUser := r.Group("/user")
	authUser.Use(middlewares.AuthUserCheck())
	{
		authUser.GET("/question", api.GetQuestions)
		authUser.POST("/judgement", api.Judgement)

		authUser.GET("/wrong-list")
		authUser.DELETE("/delete")

		authUser.GET("/logout", api.Logout)
	}

	return r
}
