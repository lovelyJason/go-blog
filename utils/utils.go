package utils

import (
	gojwt "github.com/dgrijalva/jwt-go"
	"time"
)

// var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var jwtSecret = []byte("jasonhuang")

type Claims struct {
	Uid int
	gojwt.StandardClaims
}

func Award(uid *int) (string, error) {
	expireTime := time.Now().Add(1 * 24 * time.Hour)
	claims := &Claims{
		Uid: *uid,
		StandardClaims: gojwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	//	生成token
	token := gojwt.NewWithClaims(gojwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (*gojwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := gojwt.ParseWithClaims(tokenStr, claims, func(t *gojwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, claims, nil
}
