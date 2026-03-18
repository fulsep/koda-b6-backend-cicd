package lib

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AppSecret = []byte(os.Getenv("APP_SECRET"))

type CustomClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int) string {
	claims := CustomClaims{
		userId,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, _ := token.SignedString(AppSecret)

	return ss
}

func VerifyToken(tokenStr string) bool {
	token, _ := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
		return AppSecret, nil
	})
	return token.Valid
}
