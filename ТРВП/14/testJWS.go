package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("secret_key")

func createToken() (string, error) {
	claims := jwt.MapClaims{
		"sub":  "user222",
		"name": "Dianich",
		"role": "admin",
		"exp":  time.Now().Add(time.Minute * 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

func parseToken(tokenStr string) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		fmt.Println("Invalid token:", err)
		return
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("TOKEN VALID ✅")
		fmt.Println("User ID:", claims["sub"])
		fmt.Println("Name:", claims["name"])
		fmt.Println("Role:", claims["role"])
	} else {
		fmt.Println("Invalid token")
	}
}

func main() {
	token, err := createToken()
	if err != nil {
		panic(err)
	}

	fmt.Println("JWT TOKEN:")
	fmt.Println(token)
	fmt.Println()

	parseToken(token)
}
