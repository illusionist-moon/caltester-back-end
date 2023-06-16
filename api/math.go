package api

import (
	"ChildrenMath/models"
	"ChildrenMath/pkg/e"
	"ChildrenMath/service/question"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
	"strings"
)

func GetQuestions(ctx *gin.Context) {
	op, ok := ctx.GetQuery("op")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"msg":  "未传入运算符",
			"data": nil,
		})
		return
	}
	switch op {
	case "plus":
		op = "+"
	case "minus":
		op = "-"
	case "multi":
		op = "*"
	case "div":
		op = "/"
	default:
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"msg":  "非法运算符",
			"data": nil,
		})
		return
	}
	data := question.GenerateQuestions(op)
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"msg":  e.GetMsg(e.Success),
		"data": map[string]any{
			"count":     question.Count,
			"op":        op,
			"questions": data,
		},
	})
}

func JudgeQuestion(ctx *gin.Context) {
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

	op, ok := ctx.GetPostForm("op")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"data": nil,
			"msg":  "miss operator",
		})
		return
	}
	switch op {
	case "plus":
		op = "+"
	case "minus":
		op = "-"
	case "multi":
		op = "*"
	case "div":
		op = "/"
	default:
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"msg":  "非法运算符",
			"data": nil,
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
			"msg":  "invalid answer addPoints, addPoints should be " + strconv.Itoa(question.Count),
		})
		return
	}

	// 开启一个事务，保证错题库和积分的一致性
	tx := models.DB.Begin()
	var (
		addPoints int
		nums      [3]int
		err       error
	)
	for _, q := range answers {
		numStrings := strings.Split(q, ",")
		if len(numStrings) != 3 {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"code": e.InvalidParams,
				"data": nil,
				"msg":  "invalid number addPoints in a question, addPoints must be 3",
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
			addPoints++
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
	err = models.AddPoints(tx, username, addPoints)
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
		"data": addPoints,
		"msg":  "success",
	})
	tx.Commit()
}

func GetWrongList(ctx *gin.Context) {
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

	pageStr, ok := ctx.GetQuery("page")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"data": nil,
			"msg":  "miss page",
		})
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"data": nil,
			"msg":  "invalid page",
		})
		return
	}

	wrongItems, total, getErr := models.GetWrongList(models.DB, username, page)
	if getErr != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.Error,
			"data": nil,
			"msg":  "拉取错题列表失败",
		})
		return
	}
	fmt.Println(total)
	totalPages := int(math.Ceil(float64(total) / models.WrongListOffset))
	if page > totalPages {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"data": nil,
			"msg":  "页号溢出，最大为" + strconv.Itoa(totalPages),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"data": map[string]any{
			"record":       wrongItems,
			"record_count": len(wrongItems),
			"total_page":   totalPages,
		},
		"msg": e.GetMsg(e.Success),
	})
	return
}

func GetRedoProblem(ctx *gin.Context) {
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

	res, err := models.GetRedoProblem(models.DB, username)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.Error,
			"data": nil,
			"msg":  "拉取错误题目失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"data": map[string]any{
			"count":     len(res),
			"questions": res,
		},
		"msg": e.GetMsg(e.Success),
	})
}

func JudgeRedoProblem(ctx *gin.Context) {
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

	answers, ansOK := ctx.GetPostFormArray("answer[]")
	if !ansOK {
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.InvalidParams,
			"data": nil,
			"msg":  "miss answers",
		})
		return
	}

	// 开启一个事务，保证错题库和积分的一致性
	tx := models.DB.Begin()
	var (
		addPoints   int
		id          int
		nums        [3]int // 依次为 num1, num2, res
		op          string
		err         error
		deleteIDSet []int
	)
	for _, ans := range answers {
		data := strings.Split(ans, ",")
		if len(data) != 5 {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"code": e.InvalidParams,
				"data": nil,
				"msg":  "invalid data count in an answer, count must be 5: id, num1, num2, ans, op",
			})
			return
		}

		id, err = strconv.Atoi(data[0])
		if err != nil {
			tx.Rollback()
			ctx.JSON(http.StatusOK, gin.H{
				"code": e.InvalidParams,
				"data": nil,
				"msg":  err.Error(),
			})
			return
		}
		for i := 1; i <= 3; i++ {
			nums[i-1], err = strconv.Atoi(data[i])
			if err != nil {
				tx.Rollback()
				ctx.JSON(http.StatusOK, gin.H{
					"code": e.InvalidParams,
					"data": nil,
					"msg":  "number convert into int failed",
				})
				return
			}
		}
		op = data[4]
		if op != "+" && op != "-" && op != "*" && op != "/" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": e.InvalidParams,
				"data": nil,
				"msg":  "invalid op",
			})
			return
		}
		if question.Judge(nums, op) {
			addPoints++
			deleteIDSet = append(deleteIDSet, id)
		}
	}
	err = models.AddPoints(tx, username, addPoints)
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.Error,
			"data": nil,
			"msg":  "add point failed",
		})
		return
	}

	var deleteCount int64
	deleteCount, err = models.DeleteRedoProblem(tx, deleteIDSet)
	if err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.Error,
			"data": nil,
			"msg":  err.Error(),
		})
		return
	}
	fmt.Println(deleteCount)
	if int(deleteCount) != len(deleteIDSet) {
		tx.Rollback()
		ctx.JSON(http.StatusOK, gin.H{
			"code": e.Error,
			"data": nil,
			"msg":  "delete wrong problems failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": e.Success,
		"data": addPoints,
		"msg":  "success",
	})
	tx.Commit()
}
