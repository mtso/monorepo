package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

var payloadopt = flag.String("payload", "", "string payload option")

var (
	header = []byte(`{"typ":"JWT","alg":"HS256"}`)
)

func main() {
	flag.Parse()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":     "bar",
		"message": "hello",
	})
	tokenstring, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", tokenstring)
}
