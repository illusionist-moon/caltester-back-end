package models

import (
	"gorm.io/gorm"
)

const (
	RankSize = 10
)

type User struct {
	UserName string `json:"username" gorm:"column:user_name;primary_key"`
	Password string `json:"password" gorm:"column:password"`
	Points   int    `json:"points" gorm:"column:points"`
}

type Rank struct {
	UserName string `json:"username" gorm:"column:user_name"`
	Points   int    `json:"points" gorm:"column:points"`
}

func Exists(db *gorm.DB, username string) (string, bool) {
	var user User
	// 按照 用户名 进行查找
	err := db.Select("password").Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return "", false
	} else {
		return user.Password, true
	}
}

func AddPoints(db *gorm.DB, username string, points int) error {
	var target User
	err := db.Where("user_name = ?", username).Take(&target).Error
	if err != nil {
		return err
	}
	target.Points += points
	db.Select("Points").Save(&target)
	return nil
}

func CreateUser(db *gorm.DB, username, password string) error {
	user := &User{
		UserName: username,
		Password: password,
		Points:   0,
	}
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPointsRank(db *gorm.DB) ([]Rank, error) {
	var res []Rank
	err := db.Model(&User{}).Select("points, user_name").Limit(RankSize).Order("points desc, user_name").Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserPoints(db *gorm.DB, username string) (int, error) {
	var res int
	err := db.Model(&User{}).Select("points").Where("user_name = ?", username).Take(&res).Error
	if err != nil {
		return 0, err
	}
	return res, nil
}

func GetUserRank(db *gorm.DB, username string) (int, error) {
	var (
		res int64
		err error
	)
	subQuery := db.Model(&User{}).Select("points").Where("user_name = ?", username)
	err = db.Model(&User{}).Where("points > (?)", subQuery).Find(&User{}).Count(&res).Error
	if err != nil {
		return 0, err
	}
	return int(res) + 1, nil
}
