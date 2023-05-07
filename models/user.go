package models

type User struct {
	UserName string `json:"username" gorm:"column:user_name;primary_key" binding:"required,alphanum,min=1,max=20,excludes= "`
	Password string `json:"password" gorm:"column:password" binding:"required,min=8"`
	Points   int    `json:"points" gorm:"column:points" binding:"-"`
}

func GerWrongProblems() {

}
