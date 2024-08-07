package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/Sunikka/termitalk/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

var jwt_key string = os.Getenv("JWT_SECRET")

func GenerateToken(user utils.User) (string, error) {
	// TODO: What claims are necessary?
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(jwt_key))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func validateJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwt_key), nil
	})
}
