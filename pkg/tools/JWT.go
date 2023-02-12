package tools

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type CustomClaims struct {
	Username string `json:"username"`
	User_id  int64  `json:"user_id"`

	jwt.RegisteredClaims // 内嵌标准的声明
}

var CustomSecret = []byte("BAT_DY")

func GenToken(username string, userid int64) (string, error) {
	cliams := CustomClaims{
		Username: username,
		User_id:  userid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			Issuer:    "zxj", // 签发人

		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)

	return token.SignedString(CustomClaims{})
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
