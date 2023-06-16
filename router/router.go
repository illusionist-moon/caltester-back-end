package router

import (
	"ChildrenMath/api"
	"ChildrenMath/middleware"
	"ChildrenMath/pkg/settings"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())

	gin.SetMode(settings.RunMode)

	// 公有方法
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)

	// 用户组私有
	authUser := r.Group("/user")
	authUser.Use(middleware.AuthUserCheck())
	{
		authUser.GET("/question", api.GetQuestions)
		authUser.POST("/judgement", api.Judgement)

		authUser.GET("/wrong-list", api.GetWrongList)
		authUser.GET("/wrong-redo", api.RedoWrongQuestion)
		authUser.POST("/wrong-judge")

		authUser.GET("/logout", api.Logout)

		authUser.GET("/rank", api.GetPointsRank)
	}

	return r
}
