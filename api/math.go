package api

import (
	"ChildrenMath/models"
	"ChildrenMath/pkg/e"
	"ChildrenMath/service/question"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func GetQuestions(ctx *gin.Context) {
	operator, ok := ctx.GetQuery("op")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"msg":  "未传入运算符",
			"data": nil,
		})
		return
	}

	if operator != "plus" && operator != "minus" && operator != "multi" && operator != "div" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"msg":  "非法运算符",
			"data": nil,
		})
		return
	}
	data := question.GenerateQuestions(operator)
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  e.GetMsg(e.Success),
		"data": map[string]any{
			"count":     question.Count,
			"operator":  operator,
			"questions": data,
		},
	})
}

func Judgement(ctx *gin.Context) {
	val, exist := ctx.Get("username")
	// 下面这种情况理论是不存在，但还是需要写出处理
	if !exist {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.ErrorNotExistUser,
			"data": nil,
			"msg":  "用户获取出现问题",
		})
		return
	}
	username := val.(string)

	op, ok := ctx.GetPostForm("operator")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"data": nil,
			"msg":  "miss operator",
		})
		return
	}
	if op != "plus" && op != "minus" && op != "multi" && op != "div" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"data": nil,
			"msg":  "invalid operator",
		})
		return
	}

	answers, ansOK := ctx.GetPostFormArray("answer[]")
	if !ansOK {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"data": nil,
			"msg":  "miss answers",
		})
		return
	}
	if len(answers) != question.Count {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"data": nil,
			"msg":  "invalid answer count, count should be " + strconv.Itoa(question.Count),
		})
		return
	}

	// 开启一个事务，保证错题库和积分的一致性
	tx := models.DB.Begin()
	var (
		count int
		nums  [3]int
		err   error
	)
	for _, q := range answers {
		numStrings := strings.Split(q, ",")
		if len(numStrings) != 3 {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"code": e.InvalidParams,
				"data": nil,
				"msg":  "invalid number count in a question, count must be 3",
			})
			return
		}

		for i := 0; i < 3; i++ {
			nums[i], err = strconv.Atoi(numStrings[i])
			if err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusOK, gin.H{
					"code": e.InvalidParams,
					"data": nil,
					"msg":  "convert into int failed",
				})
				return
			}

		}
		correct := question.Judge(nums, op)
		if correct {
			count++
		} else {
			err = models.AddProblem(tx, username, op, nums[0], nums[1], nums[2])
			if err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusOK, gin.H{
					"code": e.Error,
					"data": nil,
					"msg":  "add incorrect answer failed",
				})
				return
			}
		}
	}
	err = models.AddPoints(tx, username, count)
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.Error,
			"data": nil,
			"msg":  "add point failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"data": count,
		"msg":  "success",
	})
	tx.Commit()
}
