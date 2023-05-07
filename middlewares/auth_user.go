package middlewares

import (
	"ChildrenMath/pkg/e"
	"ChildrenMath/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AuthUserCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		userClaim, err := util.AnalyseToken(auth)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": e.ErrorAuthCheckTokenFail,
				"msg":  e.GetMsg(e.ErrorAuthCheckTokenFail),
			})
			ctx.Abort()
			return
		}
		if time.Now().Unix() > userClaim.ExpiresAt {
			ctx.JSON(http.StatusOK, gin.H{
				"code": e.ErrorAuthCheckTokenTimeout,
				"msg":  e.GetMsg(e.ErrorAuthCheckTokenTimeout),
			})
			ctx.Abort()
			return
		}
		ctx.Set("username", userClaim.UserName)
		ctx.Next()
	}
}
