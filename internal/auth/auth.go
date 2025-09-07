package auth

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)


var jwtToken = []byte(os.Getenv("JWT_SECRET"))

type Claims struct{
	Userid int`json:"id"`
	Userlogin string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(id int, login string) (string, error){
	claims := &Claims{
		Userid: id,
		Userlogin: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)), 
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "myapp",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtToken)
}