package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type UserClaims struct {
	UserName string `json:username`
	jwt.StandardClaims
}

var myKey = []byte("Soft-ware-Engineering-Team-Project-Children-Math")

func GenerateToken(username string) (string, error) {
	// 令牌的有限时间为 24 小时
	expireTime := time.Now().Add(time.Hour * 24)
	uerClaim := &UserClaims{
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uerClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyseToken(tokenString string) (*UserClaims, error) {
	UserClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, UserClaims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token error")
	}
	return UserClaims, nil
}

//
//func main() {
//	tokenString, _ := GenerateToken("Ztyx")
//	fmt.Println(tokenString)
//	fmt.Println(AnalyseToken(tokenString))
//}
