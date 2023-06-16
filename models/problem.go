package models

import "gorm.io/gorm"

const (
	WrongListOffset  = 10
	RedoProblemCount = 10
)

type Problem struct {
	ID       int    `json:"id" gorm:"column:id;primary_key;auto_increment"`
	UserName string `json:"username" gorm:"column:user_name"`
	Num1     int    `json:"num1" gorm:"column:num1"`
	Num2     int    `json:"num2" gorm:"column:num2"`
	WrongAns int    `json:"wrong_ans" gorm:"column:wrong_ans"`
	Operator string `json:"operator" gorm:"column:operator;type:char(1)"`
}

type WrongListItem struct {
	Num1     int    `json:"num1" gorm:"column:num1"`
	Num2     int    `json:"num2" gorm:"column:num2"`
	WrongAns int    `json:"wrong_ans" gorm:"column:wrong_ans"`
	Operator string `json:"operator" gorm:"column:operator"`
}

type RedoProblem struct {
	ID       int    `json:"id" gorm:"column:id;primary_key;auto_increment"`
	Num1     int    `json:"num1" gorm:"column:num1"`
	Num2     int    `json:"num2" gorm:"column:num2"`
	Operator string `json:"operator" gorm:"column:operator;type:char(1)"`
}

func AddProblem(db *gorm.DB, username, operator string, num1, num2, wrongAns int) error {
	var op string
	switch operator {
	case "plus":
		op = "+"
	case "minus":
		op = "-"
	case "multi":
		op = "*"
	case "div":
		op = "/"
	default:
		// 事实上，这一步永远不会走到
	}
	err := db.Create(&Problem{
		UserName: username,
		Num1:     num1,
		Num2:     num2,
		WrongAns: wrongAns,
		Operator: op,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

// GetWrongList 返回值依次为(page对应页的错题，错题的总数，错误)
func GetWrongList(db *gorm.DB, username string, page int) ([]WrongListItem, int64, error) {
	var res []WrongListItem
	var total int64
	err := db.Model(&Problem{}).Omit("id", "user_name").Where("user_name = ?", username).Count(&total).Limit(WrongListOffset).Offset((page - 1) * WrongListOffset).
		Find(&res).Error
	if err != nil {
		return nil, 0, err
	}
	return res, total, nil
}

func GetRedoProblem(db *gorm.DB, username string) ([]RedoProblem, error) {
	var res []RedoProblem
	err := db.Model(&Problem{}).Omit("user_name", "wrong_ans").Where("user_name = ?", username).Limit(RedoProblemCount).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

//func DeleteProblem(id int) bool {
//	DB.Where("user_id = ?", id).Delete(&Problem{})
//	return true
//}
