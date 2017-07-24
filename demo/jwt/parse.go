package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var tokenstring = flag.String("token", "", "must provide")

func main() {
	flag.Parse()
	if *tokenstring == "" {
		log.Fatal("must provide a token")
	}

	token, err := jwt.Parse(*tokenstring, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["message"])
	} else {
		fmt.Println(err)
	}
}
