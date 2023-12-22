package main

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Foo string `json:"foo"`
	jwt.RegisteredClaims
}

func main() {
	claims := Claims{
		"bar",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "venkat's macbook air",
		},
	}
	sign := []byte("venky")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(sign)
	fmt.Println(ss)
}
