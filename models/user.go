package models

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

func AddPoints(username string, points int) error {
	var target User
	tx := DB.Take(&target, "user_name = ?", username)
	if tx.Error != nil {
		return tx.Error
	}
	target.Points += points
	DB.Select("Points").Save(&target)
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
