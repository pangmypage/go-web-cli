package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type JWTUser struct {
	ID string
	jwt.StandardClaims
}

//创建一个token
func CreatToken(id string) (string, error) {
	//设置token内容
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := viper.GetString("app.jwt_secret")
	return token.SignedString([]byte(key))
}

//解析token
func ParseToken(c *gin.Context) (*JWTUser, error) {
	Authorization := c.Request.Header.Get("Authorization")
	if Authorization == "" || len(Authorization) == 0 {
		return nil, errors.New("请登录")
	}
	key := viper.GetString("app.jwt_secret")
	t, err := jwt.Parse(Authorization, secret(key))
	if err != nil {
		return nil, errors.New("token解析失败")
	}
	if claim, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		fmt.Println(claim, ok)
		user := &JWTUser{}
		user.ID = fmt.Sprintf("%v", claim["id"])
		user.ExpiresAt = int64(claim["exp"].(float64))
		return user, nil
	} else {
		return nil, errors.New("token验证失败")
	}

}

//验证密匙
func secret(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	}
}
