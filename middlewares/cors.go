package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")
		if origin != "" {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-Custom-Header, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			ctx.Header("Access-Control-Expose-Headers", "Authorization, Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			ctx.Header("Access-Control-Max-Age", "172800")
			ctx.Set("content-type", "application/json") // 设置返回格式是json
		}
		if method == "OPTIONS" {
			ctx.JSON(http.StatusOK, "Options Request!")
		}
		//处理请求
		ctx.Next()
	}
}
