package models

import "gorm.io/gorm"

type User struct {
	UserName string `json:"username" gorm:"column:user_name;primary_key" binding:"required,alphanum,min=1,max=20,excludes= "`
	Password string `json:"password" gorm:"column:password" binding:"required,min=8"`
	Points   int    `json:"points" gorm:"column:points" binding:"-"`
}

func Exists(username string) (string, bool) {
	var user User
	// 按照 用户名 进行查找
	err := DB.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return "", false
	} else {
		return user.Password, true
	}
}

func AddPoints(tx *gorm.DB, username string, points int) error {
	var target User
	err := tx.Where("user_name = ?", username).Take(&target).Error
	//err := tx.Take(&target, username).Error
	if err != nil {
		return err
	}
	target.Points += points
	tx.Select("Points").Save(&target)
	return nil
}

func CreateUser(username, password string) error {
	user := &User{
		UserName: username,
		Password: password,
		Points:   0,
	}
	err := DB.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
