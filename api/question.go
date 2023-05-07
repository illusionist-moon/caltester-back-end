package api

import (
	"ChildrenMath/pkg/e"
	"ChildrenMath/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetQuestions(ctx *gin.Context) {
	operator, ok := ctx.GetQuery("op")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  "未传入运算符",
			"data": nil,
		})
		return
	}

	if operator != "plus" && operator != "minus" && operator != "multi" && operator != "div" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  "非法运算符",
			"data": nil,
		})
		return
	}
	data := service.GenerateQuestions(operator)
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.SUCCESS,
		"msg":  e.GetMsg(e.SUCCESS),
		"data": map[string]any{
			"count":     10,
			"operator":  operator,
			"questions": data,
		},
	})
}
