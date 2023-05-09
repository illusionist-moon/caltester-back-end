package models

type Problem struct {
	ID       int    `json:"id" gorm:"column:id;primary_key;auto_increment"`
	UserName string `json:"username" gorm:"column:user_name"`
	Num1     int    `json:"num1" gorm:"column:num1"`
	Num2     int    `json:"num2" gorm:"column:num2"`
	Operator string `json:"operator" gorm:"column:operator;type:char(1)"`
}

func AddProblem(username, operator string, num1, num2 int) error {
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
	tx := DB.Create(&Problem{
		UserName: username,
		Num1:     num1,
		Num2:     num2,
		Operator: op,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

//func DeleteProblem(id int) bool {
//	DB.Where("user_id = ?", id).Delete(&Problem{})
//	return true
//}
