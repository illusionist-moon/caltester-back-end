package models

type WrongQuestions struct {
	ID       int    `json:"id" gorm:"primary_key:auto_increment"`
	UserID   int    `json:"user_id" gorm:"primary_key"`
	Num1     int    `json:"num1"`
	Num2     int    `json:"num2"`
	Operator string `json:"operator" gorm:"type:char(1)"`
}

func AddWrongQuestion(data map[string]any) bool {
	DB.Create(&WrongQuestions{
		UserID:   data["user_id"].(int),
		Num1:     data["num1"].(int),
		Num2:     data["num2"].(int),
		Operator: data["operator"].(string),
	})
	return true
}

func DeleteWrongQuestion(id int) bool {
	DB.Where("user_id = ?", id).Delete(&WrongQuestions{})
	return true
}
