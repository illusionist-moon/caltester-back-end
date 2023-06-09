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

	r.POST("/logout", api.Logout)

	// 用户组私有
	authUser := r.Group("/user")
	authUser.Use(middleware.AuthUserCheck())
	{
		authUser.GET("/question", api.GetQuestions)
		authUser.POST("/judge", api.JudgeQuestion)

		authUser.GET("/wrong-list", api.GetWrongList)
		authUser.GET("/wrong-redo", api.GetRedoProblem)
		authUser.POST("/wrong-judge", api.JudgeRedoProblem)

		authUser.GET("/rank", api.GetPointsRank)
		authUser.GET("/points", api.GetOwnPoints)
	}

	return r
}
