package middleware

import (
	"errors"
	"fmt"

	"ctw-interview/common"
	"github.com/golang-jwt/jwt/v5"
)

var (
	TokenInvalid = errors.New("Couldn't handle this token:")
)
var mySigningKey = []byte("ushjlwmwnwht")

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		SigningKey: mySigningKey,
	}
}

// TODO SigningKey 优化配置

type UserClaims struct {
	UserId int64
	jwt.RegisteredClaims
}

func GenerateToken(userID int64) (string, error) {
	userClaims := UserClaims{
		UserId:           userID,
		RegisteredClaims: jwt.RegisteredClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	return token.SignedString(mySigningKey)
}

func (j *JWT) ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		common.SysError(fmt.Sprintf("jwt ParseWithClaims error: %v", err))
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
