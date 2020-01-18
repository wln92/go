package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	key = []byte("Hello World! This is secret!")
)

func GenToken() string {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 1000000),
		Issuer: "test",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return ""
	}
	return ss
}

func checkToken(token string) bool {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error){
		return key, nil
	})
	if err != nil {
		fmt.Printf("check token err:%v\n", err)
	}
	fmt.Printf("%v,%v, %v, %v", t.Header, t.Method, t.Signature, t.Claims)
	return true
}

func main() {
	token := GenToken()
	fmt.Printf("gen token:%v\n", token)
	checkToken(token)
}
