package main

import (
	"fmt"
	"jwt-verify/token"
	"time"
)

const (
	id     = 121923108723812741
	email  = "nuts_plp@126.com"
	issuer = "pss"
	secret = "密文"
	expire = time.Hour * 3
)

func main() {
	generateToken, err := token.GenerateToken(id, email, issuer, secret, expire)
	if err != nil {
		return
	}
	fmt.Println(generateToken)
	ok, claim := token.ParseToken(generateToken, secret)
	if !ok {
		return
	}
	fmt.Println(claim)

}
