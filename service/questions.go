package service

import (
	"math/rand"
	"time"
)

const (
	Multi_MAX = 10
	OTHER_MAX = 100
)

type Questions struct {
	Num1 uint32
	Num2 uint32
}

// GenerateQuestions api保证了传入该函数的operator一定是正确且合法的
func GenerateQuestions(operator string) []*Questions {
	var max uint32
	if operator == "multi" {
		max = Multi_MAX
	} else {
		max = OTHER_MAX
	}
	ans := make([]*Questions, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		n1 := r.Uint32() % max
		var n2 uint32
		if operator == "div" || operator == "minus" {
			for {
				if n1 != 0 {
					break
				}
				n1 = r.Uint32() % max
			}
			n2 = r.Uint32() % n1
		} else {
			n2 = r.Uint32() % max
		}
		if operator == "div" {
			for {
				if n2 != 0 {
					break
				}
				n2 = r.Uint32() % n1
			}
		}
		ans = append(ans, &Questions{
			Num1: n1,
			Num2: n2,
		})
	}
	return ans
}
